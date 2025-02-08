package loggers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func setupTestLogger(_ *testing.T) (*observer.ObservedLogs, *zap.SugaredLogger) {
	core, recorded := observer.New(zapcore.DebugLevel)
	logger := zap.New(core).Sugar()
	Logger_2 = logger
	return recorded, logger
}

func TestLogWithLevelAndContext(t *testing.T) {
	recorded, _ := setupTestLogger(t)

	tests := []struct {
		name          string
		level         zapcore.Level
		msg           string
		keysAndValues []interface{}
		expectedLogs  int
	}{
		{
			name:          "Debug level logging",
			level:         zapcore.DebugLevel,
			msg:           "debug message",
			keysAndValues: []interface{}{"key", "value"},
			expectedLogs:  1,
		},
		{
			name:          "Info level logging",
			level:         zapcore.InfoLevel,
			msg:           "info message",
			keysAndValues: []interface{}{"key", "value"},
			expectedLogs:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorded.TakeAll() // Clear previous logs before each test case
			ctx := context.Background()
			logWithLevelAndContext(ctx, 0, tt.level, tt.msg, tt.keysAndValues...)

			logs := recorded.All()
			assert.Equal(t, tt.expectedLogs, len(logs))
			if len(logs) > 0 {
				assert.Equal(t, tt.msg, logs[len(logs)-1].Message)
				assert.Equal(t, tt.level, logs[len(logs)-1].Level)
			}
		})
	}
}

func TestContextualLogging(t *testing.T) {
	recorded, _ := setupTestLogger(t)
	ctx := context.Background()

	tests := []struct {
		name     string
		logFunc  func()
		level    zapcore.Level
		message  string
		contains string
	}{
		{
			name: "CDebug logging",
			logFunc: func() {
				CDebug(ctx, "debug %s", "message")
			},
			level:    zapcore.DebugLevel,
			message:  "debug message",
			contains: "debug message",
		},
		{
			name: "CInfo logging",
			logFunc: func() {
				CInfo(ctx, "info %s", "message")
			},
			level:    zapcore.InfoLevel,
			message:  "info message",
			contains: "info message",
		},
		{
			name: "CWarn logging",
			logFunc: func() {
				CWarn(ctx, "warn %s", "message")
			},
			level:    zapcore.WarnLevel,
			message:  "warn message",
			contains: "warn message",
		},
		{
			name: "CError logging",
			logFunc: func() {
				CError(ctx, "error %s", "message")
			},
			level:    zapcore.ErrorLevel,
			message:  "error message",
			contains: "error message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorded.TakeAll() // Clear previous logs
			tt.logFunc()

			logs := recorded.All()
			assert.Equal(t, 1, len(logs))
			assert.Equal(t, tt.level, logs[0].Level)
			assert.Equal(t, tt.message, logs[0].Message)
		})
	}
}

func TestDefaultLoggerFallback(t *testing.T) {
	// Temporarily set Logger_2 to nil to test default logger
	originalLogger := Logger_2
	Logger_2 = nil
	defer func() {
		Logger_2 = originalLogger
	}()

	// Test panic recovery
	defer func() {
		if r := recover(); r != nil {
			// Expected panic
			assert.Contains(t, r, "panic message")
		}
	}()

	ctx := context.Background()

	// Test normal logging (should not panic)
	CInfo(ctx, "info message")
	CDebug(ctx, "debug message")
	CWarn(ctx, "warn message")
	CError(ctx, "error message")

	// Test panic level (should panic)
	CPanic(ctx, "panic %s", "message")
}

func TestUsingDefaultLogger(t *testing.T) {
	// Save original logging functions
	originalLog := Log
	originalLogw := Logw
	originalLogf := Logf
	originalLogln := Logln

	defer func() {
		// Restore original logging functions
		Log = originalLog
		Logw = originalLogw
		Logf = originalLogf
		Logln = originalLogln
	}()

	UsingDefaultLogger()

	// Test that logging functions are replaced with default implementations
	Log(zapcore.InfoLevel, "test message")
	Logw(zapcore.InfoLevel, "test message", "key", "value")
	Logf(zapcore.InfoLevel, "test %s", "message")
	Logln(zapcore.InfoLevel, "test message")
}
