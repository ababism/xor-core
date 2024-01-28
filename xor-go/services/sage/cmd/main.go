package main

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	XorLogger "xor-go/libs/xor_logger"
	"xor-go/services/sage/internal/config"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config-path", "", "path to .yaml config")
	flag.Parse()

	if configPath == "" {
		path := os.Getenv("CONFIG_PATH")
		if path == "" {
			log.Panic("no config file provided")
		}
		configPath = path
	}
	var appConfig config.Config
	if err := cleanenv.ReadConfig(configPath, &appConfig); err != nil {
		log.Panicf("failed reading config from path: %s, with err: %v", configPath, err)
	}

	logger, err := XorLogger.InitLogger(appConfig.LoggerConfig.ToXorLoggerConfig())
	if err != nil {
		log.Panicf("failed initializing logger with err: %v", err)
	}
	logger.Info("OKx")
	logger.Error("NOT OK")
}
