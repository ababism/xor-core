package config

import XorLogger "xor-go/libs/xor_logger"

type Config struct {
	LoggerConfig LoggerConfig `yaml:"logger" env-prefix:"LOGGER"`
}

type LoggerConfig struct {
	Environment      string   `yaml:"environment" env:"ENVIRONMENT"`
	Level            string   `yaml:"level" env:"LEVEL"`
	Encoding         string   `yaml:"encoding" env:"ENCODING"`
	OutputPaths      []string `yaml:"output_paths" env:"OUTPUT_PATHS"`
	ErrorOutputPaths []string `yaml:"error_output_paths" env:"ERROR_OUTPUT_PATHS"`
}

func (c *LoggerConfig) ToXorLoggerConfig() *XorLogger.Config {
	return &XorLogger.Config{
		Environment:      c.Environment,
		Level:            c.Level,
		OutputPaths:      c.OutputPaths,
		ErrorOutputPaths: c.ErrorOutputPaths,
		Encoding:         c.Encoding,
	}
}
