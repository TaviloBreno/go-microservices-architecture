package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServicePort string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPass      string
	DBName      string
	RabbitMQURL string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		ServicePort: getEnv("SERVICE_PORT", "50053"),
		DBHost:      getEnv("DB_HOST", "mysql"),
		DBPort:      getEnv("DB_PORT", "3306"),
		DBUser:      getEnv("DB_USER", "root"),
		DBPass:      getEnv("DB_PASSWORD", "secret"),
		DBName:      getEnv("DB_NAME", "order_service"),
		RabbitMQURL: getEnv("RABBITMQ_URL", "amqp://guest:guest@rabbitmq:5672/"),
	}

	log.Printf("Config carregada: %+v", cfg)
	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// GetEnv retorna uma variável de ambiente ou um valor padrão (exportada)
func GetEnv(key, fallback string) string {
	return getEnv(key, fallback)
}
