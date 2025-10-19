package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/seu-usuario/go-microservices-architecture/services/order/internal/config"
	"github.com/seu-usuario/go-microservices-architecture/services/order/internal/entity"
)

func main() {
	// 🧩 Conexão com DB
	db := config.ConnectDatabase()

	// ⚙️ Migração automática
	err := db.AutoMigrate(&entity.Order{})
	if err != nil {
		log.Fatalf("Erro ao migrar banco: %v", err)
	}

	fmt.Println("📦 Migração concluída para Order")

	// 🚀 Inicializa servidor gRPC (placeholder)
	port := ":50053"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}

	fmt.Printf("🚀 OrderService rodando em gRPC %s\n", port)
	s := grpc.NewServer()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Erro ao servir: %v", err)
	}
}
