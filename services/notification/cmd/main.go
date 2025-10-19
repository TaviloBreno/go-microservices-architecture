package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"notification-service/internal/database"
	"notification-service/internal/messaging"
	"notification-service/internal/repository"
	"notification-service/internal/service"
	"notification-service/internal/transport"
	pb "notification-service/proto"

	"google.golang.org/grpc"
)

func main() {
	log.Println("🚀 Iniciando Notification Service...")

	// Inicializar conexão com o banco de dados
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("❌ Falha ao conectar com o banco de dados: %v", err)
	}

	// Executar migrações
	err = database.Migrate(db)
	if err != nil {
		log.Fatalf("❌ Falha ao executar migrações: %v", err)
	}

	// Inicializar dependências
	notificationRepo := repository.NewNotificationRepository(db)
	notificationService := service.NewNotificationService(notificationRepo)

	// Inicializar servidor gRPC
	grpcServer := grpc.NewServer()
	notificationGRPCServer := transport.NewNotificationGRPCServer(notificationService)
	pb.RegisterNotificationServiceServer(grpcServer, notificationGRPCServer)

	// Inicializar consumer RabbitMQ
	consumer, err := messaging.NewPaymentConsumer(notificationService)
	if err != nil {
		log.Fatalf("❌ Falha ao inicializar consumer: %v", err)
	}
	defer consumer.Close()

	// Criar contexto para graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Configurar consumer para processar mensagens
	go func() {
		log.Println("🎧 Iniciando consumer RabbitMQ...")
		err := consumer.ConsumePayments(ctx)
		if err != nil {
			log.Printf("❌ Erro no consumer: %v", err)
		}
	}()

	// Iniciar servidor gRPC
	lis, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatalf("❌ Falha ao criar listener: %v", err)
	}

	// Canal para capturar sinais do sistema
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Iniciar servidor gRPC em goroutine
	go func() {
		fmt.Println("📡 Notification service rodando na porta :50055")
		log.Println("🎯 gRPC server listening on :50055")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("❌ Falha ao iniciar servidor gRPC: %v", err)
		}
	}()

	log.Println("✅ Notification Service iniciado com sucesso!")
	log.Println("📋 Serviços disponíveis:")
	log.Println("   📡 gRPC Server: :50055")
	log.Println("   🐰 RabbitMQ Consumer: payments queue")
	log.Println("   🗄️  MySQL Database: notification_service")

	// Aguardar sinal de parada
	<-sigChan
	fmt.Println("\n🛑 Recebido sinal de parada. Desligando notification service...")

	// Graceful shutdown
	cancel() // Cancelar contexto do consumer
	grpcServer.GracefulStop()

	fmt.Println("✅ Notification service desligado com sucesso.")
}
