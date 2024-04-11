package main

import (
	"log"
	"xor-go/services/sage/internal/app"
	"xor-go/services/sage/internal/config"
)

const (
	configPath          = "./services/sage/config/config.local.yml"
	resourcesConfigPath = "./services/sage/config/resources-config.yml"
)

func main() {
	var cfg config.Config
	err := config.ParseConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf(err.Error())
	}

	var resourcesCfg config.ResourcesConfig
	err = config.ParseConfig(resourcesConfigPath, &resourcesCfg)
	if err != nil {
		log.Fatalf(err.Error())
	}

	application, err := app.NewApp(&cfg, &resourcesCfg)
	if err != nil {
		log.Fatalf("Failed to create app: %v", err)
	}
	application.Start()
}
