package slog

// DebugLevel Level = iota - 1
// // InfoLevel is the default logging priority.
// InfoLevel = 0
// // WarnLevel logs are more important than Info, but don't need individual
// // human review.
// WarnLevel = 1
// // ErrorLevel logs are high-priority. If an application is running smoothly,
// // it shouldn't generate any error-level logs.
// ErrorLevel = 2
// // DPanicLevel logs are particularly important errors. In development the
// // logger panics after writing the message.
// DPanicLevel = 3
// // PanicLevel logs a message, then panics.
// PanicLevel = 4
// // FatalLevel logs a message, then calls os.Exit(1).
// FatalLevel = 5

type LogConfig struct {
	Name         string        `mapstructure:"name"`  // logger name, can be service name
	Level        int           `mapstructure:"level"` // zapcore/level.go
	Dir          string        `mapstructure:"dir"`
	Console      bool          `mapstructure:"console"`
	File         bool          `mapstructure:"file"`
	RotateConfig *RotateConfig `mapstructure:"rotate"`
	DebugRotate  *RotateConfig `mapstructure:"debug_rotate"`  // using rotate config, if nil. Debug log will not be collected by log service
	OutputRotate *RotateConfig `mapstructure:"output_rotate"` // using rotate config, if nil
	ErrorRotate  *RotateConfig `mapstructure:"error_rotate"`  // using rotate config, if nil
}

type RotateConfig struct {
	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSize int `mapstructure:"max_size"`
	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	MaxAge int `mapstructure:"max_age"`
	// MaxBackups is the maximum number of old log files to retain.  The default
	// is to retain all old log files (though MaxAge may still cause them to get
	// deleted.)
	MaxBackups int `mapstructure:"max_backups"`
	// Compress determines if the rotated log files should be compressed
	// using gzip. The default is not to perform compression.
	Compress bool `mapstructure:"compress"`
}

const (
	DebugLogFile      = "debug.log"
	OutputLogFile     = "output.log"
	ErrorLogFile      = "error.log"
	DefaultLoggerName = "slog"
)

var defaultRotateConfig = &RotateConfig{
	MaxSize:    200, // megabytes
	MaxBackups: 0,
	MaxAge:     7, // days
	Compress:   false,
}

// init rotate config
func GetRotateConfigs(config *LogConfig) (debugConfig *RotateConfig, outputConfig *RotateConfig, errorConfig *RotateConfig) {
	if config.RotateConfig == nil {
		config.RotateConfig = defaultRotateConfig
	}
	debugConfig = getRotateConfig(config.DebugRotate, config.RotateConfig)
	outputConfig = getRotateConfig(config.OutputRotate, config.RotateConfig)
	errorConfig = getRotateConfig(config.ErrorRotate, config.RotateConfig)
	return
}

func getRotateConfig(rotateConfig *RotateConfig, defaultConfig *RotateConfig) *RotateConfig {
	if rotateConfig == nil {
		rotateConfig = defaultConfig
	}
	return rotateConfig
}
