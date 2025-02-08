package slog

import (
	"fmt"
	"os"
	"path"

	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var hostname string

func InitEncoders() (jsonEncoder zapcore.Encoder, consoleEncoder zapcore.Encoder) {
	hostname, _ = os.Hostname()
	jsonEncoder = GetJsonEncoder(hostname)
	consoleEncoder = GetConsoleEncoder(hostname)
	return
}

func GetJsonEncoder(hostname string) zapcore.Encoder {
	config := zapcore.EncoderConfig{
		TimeKey:        "dt",
		LevelKey:       "lv",
		NameKey:        "name",
		CallerKey:      "cal",
		MessageKey:     "msg",
		FunctionKey:    "",
		StacktraceKey:  "st",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	encoder := zapcore.NewJSONEncoder(config)
	encoder.AddString("host", hostname)
	return encoder
}

func GetConsoleEncoder(hostname string) zapcore.Encoder {
	config := zapcore.EncoderConfig{
		TimeKey:          "dt",
		LevelKey:         "lv",
		NameKey:          "name",
		CallerKey:        "cal",
		MessageKey:       "msg",
		FunctionKey:      "",
		StacktraceKey:    "st",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      levelEncoderWithHostname,
		EncodeTime:       zapcore.RFC3339TimeEncoder,
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		EncodeCaller:     CallerEncoder,
		EncodeName:       NameEncoder,
		ConsoleSeparator: " ",
	}
	encoder := zapcore.NewConsoleEncoder(config)
	return encoder
}

func levelEncoderWithHostname(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("[%s]", level.CapitalString()))
}

func CallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("[%s]:", caller.TrimmedPath()))
}

func NameEncoder(loggerName string, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("%s@%s", loggerName, hostname))
}

func GetFileSyncer(rotateConfig *RotateConfig, dir string, filename string) (zapcore.WriteSyncer, func() (err error)) {
	file := lumberjack.Logger{
		Filename:   path.Join(dir, filename),
		MaxSize:    rotateConfig.MaxSize, // megabytes
		MaxBackups: rotateConfig.MaxBackups,
		MaxAge:     rotateConfig.MaxAge, // days
		LocalTime:  true,
		Compress:   rotateConfig.Compress,
	}
	syncer := zapcore.AddSync(&file)
	return syncer, file.Close
}
