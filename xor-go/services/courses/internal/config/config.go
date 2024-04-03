package config

import (
	"github.com/spf13/viper"
	"log"
	"xor-go/pkg/app"
	configLib "xor-go/pkg/config"
	"xor-go/pkg/graceful_shutdown"
	"xor-go/pkg/http_server"
	"xor-go/pkg/metrics"
	"xor-go/pkg/mylogger"
	"xor-go/pkg/mytracer"
	kafkaConsumer "xor-go/services/courses/internal/daemons/kafkaConsumer"
	"xor-go/services/courses/internal/daemons/scraper"
	"xor-go/services/courses/internal/repository/financesclient"
	"xor-go/services/courses/internal/repository/kafkaproducer"
	"xor-go/services/courses/internal/repository/mongo"
)

type Config struct {
	App              *app.Config                  `mapstructure:"app"`
	Http             *http_server.Config          `mapstructure:"http"`
	FinancesClient   *financesclient.ClientConfig `mapstructure:"location_client"`
	Logger           *mylogger.Config             `mapstructure:"logger"`
	Mongo            *mongo.Config                `mapstructure:"mongo"`
	MigrationsMongo  *mongo.ConfigMigrations      `mapstructure:"migrations_mongo"`
	Metrics          *metrics.Config              `mapstructure:"metrics"`
	GracefulShutdown *graceful_shutdown.Config    `mapstructure:"graceful_shutdown"`
	KafkaReader      *kafkaConsumer.Config        `mapstructure:"kafka_reader"`
	KafkaWriter      *kafkaproducer.Config        `mapstructure:"kafka_writer"`
	Tracer           *mytracer.Config             `mapstructure:"tracer"`
	Scraper          *scraper.Config              `mapstructure:"scraper"`
}

func NewConfig(filePath string, appName string) (*Config, error) {
	viper.SetConfigFile(filePath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error while reading config file: %v", err)
	}

	// Загрузка конфигурации в структуру Config
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("error while unmarshalling config file: %v", err)
	}

	// Замена значений из переменных окружения, если они заданы
	configLib.ReplaceWithEnv(&config, appName)

	return &config, nil
}
