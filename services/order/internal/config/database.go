package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	user := getEnv("DB_USER", "root")
	pass := getEnv("DB_PASSWORD", "secret")
	host := getEnv("DB_HOST", "mysql")
	port := getEnv("DB_PORT", "3306")
	name := getEnv("DB_NAME", "order_service")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name)

	log.Printf("🔗 Tentando conectar ao banco de dados: %s@%s:%s/%s", user, host, port, name)

	var db *gorm.DB
	var err error
	maxRetries := 10

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("✅ Conectado ao banco de dados MySQL com sucesso")
			return db
		}

		log.Printf("⏳ Tentativa %d/%d falhou: %v. Aguardando 3 segundos...", i+1, maxRetries, err)
		time.Sleep(3 * time.Second)
	}

	log.Fatalf("❌ Erro ao conectar ao banco após %d tentativas: %v", maxRetries, err)
	return nil
}
