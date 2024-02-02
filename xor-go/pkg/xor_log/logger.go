package xor_log

import (
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"xor-go/pkg/xor_app"
)

const (
	serviceTag     = "SYSTEM_SERVICE"
	environmentTag = "SYSTEM_ENV"
	dcTag          = "SYSTEM_DC"
)

type Config struct {
	Level            string   `yaml:"level" env:"LEVEL"`
	Encoding         string   `yaml:"encoding" env:"ENCODING"`
	OutputPaths      []string `yaml:"output_paths" env:"OUTPUT_PATHS"`
	ErrorOutputPaths []string `yaml:"error_output_paths" env:"ERROR_OUTPUT_PATHS"`
}

// TODO support using sentry

func Init(cfg *Config, appCfg *xor_app.Config) (*zap.Logger, error) {
	cfgZap, err := configByEnv(appCfg.Environment)
	if err != nil {
		return nil, err
	}
	var level zapcore.Level
	err = level.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return nil, errors.WithMessage(err, "failed to unmarshall logger level")
	}

	// TODO support all options from original zap.Config
	cfgZap.Level = zap.NewAtomicLevelAt(level)
	cfgZap.OutputPaths = cfg.OutputPaths
	cfgZap.ErrorOutputPaths = cfg.ErrorOutputPaths
	cfgZap.Encoding = cfg.Encoding
	cfgZap.InitialFields = map[string]interface{}{
		serviceTag:     appCfg.Service,
		environmentTag: string(appCfg.Environment),
		dcTag:          appCfg.Dc,
	}

	logger, err := cfgZap.Build()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to build logger")
	}

	return logger, nil
}

func configByEnv(env xor_app.Environment) (*zap.Config, error) {
	var config zap.Config
	var err error = nil
	switch env {
	case xor_app.DevEnvironment:
		config = zap.NewDevelopmentConfig()
	case xor_app.ProdEnvironment:
		config = zap.NewProductionConfig()
	default:
		err = fmt.Errorf("failed to get logger config for env: %s", env)
	}
	if err != nil {
		return nil, err
	}
	return &config, nil
}
