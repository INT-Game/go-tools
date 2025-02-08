package slog

import (
	"context"
	"github.com/INT-Game/go-tools/slog/log_context"
	"github.com/INT-Game/go-tools/slog/loggers"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger
var ZapLogger *zap.Logger

var debugCloseFunc func() (err error) = func() (err error) { return }
var outputCloseFunc func() (err error) = func() (err error) { return }
var errorCloseFunc func() (err error) = func() (err error) { return }
var fileDebugSyncer zapcore.WriteSyncer
var fileOutputSyncer zapcore.WriteSyncer
var fileErrorSyncer zapcore.WriteSyncer

func Init(config LogConfig) {
	// Config stderr and stdout files
	priorityDebug := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zapcore.DebugLevel
	})
	priorityOutput := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return zapcore.DebugLevel < lvl && lvl < zapcore.ErrorLevel
	})
	priorityError := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	// High-priority output should also go to standard error, and low-priority
	// output should also go to standard out.
	consoleStdout := zapcore.Lock(os.Stdout)
	consoleStderr := zapcore.Lock(os.Stderr)

	err := os.MkdirAll(config.Dir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// init rotate config & write syncer
	debugConfig, outputConfig, errorConfig := GetRotateConfigs(&config)
	fileDebugSyncer, debugCloseFunc = GetFileSyncer(debugConfig, config.Dir, DebugLogFile)
	fileOutputSyncer, outputCloseFunc = GetFileSyncer(outputConfig, config.Dir, OutputLogFile)
	fileErrorSyncer, errorCloseFunc = GetFileSyncer(errorConfig, config.Dir, ErrorLogFile)

	// Get encoders and their configs
	jsonEncoder, consoleEncoder := InitEncoders()

	zapcores := []zapcore.Core{}
	if config.File {
		zapcores = append(zapcores, zapcore.NewCore(jsonEncoder, fileDebugSyncer, priorityDebug))
		zapcores = append(zapcores, zapcore.NewCore(jsonEncoder, fileOutputSyncer, priorityOutput))
		zapcores = append(zapcores, zapcore.NewCore(jsonEncoder, fileErrorSyncer, priorityError))
	}
	if config.Console {
		zapcores = append(zapcores, zapcore.NewCore(consoleEncoder, consoleStdout, priorityDebug))
		zapcores = append(zapcores, zapcore.NewCore(consoleEncoder, consoleStdout, priorityOutput))
		zapcores = append(zapcores, zapcore.NewCore(consoleEncoder, consoleStderr, priorityError))
	}
	core := zapcore.NewTee(zapcores...)

	if config.Name == "" {
		config.Name = DefaultLoggerName
	}
	ZapLogger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)).Named(config.Name)
	loggers.Logger_2 = ZapLogger.WithOptions(zap.AddCallerSkip(2)).Sugar()
	Logger = ZapLogger.Sugar()
}

func Close() {
	if Logger != nil {
		err := Logger.Sync()
		if err != nil {
			loggers.DefaultPrintln(err.Error())
		}
		Logger = nil
		loggers.Logger_2 = nil
		ZapLogger = nil
	}
	loggers.UsingDefaultLogger()
	err := debugCloseFunc()
	if err != nil {
		loggers.DefaultPrintln(err.Error())
	}

	err = outputCloseFunc()
	if err != nil {
		loggers.DefaultPrintln(err.Error())
	}
	err = errorCloseFunc()
	if err != nil {
		loggers.DefaultPrintln(err.Error())
	}
}

var CLog = loggers.CLog
var CLogln = loggers.CLogln
var CLogw = loggers.CLogw
var CDebug = loggers.CDebug
var CDebugln = loggers.CDebugln
var CDebugw = loggers.CDebugw
var CInfo = loggers.CInfo
var CInfoln = loggers.CInfoln
var CInfow = loggers.CInfow
var CWarn = loggers.CWarn
var CWarnln = loggers.CWarnln
var CWarnw = loggers.CWarnw
var CError = loggers.CError
var CErrorln = loggers.CErrorln
var CErrorw = loggers.CErrorw
var CDPanic = loggers.CDPanic
var CDPanicln = loggers.CDPanicln
var CDPanicw = loggers.CDPanicw
var CPanic = loggers.CPanic
var CPanicln = loggers.CPanicln
var CPanicw = loggers.CPanicw
var CFatal = loggers.CFatal
var CFatalln = loggers.CFatalln
var CFatalw = loggers.CFatalw
var GetLogContext = log_context.GetLogContext
var SetContextKeyValue = log_context.SetLogContextKeyValue

func GetContextLogger(ctx context.Context) *zap.SugaredLogger {
	kvs := log_context.GetLogContext(ctx)
	if Logger != nil {
		return Logger.With(kvs...)
	}
	return Logger
}

var NewSLogger = loggers.NewSLogger
