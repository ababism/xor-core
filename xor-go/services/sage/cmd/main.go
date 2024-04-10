package main

import (
	"log"
	"xor-go/services/sage/internal/app"
	"xor-go/services/sage/internal/config"
)

const (
	configPath          = "./services/sage/configs/config.dev.yml"
	resourcesConfigPath = "./services/sage/configs/resources-config.yml"
)

func main() {
	var cfg config.Config
	err := config.ParseConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf(err.Error())
	}

	var servicesCfg config.ResourcesConfig
	err = config.ParseConfig(resourcesConfigPath, &servicesCfg)
	if err != nil {
		log.Fatalf(err.Error())
	}

	application, err := app.NewApp(&cfg, &servicesCfg)
	if err != nil {
		log.Fatalf("Failed to create app: %v", err)
	}
	application.Start()
}
