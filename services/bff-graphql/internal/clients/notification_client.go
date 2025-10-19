package clients

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NotificationResponse representa uma notificação
type NotificationResponse struct {
	ID        uint32 `json:"id"`
	PaymentID uint32 `json:"payment_id"`
	OrderID   uint32 `json:"order_id"`
	Message   string `json:"message"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

// ListNotificationsRequest representa uma requisição para listar notificações
type ListNotificationsRequest struct{}

// ListNotificationsResponse representa a resposta com lista de notificações
type ListNotificationsResponse struct {
	Notifications []*NotificationResponse `json:"notifications"`
}

// NotificationClient interface para comunicação com notification-service
type NotificationClient interface {
	ListNotifications(ctx context.Context) ([]*NotificationResponse, error)
	Close() error
}

// NotificationServiceClient interface gRPC (simplificada)
type NotificationServiceClient interface {
	ListNotifications(ctx context.Context, in *ListNotificationsRequest, opts ...grpc.CallOption) (*ListNotificationsResponse, error)
}

// notificationServiceClient implementa NotificationServiceClient
type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

// NewNotificationServiceClient cria um novo cliente gRPC para NotificationService
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

// grpcNotificationClient implementa NotificationClient usando gRPC
type grpcNotificationClient struct {
	conn   *grpc.ClientConn
	client NotificationServiceClient
}

// NewNotificationClient cria um novo cliente para notification-service
func NewNotificationClient(grpcAddr string, timeout time.Duration) (NotificationClient, error) {
	log.Printf("🔌 Conectando ao Notification Service em %s...", grpcAddr)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, grpcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		log.Printf("❌ Erro ao conectar com Notification Service: %v", err)
		return nil, err
	}

	client := NewNotificationServiceClient(conn)
	log.Printf("✅ Conectado ao Notification Service com sucesso!")

	return &grpcNotificationClient{
		conn:   conn,
		client: client,
	}, nil
}

// ListNotifications lista todas as notificações
func (c *grpcNotificationClient) ListNotifications(ctx context.Context) ([]*NotificationResponse, error) {
	log.Printf("📢 Listando notificações...")

	response, err := c.client.ListNotifications(ctx, &ListNotificationsRequest{})
	if err != nil {
		log.Printf("❌ Erro ao listar notificações: %v", err)
		return nil, err
	}

	log.Printf("✅ Encontradas %d notificações", len(response.Notifications))
	return response.Notifications, nil
}

// Close fecha a conexão
func (c *grpcNotificationClient) Close() error {
	log.Printf("🔌 Fechando conexão com Notification Service...")
	return c.conn.Close()
}
