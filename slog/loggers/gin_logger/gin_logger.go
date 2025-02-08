package gin_logger

import (
	"context"
	"fmt"
	"github.com/INT-Game/go-tools/slog/log_context"
	"github.com/INT-Game/go-tools/slog/loggers"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// SetupGinEngineZapLogger sets up the gin engine log with zap logger.
func SetupGinEngineZapLogger(r *gin.Engine, zapLogger *zap.Logger) {
	r.Use(GinZapHandler)
	r.Use(GinZapRecoveryHandler)
}

// SetupGinZapLogger sets up the gin debug log with zap logger.
func SetupGinZapLogger(zapLogger *zap.Logger) {
	logger := zapLogger.WithOptions(zap.AddCallerSkip(2))
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		logger.Sugar().Debugf("%-6s %-25s --> %s (%d handlers)", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	gin.DebugPrintFunc = func(format string, values ...interface{}) {
		logger.Sugar().Debugf(format, values...)
	}
}

func GetGinTraceCtx(ctx context.Context, c *gin.Context) context.Context {
	ctx = log_context.SetLogContextKeyValue(ctx, log_context.CtxRequestId, getIdFromGinContext(c, log_context.GinCtxRequestIdKeyStr, log_context.GinCtxRequestIdKeyStr))
	ctx = log_context.SetLogContextKeyValue(ctx, log_context.CtxTraceId, getIdFromGinContext(c, log_context.GinCtxTraceIdKeyStr, log_context.GinCtxTraceIdKeyStr))
	c.Set("ctx", ctx)
	return ctx
}

// func GetGinTraceCtxWithKeys(ctx context.Context, c *gin.Context, keys ...string) context.Context {
// 	for _, key := range keys {
// 		ctx = SetContextKeyValue(ctx, key, getIdFromGinContext(c, key, ""))
// 	}
// 	return ctx
// }

func getIdFromGinContext(c *gin.Context, key string, defaultKey string) string {
	if key == "" {
		key = defaultKey
	}
	id := c.GetHeader(key)
	if id == "" {
		id = log_context.NewId()
	}
	return id
}

var GinZapHandler gin.HandlerFunc = func(c *gin.Context) {
	start := time.Now()
	// some evil middlewares modify this values
	path := c.Request.URL.Path
	query := c.Request.URL.RawQuery
	c.Next()
	track := true

	if track {
		end := time.Now()
		latency := end.Sub(start)

		fields := []any{
			"status", c.Writer.Status(),
			"method", c.Request.Method,
			"path", path,
			"query", query,
			"ip", c.ClientIP(),
			"ua", c.Request.UserAgent(),
			"latency", latency,
		}
		ctx := getCtxFromGinContext(c)

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			errorMsg := ""
			for i, e := range c.Errors.Errors() {
				errorMsg += fmt.Sprintf("[%d]: %s\n", i, e)
			}
			loggers.CLogw(ctx, zap.ErrorLevel, 2, errorMsg, fields...)
		} else {
			msg := fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL.Path)
			loggers.CLogw(ctx, zap.InfoLevel, 2, msg, fields...)
		}
	}
}

var GinZapRecoveryHandler gin.HandlerFunc = func(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			// Check for a broken connection, as it is not really a
			// condition that warrants a panic stack trace.
			var brokenPipe bool
			if ne, ok := err.(*net.OpError); ok {
				if se, ok := ne.Err.(*os.SyscallError); ok {
					if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
						strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
						brokenPipe = true
					}
				}
			}

			httpRequest, _ := httputil.DumpRequest(c.Request, false)
			ctx := getCtxFromGinContext(c)

			if brokenPipe {
				loggers.CLogw(
					ctx, zap.ErrorLevel,
					2,
					fmt.Sprintf("broken pipe request: %s err: %s", string(httpRequest), err),
				)
				// If the connection is dead, we can't write a status to it.
				c.Error(err.(error)) //nolint: errcheck
				c.Abort()
				return
			}

			loggers.CLogw(
				ctx, zap.ErrorLevel,
				2,
				fmt.Sprintf("[Recovery from panic] request: %s err: %s", string(httpRequest), err),
			)

			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}()
	c.Next()
}

func getCtxFromGinContext(c *gin.Context) (ctx context.Context) {
	cctx, ok := c.Get("ctx")
	if ok {
		ctx, ok = cctx.(context.Context)
		if ok {
			return ctx
		}
	}
	return context.Background()
}
