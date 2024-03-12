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
const ServiceName = "Finances Service"
const ServiceCapsName = "FINANCES SERVICE"

func init() {
	if err := godotenv.Load(MainEnvName); err != nil {
		log.Print(fmt.Sprintf("No '%s' file found", MainEnvName))
	}
}

func main() {
	ctx := context.Background()

	configPath := os.Getenv("CONFIG_" + ServiceCapsName)
	log.Println(ServiceName+" config path: ", configPath)

	// Собираем конфиг приложения
	cfg, err := config.NewConfig(configPath, ServiceCapsName)
	if err != nil {
		log.Fatalf("Fail to parse %s config: %v", ServiceName, err)
	}

	// Создаем наше приложение
	application, err := app.NewApp(cfg)
	if err != nil {
		log.Fatal(fmt.Sprintf("Fail to create %s app: %s", cfg.App, err))
	}

	// Запускаем приложение
	application.Start(ctx)
}
