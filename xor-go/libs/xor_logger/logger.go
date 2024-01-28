package xor_logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// TODO support using sentry

func InitLogger(cfg *Config) (*zap.Logger, error) {
	env, err := ParseEnvironment(cfg.Environment)
	if err != nil {
		return nil, err
	}
	configZap, err := GetLoggerConfig(env)
	if err != nil {
		return nil, err
	}

	var level zapcore.Level
	err = level.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshall xor_logger level: %v", err)
	}

	// TODO support all options from original zap.Config
	configZap.Level = zap.NewAtomicLevelAt(level)
	configZap.OutputPaths = cfg.OutputPaths
	configZap.ErrorOutputPaths = cfg.ErrorOutputPaths
	configZap.Encoding = cfg.Encoding
	fmt.Println(cfg.ErrorOutputPaths)

	logger, err := configZap.Build()
	if err != nil {
		return nil, fmt.Errorf("xor_logger build failed: %v", err)
	}

	return logger, nil
}

func GetLoggerConfig(env Environment) (*zap.Config, error) {
	var config zap.Config
	var err error = nil
	switch env {
	case DevEnvironment:
		config = zap.NewDevelopmentConfig()
	case ProdEnvironment:
		config = zap.NewProductionConfig()
	default:
		err = fmt.Errorf("cannot get xor_logger config for env: %s", env)
	}
	if err != nil {
		return nil, err
	}
	return &config, nil
}
