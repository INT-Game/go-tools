package loggers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
)

func TestNewSLogger(t *testing.T) {
	tests := []struct {
		name          string
		template      string
		keysAndValues []any
		want          *SLogger
	}{
		{
			name:     "create with template only",
			template: "test-%s",
			want:     &SLogger{Template: "test-%s"},
		},
		{
			name:          "create with template and kv",
			template:      "test-%s",
			keysAndValues: []any{"key", "value"},
			want:          &SLogger{Template: "test-%s", KeysAndValues: []any{"key", "value"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSLogger(tt.template, tt.keysAndValues...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSLogger_With(t *testing.T) {
	logger := NewSLogger("test-%s")

	logger = logger.With("key1", "value1")
	assert.Equal(t, []any{"key1", "value1"}, logger.KeysAndValues)

	logger = logger.With("key2", "value2")
	assert.Equal(t, []any{"key1", "value1", "key2", "value2"}, logger.KeysAndValues)
}

func TestSLogger_GetMsg(t *testing.T) {
	tests := []struct {
		name     string
		logger   *SLogger
		msg      string
		expected string
	}{
		{
			name:     "nil logger",
			logger:   nil,
			msg:      "test message",
			expected: "test message",
		},
		{
			name:     "empty template",
			logger:   &SLogger{Template: ""},
			msg:      "test message",
			expected: "test message",
		},
		{
			name:     "with template",
			logger:   &SLogger{Template: "PREFIX: %s"},
			msg:      "test message",
			expected: "PREFIX: test message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.logger.GetMsg(tt.msg)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestSLogger_GetKeysAndValues(t *testing.T) {
	tests := []struct {
		name     string
		logger   *SLogger
		expected []any
	}{
		{
			name:     "nil logger",
			logger:   nil,
			expected: nil,
		},
		{
			name:     "empty kv",
			logger:   &SLogger{},
			expected: nil,
		},
		{
			name:     "with kv",
			logger:   &SLogger{KeysAndValues: []any{"key", "value"}},
			expected: []any{"key", "value"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.logger.GetKeysAndValues()
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestSLogger_LoggingMethods(t *testing.T) {
	// This is a basic smoke test to ensure logging methods don't panic
	ctx := context.Background()
	logger := NewSLogger("test-%s", "key", "value")

	// Test various logging methods
	logger.CDebug(ctx, "debug %s", "message")
	logger.CInfo(ctx, "info %s", "message")
	logger.CInfoln(ctx, "info", "message")
	logger.CInfow(ctx, "info message", "extra_key", "extra_value")
	logger.CWarn(ctx, "warn %s", "message")
	logger.CWarnln(ctx, "warn", "message")
	logger.CWarnw(ctx, "warn message", "extra_key", "extra_value")
	logger.CError(ctx, "error %s", "message")
	logger.CErrorln(ctx, "error", "message")
	logger.CErrorw(ctx, "error message", "extra_key", "extra_value")

	// Test the generic logging method
	logger.CLog(ctx, zapcore.InfoLevel, 0, "test %s", "message")
	logger.CLogln(ctx, zapcore.InfoLevel, 0, "test", "message")
	logger.CLogw(ctx, zapcore.InfoLevel, 0, "test message", "extra_key", "extra_value")

	// These should panic
	// DPanic only panic in development mode
	logger.CDPanic(ctx, "panic %s", "message")
	logger.CDPanicln(ctx, "panic", "message")
	logger.CDPanicw(ctx, "panic message", "extra_key", "extra_value")
	assert.Panics(t, func() { logger.CPanic(ctx, "panic %s", "message") })
	assert.Panics(t, func() { logger.CPanicln(ctx, "panic", "message") })
	assert.Panics(t, func() { logger.CPanicw(ctx, "panic message", "extra_key", "extra_value") })
	// Fatal will exit the program directly
	// logger.CFatal(ctx, "fatal %s", "message")
	// logger.CFatalln(ctx, "fatal", "message")
	// logger.CFatalw(ctx, "fatal message", "extra_key", "extra_value")

}

func TestNilSLogger_LoggingMethods(t *testing.T) {
	// Test that logging with nil logger doesn't panic
	var logger *SLogger
	ctx := context.Background()

	// These should not panic
	logger.CDebug(ctx, "debug %s", "message")
	logger.CInfo(ctx, "info %s", "message")
	logger.CInfoln(ctx, "info", "message")
	logger.CInfow(ctx, "info message", "extra_key", "extra_value")
	// These should panic
	assert.Panics(t, func() { logger.CPanic(ctx, "panic %s", "message") })
	assert.Panics(t, func() { logger.CPanicln(ctx, "panic", "message") })
	assert.Panics(t, func() { logger.CPanicw(ctx, "panic message", "extra_key", "extra_value") })
}
