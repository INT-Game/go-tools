package loggers

import (
	"go.uber.org/zap"
	"xorm.io/xorm/log"
)

type xormZapSugaredLogger struct {
	logger  *zap.SugaredLogger
	showSQL bool
	level   log.LogLevel
}

func NewXormZapSugaredLogger(logger *zap.SugaredLogger) *xormZapSugaredLogger {
	return &xormZapSugaredLogger{
		logger: logger.WithOptions(zap.AddCallerSkip(2)), // xorm also wraps the logger, so we need to skip 2 callers
	}
}

func (l *xormZapSugaredLogger) Debug(v ...interface{}) {
	l.logger.Debug(v...)
}

func (l *xormZapSugaredLogger) Debugf(format string, v ...interface{}) {
	l.logger.Debugf(format, v...)
}

func (l *xormZapSugaredLogger) Error(v ...interface{}) {
	l.logger.Error(v...)
}

func (l *xormZapSugaredLogger) Errorf(format string, v ...interface{}) {
	l.logger.Errorf(format, v...)
}

func (l *xormZapSugaredLogger) Info(v ...interface{}) {
	l.logger.Info(v...)
}

func (l *xormZapSugaredLogger) Infof(format string, v ...interface{}) {
	l.logger.Infof(format, v...)
}

func (l *xormZapSugaredLogger) Warn(v ...interface{}) {
	l.logger.Warn(v...)
}

func (l *xormZapSugaredLogger) Warnf(format string, v ...interface{}) {
	l.logger.Warnf(format, v...)
}

func (l *xormZapSugaredLogger) Level() log.LogLevel {
	return l.level
}

func (l *xormZapSugaredLogger) SetLevel(level log.LogLevel) {
	l.level = level
}

func (l *xormZapSugaredLogger) ShowSQL(show ...bool) {
	if len(show) == 0 {
		l.showSQL = true
		return
	}
	l.showSQL = show[0]
}

func (l *xormZapSugaredLogger) IsShowSQL() bool {
	return l.showSQL
}
