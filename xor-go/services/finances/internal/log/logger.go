package log

import (
	"go.uber.org/zap"
	"xor-go/pkg/xapp"
	"xor-go/pkg/xlogger"
)

// TODO убрать *

var Logger *zap.Logger

func Init(cfgLogger *xlogger.Config, cfgApp *xapp.Config) error {
	var err error
	Logger, err = xlogger.Init(cfgLogger, cfgApp)
	if err != nil {
		return err
	}

	return err
}
