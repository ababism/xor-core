package xlogger

import (
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"xor-go/pkg/xapp"
)

const (
	serviceTag     = "SYSTEM_SERVICE"
	environmentTag = "SYSTEM_ENV"
	dcTag          = "SYSTEM_DC"
)

type Config struct {
	Level            string   `yaml:"level" env:"LEVEL"`
	Env              string   `yaml:"env" env:"ENV"`
	Encoding         string   `yaml:"encoding" env:"ENCODING"`
	OutputPaths      []string `yaml:"output_paths" env:"OUTPUT_PATHS"`
	ErrorOutputPaths []string `yaml:"error_output_paths" env:"ERROR_OUTPUT_PATHS"`
}

// TODO support using sentry

func Init(cfg *Config, appCfg *xapp.Config) (*zap.Logger, error) {
	cfgZap, err := configByEnv(appCfg.Environment)
	if err != nil {
		return nil, err
	}
	var level zapcore.Level
	err = level.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return nil, errors.WithMessage(err, "failed to unmarshall logger level")
	}

	cfgZap.Level = zap.NewAtomicLevelAt(level)
	cfgZap.OutputPaths = cfg.OutputPaths
	cfgZap.ErrorOutputPaths = cfg.ErrorOutputPaths
	cfgZap.Encoding = cfg.Encoding
	cfgZap.InitialFields = map[string]interface{}{
		serviceTag:     appCfg.Name,
		environmentTag: string(appCfg.Environment),
		dcTag:          appCfg.Dc,
	}

	logger, err := cfgZap.Build()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to build logger")
	}

	return logger, nil
}

func configByEnv(env xapp.Environment) (*zap.Config, error) {
	var config zap.Config
	var err error = nil
	switch env {
	case xapp.DevelopmentEnv:
		config = zap.NewDevelopmentConfig()
	case xapp.ProductionEnv:
		config = zap.NewProductionConfig()
	default:
		err = fmt.Errorf("failed to get logger config for env: %s", env)
	}
	if err != nil {
		return nil, err
	}
	return &config, nil
}
