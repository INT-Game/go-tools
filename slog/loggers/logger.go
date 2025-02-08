package loggers

import (
	"context"
	"fmt"
	"os"
	"slog/log_context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger_2 *zap.SugaredLogger // skipCaller(2) Sugared Logger

func DefaultPrint(args ...interface{})                   { fmt.Println(args...) }
func DefaultPrintf(template string, args ...interface{}) { fmt.Println(fmt.Sprintf(template, args...)) }
func DefaultPrintln(args ...interface{})                 { fmt.Println(args...) }
func DefaultPrintw(msg string, keysAndValues ...interface{}) {
	fmt.Printf("%s %v\n", msg, keysAndValues)
}
func DefaultError(args ...interface{}) { fmt.Fprintln(os.Stderr, args...) }
func DefaultErrorf(template string, args ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(template, args...))
}
func DefaultErrorln(args ...interface{}) { fmt.Fprintln(os.Stderr, args...) }
func DefaultErrorw(msg string, keysAndValues ...interface{}) {
	fmt.Fprintf(os.Stderr, "%s %v\n", msg, keysAndValues)
}

func DefaultPanic(args ...interface{})                   { panic(args) }
func DefaultPanicf(template string, args ...interface{}) { panic(fmt.Sprintf(template, args...)) }
func DefaultPanicln(args ...interface{})                 { panic(args) }
func DefaultPanicw(msg string, keysAndValues ...interface{}) {
	panic(fmt.Sprintf("%s %v", msg, keysAndValues))
}

var Log = func(lvl zapcore.Level, args ...interface{}) {
	DefaultPrint(args)
}
var Logw = func(lvl zapcore.Level, msg string, keysAndValues ...interface{}) {
	DefaultPrintw(msg, keysAndValues)
}
var Logf = func(lvl zapcore.Level, template string, args ...interface{}) {
	DefaultPrintf(template, args)
}
var Logln = func(lvl zapcore.Level, args ...interface{}) {
	DefaultPrintln(args)
}

func CLog(ctx context.Context, level zapcore.Level, extra_skip int, template string, args ...interface{}) {
	logWithLevelAndContext(ctx, extra_skip, level, fmt.Sprintf(template, args...))
}
func CLogln(ctx context.Context, level zapcore.Level, extra_skip int, args ...interface{}) {
	logWithLevelAndContext(ctx, extra_skip, level, fmt.Sprint(args...))
}
func CLogw(ctx context.Context, level zapcore.Level, extra_skip int, msg string, keysAndValues ...interface{}) {
	logWithLevelAndContext(ctx, extra_skip, level, msg, keysAndValues...)
}
func CDebug(ctx context.Context, template string, args ...interface{}) {
	logWithLevelAndContext(ctx, 0, zap.DebugLevel, fmt.Sprintf(template, args...))
}
func CDebugln(ctx context.Context, args ...interface{}) {
	logWithLevelAndContext(ctx, 0, zap.DebugLevel, fmt.Sprint(args...))
}
func CDebugw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	logWithLevelAndContext(ctx, 0, zap.DebugLevel, msg, keysAndValues...)
}
func CInfo(ctx context.Context, template string, args ...interface{}) {
	logWithLevelAndContext(ctx, 0, zap.InfoLevel, fmt.Sprintf(template, args...))
}
func CInfoln(ctx context.Context, args ...interface{}) {
	logWithLevelAndContext(ctx, 0, zap.InfoLevel, fmt.Sprint(args...))
}
func CInfow(ctx context.Context, msg string, keysAndValues ...interface{}) {
	logWithLevelAndContext(ctx, 0, zap.InfoLevel, msg, keysAndValues...)
}
func CWarn(ctx context.Context, template string, args ...interface{}) {
	logWithLevelAndContext(ctx, 0, zap.WarnLevel, fmt.Sprintf(template, args...))
}
func CWarnln(ctx context.Context, args ...interface{}) {
	logWithLevelAndContext(ctx, 0, zap.WarnLevel, fmt.Sprint(args...))
}
func CWarnw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	logWithLevelAndContext(ctx, 0, zap.WarnLevel, msg, keysAndValues...)
}
func CError(ctx context.Context, template string, args ...interface{}) {
	logWithLevelAndContext(ctx, 0, zap.ErrorLevel, fmt.Sprintf(template, args...))
}
func CErrorln(ctx context.Context, args ...interface{}) {
	logWithLevelAndContext(ctx, 0, zap.ErrorLevel, fmt.Sprint(args...))
}
func CErrorw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	logWithLevelAndContext(ctx, 0, zap.ErrorLevel, msg, keysAndValues...)
}
func CDPanic(ctx context.Context, template string, args ...interface{}) {
	logWithLevelAndContext(ctx, 0, zap.DPanicLevel, fmt.Sprintf(template, args...))
}
func CDPanicln(ctx context.Context, args ...interface{}) {
	logWithLevelAndContext(ctx, 0, zap.DPanicLevel, fmt.Sprint(args...))
}
func CDPanicw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	logWithLevelAndContext(ctx, 0, zap.DPanicLevel, msg, keysAndValues...)
}
func CPanic(ctx context.Context, template string, args ...interface{}) {
	logWithLevelAndContext(ctx, 0, zap.PanicLevel, fmt.Sprintf(template, args...))
}
func CPanicln(ctx context.Context, args ...interface{}) {
	logWithLevelAndContext(ctx, 0, zap.PanicLevel, fmt.Sprint(args...))
}
func CPanicw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	logWithLevelAndContext(ctx, 0, zap.PanicLevel, msg, keysAndValues...)
}
func CFatal(ctx context.Context, template string, args ...interface{}) {
	logWithLevelAndContext(ctx, 0, zap.FatalLevel, fmt.Sprintf(template, args...))
}
func CFatalln(ctx context.Context, args ...interface{}) {
	logWithLevelAndContext(ctx, 0, zap.FatalLevel, fmt.Sprint(args...))
}
func CFatalw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	logWithLevelAndContext(ctx, 0, zap.FatalLevel, msg, keysAndValues...)
}

func logWithLevelAndContext(ctx context.Context, extra_skip int, level zapcore.Level, msg string, keysAndValues ...interface{}) {
	if ctx == nil {
		ctx = context.Background()
	}
	kvs := log_context.GetLogContext(ctx)
	kvs = append(kvs, keysAndValues...)
	if Logger_2 == nil {
		if level == zapcore.PanicLevel {
			DefaultPanicw(msg, kvs...)
		} else {
			DefaultPrintw(msg, kvs...)
		}
	} else if extra_skip == 0 {
		Logger_2.Logw(level, msg, kvs...)
	} else {
		Logger_2.WithOptions(zap.AddCallerSkip(extra_skip)).Logw(level, msg, kvs...)
	}
}

func UsingDefaultLogger() {
	Log = func(_ zapcore.Level, args ...interface{}) {
		DefaultPrint(args)
	}
	Logw = func(_ zapcore.Level, msg string, keysAndValues ...interface{}) {
		DefaultPrintw(msg, keysAndValues)
	}
	Logf = func(_ zapcore.Level, template string, args ...interface{}) {
		DefaultPrintf(template, args)
	}
	Logln = func(_ zapcore.Level, args ...interface{}) {
		DefaultPrintln(args)
	}
}
