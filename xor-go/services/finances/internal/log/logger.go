package log

import (
	"go.uber.org/zap"
	"xor-go/pkg/xlogger"
	"xor-go/services/finances/internal/config"
)

// TODO убрать *

var Logger *zap.Logger

func Init(cfg *config.Config) error {
	var err error
	Logger, err = xlogger.Init(cfg.Logger, cfg.App)
	if err != nil {
		return err
	}

	return err
}
