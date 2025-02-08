package slog

import (
	"testing"
)

func TestGetRotateConfigs(t *testing.T) {
	t.Run("All rotateConfigs are set", func(t *testing.T) {
		// Create a sample LogConfig
		config := &LogConfig{
			RotateConfig: &RotateConfig{
				MaxSize:    100,
				MaxAge:     30,
				MaxBackups: 5,
				Compress:   true,
			},
			DebugRotate: &RotateConfig{
				MaxSize:    50,
				MaxAge:     7,
				MaxBackups: 3,
				Compress:   false,
			},
			OutputRotate: &RotateConfig{
				MaxSize:    150,
				MaxAge:     14,
				MaxBackups: 2,
				Compress:   true,
			},
			ErrorRotate: &RotateConfig{
				MaxSize:    80,
				MaxAge:     21,
				MaxBackups: 4,
				Compress:   false,
			},
		}

		// Call the function
		debugConfig, outputConfig, errorConfig := GetRotateConfigs(config)

		// Assert the returned values
		if config.DebugRotate != debugConfig {
			t.Errorf("Expected debugConfig to be %v, but got %v", config.DebugRotate, debugConfig)
		}

		if config.OutputRotate != outputConfig {
			t.Errorf("Expected outputConfig to be %v, but got %v", config.OutputRotate, outputConfig)
		}

		if config.ErrorRotate != errorConfig {
			t.Errorf("Expected errorConfig to be %v, but got %v", config.ErrorRotate, errorConfig)
		}
	})

	t.Run("All rotateConfigs are nil", func(t *testing.T) {
		// Create a sample LogConfig
		config := &LogConfig{}

		// Call the function
		debugConfig, outputConfig, errorConfig := GetRotateConfigs(config)

		// Assert the returned values
		if debugConfig != defaultRotateConfig {
			t.Errorf("Expected debugConfig to be %v, but got %v", defaultRotateConfig, debugConfig)
		}

		if outputConfig != defaultRotateConfig {
			t.Errorf("Expected outputConfig to be %v, but got %v", defaultRotateConfig, outputConfig)
		}

		if errorConfig != defaultRotateConfig {
			t.Errorf("Expected errorConfig to be %v, but got %v", defaultRotateConfig, errorConfig)
		}
	})

	t.Run("Basic rotateConfig is set", func(t *testing.T) {
		// Create a sample LogConfig
		config := &LogConfig{
			RotateConfig: &RotateConfig{
				MaxSize:    100,
				MaxAge:     30,
				MaxBackups: 5,
				Compress:   true,
			},
		}

		// Call the function
		debugConfig, outputConfig, errorConfig := GetRotateConfigs(config)

		// Assert the returned values
		if config.RotateConfig != debugConfig {
			t.Errorf("Expected debugConfig to be %v, but got %v", config.RotateConfig, debugConfig)
		}

		if config.RotateConfig != outputConfig {
			t.Errorf("Expected outputConfig to be %v, but got %v", config.RotateConfig, outputConfig)
		}

		if config.RotateConfig != errorConfig {
			t.Errorf("Expected errorConfig to be %v, but got %v", config.RotateConfig, errorConfig)
		}
	})
}
