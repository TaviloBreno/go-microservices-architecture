package clients

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// OrderRequest representa uma requisição de criação de pedido
type OrderRequest struct {
	Customer  string  `json:"customer"`
	ProductID uint32  `json:"product_id"`
	Quantity  int32   `json:"quantity"`
	Price     float64 `json:"price"`
}

// OrderResponse representa a resposta de um pedido
type OrderResponse struct {
	ID        uint32  `json:"id"`
	Customer  string  `json:"customer"`
	ProductID uint32  `json:"product_id"`
	Quantity  int32   `json:"quantity"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"created_at"`
}

// ListOrdersRequest representa uma requisição para listar pedidos
type ListOrdersRequest struct{}

// ListOrdersResponse representa a resposta com lista de pedidos
type ListOrdersResponse struct {
	Orders []*OrderResponse `json:"orders"`
}

// OrderClient interface para comunicação com order-service
type OrderClient interface {
	CreateOrder(ctx context.Context, req *OrderRequest) (*OrderResponse, error)
	ListOrders(ctx context.Context) ([]*OrderResponse, error)
	Close() error
}

// grpcOrderClient implementa OrderClient usando gRPC
type grpcOrderClient struct {
	conn   *grpc.ClientConn
	client OrderServiceClient
}

// OrderServiceClient interface gRPC (simplificada para evitar dependências proto)
type OrderServiceClient interface {
	CreateOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error)
	ListOrders(ctx context.Context, in *ListOrdersRequest, opts ...grpc.CallOption) (*ListOrdersResponse, error)
}

// orderServiceClient implementa OrderServiceClient
type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

// NewOrderServiceClient cria um novo cliente gRPC para OrderService
func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, "/order.OrderService/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) ListOrders(ctx context.Context, in *ListOrdersRequest, opts ...grpc.CallOption) (*ListOrdersResponse, error) {
	out := new(ListOrdersResponse)
	err := c.cc.Invoke(ctx, "/order.OrderService/ListOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewOrderClient cria um novo cliente para order-service
func NewOrderClient(grpcAddr string, timeout time.Duration) (OrderClient, error) {
	log.Printf("🔌 Conectando ao Order Service em %s...", grpcAddr)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, grpcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		log.Printf("❌ Erro ao conectar com Order Service: %v", err)
		return nil, err
	}

	client := NewOrderServiceClient(conn)
	log.Printf("✅ Conectado ao Order Service com sucesso!")

	return &grpcOrderClient{
		conn:   conn,
		client: client,
	}, nil
}

// CreateOrder cria um novo pedido
func (c *grpcOrderClient) CreateOrder(ctx context.Context, req *OrderRequest) (*OrderResponse, error) {
	log.Printf("📦 Criando pedido: %+v", req)

	response, err := c.client.CreateOrder(ctx, req)
	if err != nil {
		log.Printf("❌ Erro ao criar pedido: %v", err)
		return nil, err
	}

	log.Printf("✅ Pedido criado com sucesso: ID %d", response.ID)
	return response, nil
}

// ListOrders lista todos os pedidos
func (c *grpcOrderClient) ListOrders(ctx context.Context) ([]*OrderResponse, error) {
	log.Printf("📋 Listando pedidos...")

	response, err := c.client.ListOrders(ctx, &ListOrdersRequest{})
	if err != nil {
		log.Printf("❌ Erro ao listar pedidos: %v", err)
		return nil, err
	}

	log.Printf("✅ Encontrados %d pedidos", len(response.Orders))
	return response.Orders, nil
}

// Close fecha a conexão
func (c *grpcOrderClient) Close() error {
	log.Printf("🔌 Fechando conexão com Order Service...")
	return c.conn.Close()
}
