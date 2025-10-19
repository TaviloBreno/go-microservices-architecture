package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

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

	// Configuração gRPC será feita sob demanda para evitar falhas no startup
	var notificationClient NotificationServiceClient

	// Endpoint de teste
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok", "message": "BFF funcionando!"})
	})

	// Endpoint para listar notificações
	http.HandleFunc("/notifications", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Conectar ao notification service sob demanda
		notificationConn, err := grpc.Dial("notification-service:50055", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("Erro ao conectar com notification service: %v", err)
			http.Error(w, "Erro ao conectar com notification service", http.StatusInternalServerError)
			return
		}
		defer notificationConn.Close()

		notificationClient = NewNotificationServiceClient(notificationConn)

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

	// Root endpoint (deve ser registrado por último)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Write([]byte("BFF Online 🚀"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
