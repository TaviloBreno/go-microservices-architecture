package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/vektah/gqlparser/v2/ast"

	"github.com/seu-usuario/go-microservices-architecture/bff-graphql/graph"
	"github.com/seu-usuario/go-microservices-architecture/bff-graphql/internal/clients"
	"github.com/seu-usuario/go-microservices-architecture/bff-graphql/internal/config"
	"github.com/seu-usuario/go-microservices-architecture/bff-graphql/internal/metrics"
	"github.com/seu-usuario/go-microservices-architecture/bff-graphql/internal/telemetry"
)

func main() {
	log.Println("🚀 Iniciando BFF GraphQL Server")

	// 📊 Inicializar métricas Prometheus
	log.Println("📊 Inicializando métricas Prometheus...")
	metrics.Init()

	// 🔍 Inicializar OpenTelemetry Tracing
	log.Println("🔍 Inicializando OpenTelemetry Tracing...")
	ctx := context.Background()
	shutdown := telemetry.InitTracer("bff-service")
	defer shutdown(ctx)

	// Carregar configuração
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("❌ Erro ao carregar configuração:", err)
	}

	log.Printf("⚙️ Configuração carregada - Port: %s", cfg.Port)
	log.Printf("🔗 Conectando com microserviços:")
	log.Printf("   📦 Order Service: %s", cfg.OrderServiceURL)
	log.Printf("   👤 User Service: %s", cfg.UserServiceURL)
	log.Printf("   💳 Payment Service: %s", cfg.PaymentServiceURL)
	log.Printf("   📧 Notification Service: %s", cfg.NotificationServiceURL)

	// Inicializar clientes gRPC com timeout
	timeout := 10 * time.Second

	orderClient, err := clients.NewOrderClient(cfg.OrderServiceURL, timeout)
	if err != nil {
		log.Printf("⚠️ Aviso: Não foi possível conectar ao Order Service: %v", err)
	} else {
		log.Println("✅ Order Client conectado")
	}

	userClient, err := clients.NewUserClient(cfg.UserServiceURL, timeout)
	if err != nil {
		log.Printf("⚠️ Aviso: Não foi possível conectar ao User Service: %v", err)
	} else {
		log.Println("✅ User Client conectado")
	}

	paymentClient, err := clients.NewPaymentClient(cfg.PaymentServiceURL, timeout)
	if err != nil {
		log.Printf("⚠️ Aviso: Não foi possível conectar ao Payment Service: %v", err)
	} else {
		log.Println("✅ Payment Client conectado")
	}

	notificationClient, err := clients.NewNotificationClient(cfg.NotificationServiceURL, timeout)
	if err != nil {
		log.Printf("⚠️ Aviso: Não foi possível conectar ao Notification Service: %v", err)
	} else {
		log.Println("✅ Notification Client conectado")
	}

	// Criar resolver com todos os clientes
	resolver := &graph.Resolver{
		OrderClient:        orderClient,
		UserClient:         userClient,
		PaymentClient:      paymentClient,
		NotificationClient: notificationClient,
	}

	// Configurar servidor GraphQL
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	// Configurar transporte WebSocket para subscriptions
	srv.AddTransport(transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // Permitir qualquer origem para desenvolvimento
			},
		},
	})

	// Adicionar transporte HTTP
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	// Configurar cache e extensões
	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	// Configurar CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           86400, // 24 horas
	})

	// Configurar rotas
	mux := http.NewServeMux()

	// Endpoint GraphQL principal
	mux.Handle("/", corsHandler.Handler(srv))

	// Playground para desenvolvimento
	if os.Getenv("GO_ENV") != "production" {
		mux.Handle("/playground", playground.Handler("GraphQL Playground", "/"))
		log.Println("🎮 GraphQL Playground disponível em http://localhost:" + cfg.Port + "/playground")
	}

	// Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"healthy","service":"bff-graphql","timestamp":"` + time.Now().Format(time.RFC3339) + `"}`))
	})

	// Middleware de logging
	loggedMux := loggingMiddleware(mux)

	// Iniciar servidor
	log.Printf("🌍 Servidor GraphQL rodando na porta %s", cfg.Port)
	log.Printf("📡 GraphQL endpoint: http://localhost:%s/", cfg.Port)
	log.Printf("🏥 Health endpoint: http://localhost:%s/health", cfg.Port)

	if err := http.ListenAndServe(":"+cfg.Port, loggedMux); err != nil {
		log.Fatal("❌ Erro ao iniciar servidor:", err)
	}
}

// loggingMiddleware adiciona logs para todas as requisições HTTP
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Criar um ResponseWriter que captura o status code
		wrappedWriter := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		// Processar request
		next.ServeHTTP(wrappedWriter, r)

		// Log da requisição
		duration := time.Since(start)
		log.Printf("📊 %s %s %d %s %s",
			r.Method,
			r.URL.Path,
			wrappedWriter.statusCode,
			duration.String(),
			r.Header.Get("User-Agent"),
		)
	})
}

// responseWriter wrapper para capturar status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// Função auxiliar para converter string para int com valor padrão
func parseIntWithDefault(s string, defaultValue int) int {
	if val, err := strconv.Atoi(s); err == nil {
		return val
	}
	return defaultValue
}
