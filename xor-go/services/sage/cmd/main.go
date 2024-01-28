package main

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"xor-go/services/sage/internal/app"
	"xor-go/services/sage/internal/config"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config-path", "", "path to .yaml config")
	flag.Parse()

	if configPath == "" {
		path := os.Getenv("CONFIG_PATH")
		if path == "" {
			log.Fatal("no config file provided")
		}
		configPath = path
	}
	var appConfig config.Config
	if err := cleanenv.ReadConfig(configPath, &appConfig); err != nil {
		log.Fatalf("failed reading config from path: {%s}, with error: {%v}", configPath, err)
	}

	application, err := app.NewApp(&appConfig)
	if err != nil {
		log.Fatalf("failed creating app with error: {%v}", err)
	}
	application.Start()
}
