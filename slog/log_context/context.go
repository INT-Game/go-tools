package log_context

import (
	"context"

	"github.com/google/uuid"
)

type ctxKey string

const LogContextKeyStr = ctxKey("LCK")
const UnnamedKeyStr = "unnamed"
const GinCtxRequestIdKeyStr = "x-req-id"
const GinCtxTraceIdKeyStr = "x-tra-id"
const CtxRequestId = "reqId"
const CtxTraceId = "traId"
const ArrayLimit = 100

func CopyLogContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.Background()
	}
	return context.WithValue(ctx, LogContextKeyStr, getContextLogValues(ctx).Copy())
}

func GetLogContext(ctx context.Context) []any {
	if ctx == nil {
		ctx = context.Background()
	}
	lv := getContextLogValues(ctx)
	return lv.All()
}

func SetLogContextKeyValue(ctx context.Context, key string, value any) context.Context {
	lv := getContextLogValues(ctx)
	if lv == nil {
		lv = &LogValues{}
		val := ctx.Value(LogContextKeyStr)
		if val != nil {
			lv.Set(UnnamedKeyStr, val)
		}
		lv.Set(key, value)
		ctx = context.WithValue(ctx, LogContextKeyStr, lv)
		return ctx
	}
	lv.Set(key, value)
	return ctx
}

func SetLogContextKeyStringValueIfNotEmpty(ctx context.Context, key string, value string) context.Context {
	if value != "" {
		return SetLogContextKeyValue(ctx, key, value)
	}
	return ctx
}

func GetLogContextValue(ctx context.Context, key string) (value any, ok bool) {
	lv := getContextLogValues(ctx)
	return lv.Get(key)
}

func GetLogContextValueAsString(ctx context.Context, key string) (value string, ok bool) {
	lv := getContextLogValues(ctx)
	return lv.GetStr(key)
}

func getContextLogValues(ctx context.Context) *LogValues {
	val := ctx.Value(LogContextKeyStr)
	lv, ok := val.(*LogValues)
	if !ok {
		return nil
	}
	return lv
}

func NewTrackLogContext(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = SetLogContextKeyValue(ctx, CtxRequestId, NewId())
	ctx = SetLogContextKeyValue(ctx, CtxTraceId, NewId())
	return ctx
}

func SetTrackLogContext(ctx context.Context, reqId string, traId string) context.Context {
	if reqId == "" {
		reqId = NewId()
	}
	if traId == "" {
		traId = NewId()
	}
	ctx = SetLogContextKeyValue(ctx, CtxRequestId, reqId)
	ctx = SetLogContextKeyValue(ctx, CtxTraceId, traId)
	return ctx
}

func GetLogTrackContext(ctx context.Context) (reqId string, traId string) {
	reqId, _ = GetLogContextValueAsString(ctx, CtxRequestId)
	traId, _ = GetLogContextValueAsString(ctx, CtxTraceId)
	if reqId == "" {
		reqId = NewId()
	}
	if traId == "" {
		traId = NewId()
	}
	return
}

func NewId() string {
	return uuid.New().String()[:13]
}
