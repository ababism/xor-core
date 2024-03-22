package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/net/context"
	"log"
	"os"
	"xor-go/services/finances/internal/app"
	"xor-go/services/finances/internal/config"
)

const MainEnvName = ".env"
const ServiceName = "Finances"
const ServiceCapsName = "FINANCES"

func init() {
	envPath := fmt.Sprintf("services/finances/%s", MainEnvName)
	if err := godotenv.Load(envPath); err != nil {
		log.Print(fmt.Sprintf("No '%s' file found", MainEnvName))
	}
}

func main() {
	ctx := context.Background()

	cfgEnvName := "CONFIG_" + ServiceCapsName
	configPath := os.Getenv(cfgEnvName)
	log.Printf("%s config path (%s): %s", ServiceName, cfgEnvName, configPath)

	// Собираем конфиг приложения
	cfg, err := config.NewConfig(configPath, ServiceCapsName)
	if err != nil {
		log.Fatalf("Fail to parse %s config: %v", ServiceName, err)
	}

	// Создаем наше приложение
	application, err := app.NewApp(cfg)
	if err != nil {
		log.Fatal(fmt.Sprintf("Fail to create '%s' service: %s", cfg.App.Name, err))
	}

	// Запускаем приложение
	application.Start(ctx)
}
