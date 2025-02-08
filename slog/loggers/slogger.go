package loggers

import (
	"context"
	"fmt"

	"go.uber.org/zap/zapcore"
)

const ()

type SLogger struct {
	Template      string
	KeysAndValues []any
}

func NewSLogger(template string, keysAndValues ...any) *SLogger {
	return &SLogger{template, keysAndValues}
}

func (s *SLogger) With(key string, value interface{}) *SLogger {
	s.KeysAndValues = append(s.KeysAndValues, key, value)
	return s
}

func (s *SLogger) GetMsg(msg string) string {
	if s == nil || s.Template == "" {
		return msg
	}
	return fmt.Sprintf(s.Template, msg)
}

func (s *SLogger) GetKeysAndValues() []any {
	if s == nil {
		return nil
	}
	return s.KeysAndValues
}

func (s *SLogger) CLog(ctx context.Context, level zapcore.Level, extra_skip int, template string, args ...interface{}) {
	logWithLevelAndContext(ctx, extra_skip, level, s.GetMsg(fmt.Sprintf(template, args...)), s.GetKeysAndValues()...)
}

func (s *SLogger) CLogln(ctx context.Context, level zapcore.Level, extra_skip int, args ...interface{}) {
	logWithLevelAndContext(ctx, extra_skip, level, s.GetMsg(fmt.Sprint(args...)), s.GetKeysAndValues()...)
}

func (s *SLogger) CLogw(ctx context.Context, level zapcore.Level, extra_skip int, msg string, keysAndValues ...interface{}) {
	kvs := append(s.GetKeysAndValues(), keysAndValues...)
	logWithLevelAndContext(ctx, extra_skip, level, s.GetMsg(msg), kvs...)
}

func (s *SLogger) CDebug(ctx context.Context, template string, args ...interface{}) {
	logWithLevelAndContext(ctx, 0, zapcore.DebugLevel, s.GetMsg(fmt.Sprintf(template, args...)), s.GetKeysAndValues()...)
}

func (s *SLogger) CInfo(ctx context.Context, template string, args ...interface{}) {
	logWithLevelAndContext(ctx, 0, zapcore.InfoLevel, s.GetMsg(fmt.Sprintf(template, args...)), s.GetKeysAndValues()...)
}

func (s *SLogger) CInfoln(ctx context.Context, args ...interface{}) {
	logWithLevelAndContext(ctx, 0, zapcore.InfoLevel, s.GetMsg(fmt.Sprint(args...)), s.GetKeysAndValues()...)
}

func (s *SLogger) CInfow(ctx context.Context, msg string, keysAndValues ...interface{}) {
	kvs := append(s.GetKeysAndValues(), keysAndValues...)
	logWithLevelAndContext(ctx, 0, zapcore.InfoLevel, s.GetMsg(msg), kvs...)
}

func (s *SLogger) CWarn(ctx context.Context, template string, args ...interface{}) {
	logWithLevelAndContext(ctx, 0, zapcore.WarnLevel, s.GetMsg(fmt.Sprintf(template, args...)), s.GetKeysAndValues()...)
}

func (s *SLogger) CWarnln(ctx context.Context, args ...interface{}) {
	logWithLevelAndContext(ctx, 0, zapcore.WarnLevel, s.GetMsg(fmt.Sprint(args...)), s.GetKeysAndValues()...)
}

func (s *SLogger) CWarnw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	kvs := append(s.GetKeysAndValues(), keysAndValues...)
	logWithLevelAndContext(ctx, 0, zapcore.WarnLevel, s.GetMsg(msg), kvs...)
}

func (s *SLogger) CError(ctx context.Context, template string, args ...interface{}) {
	logWithLevelAndContext(ctx, 0, zapcore.ErrorLevel, s.GetMsg(fmt.Sprintf(template, args...)), s.GetKeysAndValues()...)
}

func (s *SLogger) CErrorln(ctx context.Context, args ...interface{}) {
	logWithLevelAndContext(ctx, 0, zapcore.ErrorLevel, s.GetMsg(fmt.Sprint(args...)), s.GetKeysAndValues()...)
}

func (s *SLogger) CErrorw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	kvs := append(s.GetKeysAndValues(), keysAndValues...)
	logWithLevelAndContext(ctx, 0, zapcore.ErrorLevel, s.GetMsg(msg), kvs...)
}

func (s *SLogger) CDPanic(ctx context.Context, template string, args ...interface{}) {
	logWithLevelAndContext(ctx, 0, zapcore.DPanicLevel, s.GetMsg(fmt.Sprintf(template, args...)), s.GetKeysAndValues()...)
}

func (s *SLogger) CDPanicln(ctx context.Context, args ...interface{}) {
	logWithLevelAndContext(ctx, 0, zapcore.DPanicLevel, s.GetMsg(fmt.Sprint(args...)), s.GetKeysAndValues()...)
}

func (s *SLogger) CDPanicw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	kvs := append(s.GetKeysAndValues(), keysAndValues...)
	logWithLevelAndContext(ctx, 0, zapcore.DPanicLevel, s.GetMsg(msg), kvs...)
}

func (s *SLogger) CPanic(ctx context.Context, template string, args ...interface{}) {
	logWithLevelAndContext(ctx, 0, zapcore.PanicLevel, s.GetMsg(fmt.Sprintf(template, args...)), s.GetKeysAndValues()...)
}

func (s *SLogger) CPanicln(ctx context.Context, args ...interface{}) {
	logWithLevelAndContext(ctx, 0, zapcore.PanicLevel, s.GetMsg(fmt.Sprint(args...)), s.GetKeysAndValues()...)
}

func (s *SLogger) CPanicw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	kvs := append(s.GetKeysAndValues(), keysAndValues...)
	logWithLevelAndContext(ctx, 0, zapcore.PanicLevel, s.GetMsg(msg), kvs...)
}

func (s *SLogger) CFatal(ctx context.Context, template string, args ...interface{}) {
	logWithLevelAndContext(ctx, 0, zapcore.FatalLevel, s.GetMsg(fmt.Sprintf(template, args...)), s.GetKeysAndValues()...)
}

func (s *SLogger) CFatalln(ctx context.Context, args ...interface{}) {
	logWithLevelAndContext(ctx, 0, zapcore.FatalLevel, s.GetMsg(fmt.Sprint(args...)), s.GetKeysAndValues()...)
}

func (s *SLogger) CFatalw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	kvs := append(s.GetKeysAndValues(), keysAndValues...)
	logWithLevelAndContext(ctx, 0, zapcore.FatalLevel, s.GetMsg(msg), kvs...)
}
