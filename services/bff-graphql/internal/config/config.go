package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config estrutura para configurações do BFF GraphQL
type Config struct {
	// Servidor GraphQL
	GraphQLPort string
	Port        string // Alias para GraphQLPort

	// Endereços dos serviços gRPC
	OrderGRPCAddr        string
	UserGRPCAddr         string
	PaymentGRPCAddr      string
	NotificationGRPCAddr string

	// Aliases para compatibilidade com main.go
	OrderServiceURL        string
	UserServiceURL         string
	PaymentServiceURL      string
	NotificationServiceURL string

	// Configurações gerais
	Environment string
	GRPCTimeout time.Duration
	LogLevel    string
}

// Load carrega as configurações das variáveis de ambiente (alias para LoadConfig)
func Load() (*Config, error) {
	return LoadConfig(), nil
}

// LoadConfig carrega as configurações das variáveis de ambiente
func LoadConfig() *Config {
	// Tenta carregar .env file (opcional)
	_ = godotenv.Load()

	cfg := &Config{
		GraphQLPort:          getEnv("PORT", "8080"),
		OrderGRPCAddr:        getEnv("ORDER_SERVICE_URL", "localhost:50052"),
		UserGRPCAddr:         getEnv("USER_SERVICE_URL", "localhost:50051"),
		PaymentGRPCAddr:      getEnv("PAYMENT_SERVICE_URL", "localhost:50053"),
		NotificationGRPCAddr: getEnv("NOTIFICATION_SERVICE_URL", "localhost:50055"),
		Environment:          getEnv("GO_ENV", "development"),
		GRPCTimeout:          getEnvAsDuration("GRPC_TIMEOUT", 10*time.Second),
		LogLevel:             getEnv("LOG_LEVEL", "info"),
	}

	// Adicionar campos para compatibilidade
	cfg.Port = cfg.GraphQLPort
	cfg.OrderServiceURL = cfg.OrderGRPCAddr
	cfg.UserServiceURL = cfg.UserGRPCAddr
	cfg.PaymentServiceURL = cfg.PaymentGRPCAddr
	cfg.NotificationServiceURL = cfg.NotificationGRPCAddr

	return cfg
}

// getEnv obtém variável de ambiente com valor padrão
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt obtém variável de ambiente como inteiro
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// getEnvAsDuration obtém variável de ambiente como duration
func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if seconds, err := strconv.Atoi(value); err == nil {
			return time.Duration(seconds) * time.Second
		}
	}
	return defaultValue
}
