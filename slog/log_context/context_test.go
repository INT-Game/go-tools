package log_context

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyLogContext(t *testing.T) {
	// Test nil context
	ctx := CopyLogContext(nil)
	assert.NotNil(t, ctx)

	// Test context with values
	origCtx := context.Background()
	origCtx = SetLogContextKeyValue(origCtx, "key1", "value1")
	copiedCtx := CopyLogContext(origCtx)

	val, ok := GetLogContextValue(copiedCtx, "key1")
	assert.True(t, ok)
	assert.Equal(t, "value1", val)
}

func TestGetLogContext(t *testing.T) {
	// Test nil context
	values := GetLogContext(nil)
	assert.NotNil(t, values)
	assert.Empty(t, values)

	// Test context with values
	ctx := context.Background()
	ctx = SetLogContextKeyValue(ctx, "key1", "value1")
	values = GetLogContext(ctx)
	assert.Equal(t, []any{"key1", "value1"}, values)
}

func TestSetLogContextKeyValue(t *testing.T) {
	ctx := context.Background()

	// Test setting new value
	ctx = SetLogContextKeyValue(ctx, "key1", "value1")
	val, ok := GetLogContextValue(ctx, "key1")
	assert.True(t, ok)
	assert.Equal(t, "value1", val)

	// Test overwriting existing value
	ctx = SetLogContextKeyValue(ctx, "key1", "value2")
	val, ok = GetLogContextValue(ctx, "key1")
	assert.True(t, ok)
	assert.Equal(t, "value2", val)

	// Test unnamed key handling
	ctx = context.WithValue(context.Background(), LogContextKeyStr, "unnamed_value")
	ctx = SetLogContextKeyValue(ctx, "key1", "value1")
	val, ok = GetLogContextValue(ctx, UnnamedKeyStr)
	assert.True(t, ok)
	assert.Equal(t, "unnamed_value", val)
}

func TestSetLogContextKeyStringValueIfNotEmpty(t *testing.T) {
	ctx := context.Background()

	// Test with empty string
	newCtx := SetLogContextKeyStringValueIfNotEmpty(ctx, "key1", "")
	_, ok := GetLogContextValue(newCtx, "key1")
	assert.False(t, ok)

	// Test with non-empty string
	newCtx = SetLogContextKeyStringValueIfNotEmpty(ctx, "key1", "value1")
	val, ok := GetLogContextValue(newCtx, "key1")
	assert.True(t, ok)
	assert.Equal(t, "value1", val)
}

func TestGetLogContextValue(t *testing.T) {
	ctx := context.Background()
	ctx = SetLogContextKeyValue(ctx, "key1", "value1")

	// Test existing key
	val, ok := GetLogContextValue(ctx, "key1")
	assert.True(t, ok)
	assert.Equal(t, "value1", val)

	// Test non-existing key
	val, ok = GetLogContextValue(ctx, "nonexistent")
	assert.False(t, ok)
	assert.Nil(t, val)
}

func TestGetLogContextValueAsString(t *testing.T) {
	ctx := context.Background()
	ctx = SetLogContextKeyValue(ctx, "strKey", "string")
	ctx = SetLogContextKeyValue(ctx, "intKey", 42)

	// Test string value
	str, ok := GetLogContextValueAsString(ctx, "strKey")
	assert.True(t, ok)
	assert.Equal(t, "string", str)

	// Test non-string value
	str, ok = GetLogContextValueAsString(ctx, "intKey")
	assert.False(t, ok)
	assert.Empty(t, str)
}

func TestNewTrackLogContext(t *testing.T) {
	ctx := NewTrackLogContext(nil)

	reqId, ok := GetLogContextValueAsString(ctx, CtxRequestId)
	assert.True(t, ok)
	assert.Len(t, reqId, 13)

	traId, ok := GetLogContextValueAsString(ctx, CtxTraceId)
	assert.True(t, ok)
	assert.Len(t, traId, 13)
}

func TestSetTrackLogContext(t *testing.T) {
	ctx := context.Background()

	// Test with provided IDs
	ctx = SetTrackLogContext(ctx, "req123", "tra456")
	reqId, ok := GetLogContextValueAsString(ctx, CtxRequestId)
	assert.True(t, ok)
	assert.Equal(t, "req123", reqId)

	traId, ok := GetLogContextValueAsString(ctx, CtxTraceId)
	assert.True(t, ok)
	assert.Equal(t, "tra456", traId)

	// Test with empty IDs
	ctx = SetTrackLogContext(ctx, "", "")
	reqId, ok = GetLogContextValueAsString(ctx, CtxRequestId)
	assert.True(t, ok)
	assert.Len(t, reqId, 13)

	traId, ok = GetLogContextValueAsString(ctx, CtxTraceId)
	assert.True(t, ok)
	assert.Len(t, traId, 13)
}

func TestGetLogTrackContext(t *testing.T) {
	// Test with existing IDs
	ctx := SetTrackLogContext(context.Background(), "req123", "tra456")
	reqId, traId := GetLogTrackContext(ctx)
	assert.Equal(t, "req123", reqId)
	assert.Equal(t, "tra456", traId)

	// Test with empty context
	reqId, traId = GetLogTrackContext(context.Background())
	assert.Len(t, reqId, 13)
	assert.Len(t, traId, 13)
}

func TestNewId(t *testing.T) {
	// Test ID generation
	id1 := NewId()
	id2 := NewId()

	assert.Len(t, id1, 13)
	assert.Len(t, id2, 13)
	assert.NotEqual(t, id1, id2) // IDs should be unique
}
