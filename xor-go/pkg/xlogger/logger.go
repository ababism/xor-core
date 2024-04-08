package xlogger

import (
	"fmt"
	"github.com/TheZeroSlave/zapsentry"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"xor-go/pkg/xapp"
)

const (
	serviceTag     = "SYSTEM_SERVICE"
	environmentTag = "SYSTEM_ENV"
	dcTag          = "SYSTEM_DC"
)

type Config struct {
	Level            string   `mapstructure:"level" env:"LEVEL"`
	Env              string   `mapstructure:"env" env:"ENV"`
	Encoding         string   `mapstructure:"encoding" env:"ENCODING"`
	OutputPaths      []string `mapstructure:"output_paths" env:"OUTPUT_PATHS"`
	ErrorOutputPaths []string `mapstructure:"error_output_paths" env:"ERROR_OUTPUT_PATHS"`
	SentryLevel      string   `mapstructure:"sentry_level"`
	SentryDSN        string   `mapstructure:"sentry_dsn"`
}

// TODO support using sentry

func Init(cfg *Config, appCfg *xapp.Config) (*zap.Logger, error) {
	cfgZap, err := configByEnv(appCfg.Environment)
	if err != nil {
		return nil, err
	}
	var levelSentry, level zapcore.Level
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

	// TODO add new field to config
	err = levelSentry.UnmarshalText([]byte(cfg.SentryLevel))
	if err != nil {
		log.Printf("Sentry logs level with value=%s not initialized", levelSentry)
		return nil, err
	}
	cfgSentry := zapsentry.Configuration{
		Level: levelSentry,
		Tags: map[string]string{
			"environment": cfg.Env,
			"app":         appCfg.Name,
		},
	}
	core, err := zapsentry.NewCore(
		cfgSentry,
		zapsentry.NewSentryClientFromDSN(cfg.SentryDSN),
	)
	if err != nil {
		log.Println("Zapsentry NewCore not initialized")
		return nil, err
	}

	logger = zapsentry.AttachCoreToLogger(core, logger)

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
