package gin_logger

import (
	"context"
	"net/http/httptest"
	"slog/log_context"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGinContextLogger(t *testing.T) {
	// Create a new Gin context
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = httptest.NewRequest("GET", "/", nil)

	// Call the GinContextLogger function
	c.Request.Header.Set(log_context.GinCtxRequestIdKeyStr, "123")
	c.Request.Header.Set(log_context.GinCtxTraceIdKeyStr, "456")
	// ctx := GetGinTraceCtx(context.Background(), c, GinCtxRequestIdKeyStr, GinCtxTraceIdKeyStr)
	ctx := GetGinTraceCtx(context.Background(), c)

	result := log_context.GetLogContext(ctx)

	if result[0] != "reqId" {
		t.Errorf("Expected 'reqId', got %s", result[0])
	}
	if result[1] != "123" {
		t.Errorf("Expected '123', got %s", result[1])
	}
	if result[2] != "traId" {
		t.Errorf("Expected 'traId', got %s", result[2])
	}
	if result[3] != "456" {
		t.Errorf("Expected '456', got %s", result[3])
	}
}
