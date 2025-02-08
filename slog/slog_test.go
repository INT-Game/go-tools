package slog

import (
	"context"
	"github.com/INT-Game/go-tools/slog/log_context"
	"os"
	"path"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInitSlog(t *testing.T) {
	logDir := "./tmp/test_slog"
	// Create a test LogConfig
	config := LogConfig{
		Dir:     logDir,
		File:    true,
		Console: true,
	}

	// Call the InitSlog function
	Init(config)
	logger := Logger
	// Perform assertions on the logger and its output

	// Example assertion: Check if logger is not nil
	if logger == nil {
		t.Errorf("Expected logger to be initialized, but got nil")
	}

	logger.Info("test info")
	logger.Error("test error")

	// Example assertion: Check if log files are created
	_, err := os.Stat(path.Join(logDir, OutputLogFile))
	if err != nil {
		t.Errorf("Expected output.log file to be created, but got error: %v", err)
	}

	_, err = os.Stat(path.Join(logDir, ErrorLogFile))
	if err != nil {
		t.Errorf("Expected errors.log file to be created, but got error: %v", err)
	}

	// Example assertion: Check if loggers write to the correct outputs
	outputBytes, err := os.ReadFile(path.Join(logDir, OutputLogFile))
	if err != nil {
		t.Errorf("Failed to read output.log file: %v", err)
	}
	outputContent := string(outputBytes)
	if !strings.Contains(outputContent, "test info") {
		t.Errorf("Expected 'test info' in output.log, but got: %s", outputContent)
	}

	stderrBytes, err := os.ReadFile(path.Join(logDir, ErrorLogFile))
	if err != nil {
		t.Errorf("Failed to read error.log file: %v", err)
	}
	errorContent := string(stderrBytes)
	if !strings.Contains(errorContent, "test error") {
		t.Errorf("Expected 'test error' in error.log, but got: %s", errorContent)
	}

	// Close the logger
	Close()
	err = os.RemoveAll("./tmp")
	if err != nil {
		t.Errorf("Failed to remove temporary directory: %v", err)
	}

	// test panic
	assert.Panics(t, func() { logger.Panic("test panic") })
	assert.Panics(t, func() { logger.Panicln("test panicln") })
	assert.Panics(t, func() { logger.Panicw("test panicw", "key", "value") })
}

func TestContextLogger(t *testing.T) {
	fields := map[string]any{"tid": "123", "rid": "456"}
	t.Run("Test ContextLogger without ZapLogger", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), log_context.LogContextKeyStr, fields)
		CDebugw(ctx, "this is a test msg", "0", "1")
		ctx = SetContextKeyValue(ctx, "uid", "789")
		CDebugw(ctx, "this is a test msg", "0", "1")
		CDebug(ctx, "this is a test msg")
		CDebug(ctx, "this is a test msg witha args: %s %s", "1", "2")
	})

	t.Run("Test ContextLogger with ZapLogger", func(t *testing.T) {
		Init(LogConfig{
			Dir:     "./tmpCtx",
			Console: true,
			File:    true,
		})

		ctx := context.WithValue(context.Background(), log_context.LogContextKeyStr, fields)
		CDebugw(ctx, "this is a test msg", "0", "1")
		ctx = SetContextKeyValue(ctx, "uid", "789")
		CDebugw(ctx, "this is a test msg", "0", "1")
		CDebug(ctx, "this is a test msg")
		CDebug(ctx, "this is a test msg witha args: %s %s", "1", "2")
		CDebugln(ctx, "this is a test msg witha args", "1", "2")
		Close()
		err := os.RemoveAll("./tmpCtx")
		if err != nil {
			t.Errorf("Failed to remove temporary directory: %v", err)
		}
	})
}

func TestLogRotate(t *testing.T) {
	Init(LogConfig{
		Dir:  "./tmpRotate",
		File: true,
		RotateConfig: &RotateConfig{
			MaxSize:    1,
			MaxAge:     1,
			MaxBackups: 5,
			Compress:   true,
		},
	})
	for i := 0; i < 20000; i++ {
		CInfo(context.Background(), "[%d] this is a test msg with many words to test log rotate. this is a test msg with many words to test log rotate.this is a test msg with many words to test log rotate.this is a test msg with many words to test log rotate.this is a test msg with many words to test log rotate.this is a test msg with many words to test log rotate.this is a test msg with many words to test log rotate.this is a test msg with many words to test log rotate.this is a test msg with many words to test log rotate.", i)
	}
	// wait few seconds for log rotate
	time.Sleep(5 * time.Second)
	// count files
	files, err := os.ReadDir("./tmpRotate")
	if err != nil {
		t.Errorf("Failed to read directory: %v", err)
	}
	if len(files) > 6 {
		t.Errorf("Expected 6 files, but got %d", len(files))
	}
	Close()
	err = os.RemoveAll("./tmpRotate")
	if err != nil {
		t.Errorf("Failed to remove temporary directory: %v", err)
	}
}
