package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"payment-service/internal/database"
	"payment-service/internal/messaging"
	"payment-service/internal/repository"
	"payment-service/internal/service"
	"payment-service/internal/transport"
	pb "payment-service/proto"

	"google.golang.org/grpc"
)

func main() {
	// Inicializar conexão com o banco de dados
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Falha ao conectar com o banco de dados: %v", err)
	}

	// Executar migrações
	err = database.Migrate(db)
	if err != nil {
		log.Fatalf("Falha ao executar migrações: %v", err)
	}

	// Inicializar dependências
	paymentRepo := repository.NewPaymentRepository(db)
	paymentService := service.NewPaymentService(paymentRepo)

	// Inicializar servidor gRPC
	grpcServer := grpc.NewServer()
	paymentGRPCServer := transport.NewPaymentGRPCServer(paymentService)
	pb.RegisterPaymentServiceServer(grpcServer, paymentGRPCServer)

	// Inicializar consumer RabbitMQ
	consumer, err := messaging.NewOrderConsumer(paymentService)
	if err != nil {
		log.Fatalf("Falha ao inicializar consumer: %v", err)
	}
	defer consumer.Close()

	// Configurar consumer para processar mensagens
	go func() {
		err := consumer.ConsumeOrders(context.Background())
		if err != nil {
			log.Printf("Erro no consumer: %v", err)
		}
	}()

	// Iniciar servidor gRPC
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Falha ao criar listener: %v", err)
	}

	// Canal para capturar sinais do sistema
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Iniciar servidor gRPC em goroutine
	go func() {
		fmt.Println("Payment service rodando na porta :50053")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Falha ao iniciar servidor gRPC: %v", err)
		}
	}()

	// Aguardar sinal de parada
	<-sigChan
	fmt.Println("\nRecebido sinal de parada. Desligando payment service...")

	// Graceful shutdown
	grpcServer.GracefulStop()
	fmt.Println("Payment service desligado com sucesso.")
}
