package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"xor-go/services/courses/internal/app"
	"xor-go/services/courses/internal/config"
)

const MainEnvName = ".env"
const AppCapsName = "COURSES"

func init() {
	if err := godotenv.Load(MainEnvName); err != nil {
		log.Print(fmt.Sprintf("No '%s' file found", MainEnvName))
	}
}

func main() {
	ctx := context.Background()

	configPath := os.Getenv("CONFIG_" + AppCapsName)
	log.Println("Courses config path: ", configPath)

	// Собираем конфиг приложения
	cfg, err := config.NewConfig(configPath, AppCapsName)
	if err != nil {
		log.Fatal("Fail to parse courses config: ", err)
	}

	// Создаем наше приложение
	application, err := app.NewApp(cfg)
	if err != nil {
		log.Fatal(fmt.Sprintf("Fail to create %s app: %s", cfg.App.Name, err))
	}

	// Запускаем приложение
	application.Start(ctx)
}
