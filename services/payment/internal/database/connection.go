package database

import (
	"fmt"
	"time"

	"payment-service/internal/config"
	"payment-service/internal/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect estabelece conexão com o banco de dados MySQL
func Connect() (*gorm.DB, error) {
	cfg := config.LoadConfig()

	// Configurações do banco de dados
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	// Conectar com retry
	var db *gorm.DB
	var err error

	for i := 0; i < 5; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}

		fmt.Printf("Tentativa %d de conexão com o banco falhou: %v\n", i+1, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("falha ao conectar com o banco após 5 tentativas: %w", err)
	}

	// Configurar pool de conexões
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("falha ao obter instância SQL: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("✅ Conexão com banco de dados estabelecida com sucesso!")
	return db, nil
}

// Migrate executa as migrações do banco de dados
func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&domain.Payment{},
	)

	if err != nil {
		return fmt.Errorf("falha ao executar migrações: %w", err)
	}

	fmt.Println("✅ Migrações executadas com sucesso!")
	return nil
}
