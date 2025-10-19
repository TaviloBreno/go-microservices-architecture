package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"github.com/seu-usuario/go-microservices-architecture/services/order/internal/config"
	"github.com/seu-usuario/go-microservices-architecture/services/order/internal/messaging"
	"github.com/seu-usuario/go-microservices-architecture/services/order/internal/metrics"
	"github.com/seu-usuario/go-microservices-architecture/services/order/internal/repository"
	"github.com/seu-usuario/go-microservices-architecture/services/order/internal/service"
	"github.com/seu-usuario/go-microservices-architecture/services/order/internal/telemetry"
	grpcServer "github.com/seu-usuario/go-microservices-architecture/services/order/internal/transport/grpc"
	"github.com/seu-usuario/go-microservices-architecture/services/order/proto"
)

func main() {
	// 🔧 Carregar variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}

	// 📊 Inicializar métricas Prometheus
	log.Println("📊 Inicializando métricas Prometheus...")
	metrics.Init()

	// 🔍 Inicializar OpenTelemetry Tracing
	log.Println("🔍 Inicializando OpenTelemetry Tracing...")
	ctx := context.Background()
	shutdown := telemetry.InitTracer("order-service")
	defer shutdown(ctx)

	// 🧩 Conectar ao banco de dados
	log.Println("🔗 Conectando ao banco de dados...")
	db := config.ConnectDatabase()

	// ⚙️ Executar migração automática
	if err := config.AutoMigrate(db); err != nil {
		log.Fatalf("❌ Erro ao executar migração: %v", err)
	}

	// 🐰 Conectar ao RabbitMQ
	log.Println("🐰 Conectando ao RabbitMQ...")
	rabbitmqConn, err := messaging.NewRabbitMQConnection()
	if err != nil {
		log.Printf("⚠️ Erro ao conectar ao RabbitMQ: %v", err)
		log.Println("⚠️ Continuando sem RabbitMQ - eventos não serão publicados")
		rabbitmqConn = nil
	}

	// 🏗️ Inicializar dependências
	log.Println("🏗️  Inicializando dependências...")
	orderRepo := repository.NewOrderRepository(db)

	var orderPublisher *messaging.OrderPublisher
	if rabbitmqConn != nil {
		orderPublisher = messaging.NewOrderPublisher(rabbitmqConn)
	}

	orderService := service.NewOrderService(orderRepo, orderPublisher)
	orderGRPCServer := grpcServer.NewOrderGRPCServer(orderService)

	// 🚀 Configurar servidor gRPC
	port := config.GetEnv("SERVICE_PORT", "50053")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("❌ Erro ao iniciar listener: %v", err)
	}

	// 📡 Criar e configurar servidor gRPC
	s := grpc.NewServer()
	proto.RegisterOrderServiceServer(s, orderGRPCServer)

	// 🎯 Iniciar servidor em goroutine
	go func() {
		log.Printf("🚀 OrderService rodando em gRPC :%s", port)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("❌ Erro ao servir gRPC: %v", err)
		}
	}()

	// 🛡️ Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Desligando OrderService...")
	s.GracefulStop()

	// Fechar conexão RabbitMQ se existir
	if rabbitmqConn != nil {
		if err := rabbitmqConn.Close(); err != nil {
			log.Printf("⚠️ Erro ao fechar RabbitMQ: %v", err)
		}
	}

	log.Println("✅ OrderService desligado com sucesso")
}
