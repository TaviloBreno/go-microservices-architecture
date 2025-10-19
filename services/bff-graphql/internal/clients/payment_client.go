package clients

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// PaymentStatusRequest representa uma requisição de status de pagamento
type PaymentStatusRequest struct {
	OrderID uint32 `json:"order_id"`
}

// PaymentStatusResponse representa a resposta de status de pagamento
type PaymentStatusResponse struct {
	OrderID uint32  `json:"order_id"`
	Status  string  `json:"status"`
	Amount  float64 `json:"amount"`
}

// PaymentClient interface para comunicação com payment-service
type PaymentClient interface {
	GetPaymentStatus(ctx context.Context, orderID uint32) (*PaymentStatusResponse, error)
	Close() error
}

// PaymentServiceClient interface gRPC (simplificada)
type PaymentServiceClient interface {
	GetPaymentStatus(ctx context.Context, in *PaymentStatusRequest, opts ...grpc.CallOption) (*PaymentStatusResponse, error)
}

// paymentServiceClient implementa PaymentServiceClient
type paymentServiceClient struct {
	cc grpc.ClientConnInterface
}

// NewPaymentServiceClient cria um novo cliente gRPC para PaymentService
func NewPaymentServiceClient(cc grpc.ClientConnInterface) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) GetPaymentStatus(ctx context.Context, in *PaymentStatusRequest, opts ...grpc.CallOption) (*PaymentStatusResponse, error) {
	out := new(PaymentStatusResponse)
	err := c.cc.Invoke(ctx, "/payment.PaymentService/GetPaymentStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// grpcPaymentClient implementa PaymentClient usando gRPC
type grpcPaymentClient struct {
	conn   *grpc.ClientConn
	client PaymentServiceClient
}

// NewPaymentClient cria um novo cliente para payment-service
func NewPaymentClient(grpcAddr string, timeout time.Duration) (PaymentClient, error) {
	log.Printf("🔌 Conectando ao Payment Service em %s...", grpcAddr)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, grpcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		log.Printf("❌ Erro ao conectar com Payment Service: %v", err)
		return nil, err
	}

	client := NewPaymentServiceClient(conn)
	log.Printf("✅ Conectado ao Payment Service com sucesso!")

	return &grpcPaymentClient{
		conn:   conn,
		client: client,
	}, nil
}

// GetPaymentStatus obtém o status de um pagamento
func (c *grpcPaymentClient) GetPaymentStatus(ctx context.Context, orderID uint32) (*PaymentStatusResponse, error) {
	log.Printf("💳 Consultando status do pagamento para order ID: %d", orderID)

	req := &PaymentStatusRequest{OrderID: orderID}
	response, err := c.client.GetPaymentStatus(ctx, req)
	if err != nil {
		log.Printf("❌ Erro ao consultar status do pagamento: %v", err)
		return nil, err
	}

	log.Printf("✅ Status do pagamento obtido: %s", response.Status)
	return response, nil
}

// Close fecha a conexão
func (c *grpcPaymentClient) Close() error {
	log.Printf("🔌 Fechando conexão com Payment Service...")
	return c.conn.Close()
}
