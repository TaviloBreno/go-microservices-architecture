package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Estruturas para as requisições
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateOrderRequest struct {
	UserID      uint32  `json:"user_id"`
	ProductName string  `json:"product_name"`
	Quantity    int32   `json:"quantity"`
	Price       float64 `json:"price"`
}

// Estruturas para gRPC
type ListNotificationsRequest struct{}

type NotificationResponse struct {
	ID        uint32 `json:"id"`
	PaymentID uint32 `json:"payment_id"`
	OrderID   uint32 `json:"order_id"`
	Message   string `json:"message"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

type ListNotificationsResponse struct {
	Notifications []*NotificationResponse `json:"notifications"`
}

// NotificationServiceClient interface para gRPC
type NotificationServiceClient interface {
	ListNotifications(ctx context.Context, in *ListNotificationsRequest, opts ...grpc.CallOption) (*ListNotificationsResponse, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) ListNotifications(ctx context.Context, in *ListNotificationsRequest, opts ...grpc.CallOption) (*ListNotificationsResponse, error) {
	out := new(ListNotificationsResponse)
	err := c.cc.Invoke(ctx, "/notification.NotificationService/ListNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func main() {
	fmt.Println("🌐 BFF rodando na porta 8080...")

	// Configurar conexões gRPC
	userConn, err := grpc.Dial("user-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Erro ao conectar com user service:", err)
	}
	defer userConn.Close()

	orderConn, err := grpc.Dial("order-service:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Erro ao conectar com order service:", err)
	}
	defer orderConn.Close()

	paymentConn, err := grpc.Dial("payment-service:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Erro ao conectar com payment service:", err)
	}
	defer paymentConn.Close()

	notificationConn, err := grpc.Dial("notification-service:50055", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Erro ao conectar com notification service:", err)
	}
	defer notificationConn.Close()

	notificationClient := NewNotificationServiceClient(notificationConn)

	// Root endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("BFF Online 🚀"))
	})

	// Endpoint para criar usuários
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req CreateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Chamar user service via HTTP (simplificado)
		userServiceURL := "http://user-service:8081/users"

		jsonData, _ := json.Marshal(req)
		resp, err := http.Post(userServiceURL, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			http.Error(w, "Erro ao chamar user service", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	})

	// Endpoint para criar pedidos
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req CreateOrderRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Chamar order service via HTTP (simplificado)
		orderServiceURL := "http://order-service:8082/orders"

		jsonData, _ := json.Marshal(req)
		resp, err := http.Post(orderServiceURL, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			http.Error(w, "Erro ao chamar order service", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	})

	// Endpoint para listar pagamentos
	http.HandleFunc("/payments", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Chamar payment service via HTTP (simplificado)
		paymentServiceURL := "http://payment-service:8083/payments"

		resp, err := http.Get(paymentServiceURL)
		if err != nil {
			http.Error(w, "Erro ao chamar payment service", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	})

	// Endpoint para listar notificações
	http.HandleFunc("/notifications", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		response, err := notificationClient.ListNotifications(ctx, &ListNotificationsRequest{})
		if err != nil {
			log.Printf("Erro ao chamar notification service: %v", err)
			http.Error(w, "Erro ao chamar notification service", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response.Notifications)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
