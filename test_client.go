package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// OrderRequest representa uma solicitação de pedido
type OrderRequest struct {
	UserID      uint32  `json:"user_id"`
	ProductName string  `json:"product_name"`
	Quantity    uint32  `json:"quantity"`
	Price       float64 `json:"price"`
}

// OrderResponse representa uma resposta de pedido
type OrderResponse struct {
	ID          uint32  `json:"id"`
	UserID      uint32  `json:"user_id"`
	ProductName string  `json:"product_name"`
	Quantity    uint32  `json:"quantity"`
	Price       float64 `json:"price"`
	Status      string  `json:"status"`
}

// OrderServiceClient é o cliente do serviço de pedidos
type OrderServiceClient interface {
	CreateOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

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

func main() {
	// Conectar ao order service
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Falha ao conectar: %v", err)
	}
	defer conn.Close()

	// Criar cliente
	client := NewOrderServiceClient(conn)

	// Criar pedido
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := &OrderRequest{
		UserID:      1,
		ProductName: "Produto Teste",
		Quantity:    2,
		Price:       99.90,
	}

	response, err := client.CreateOrder(ctx, request)
	if err != nil {
		log.Fatalf("Falha ao criar pedido: %v", err)
	}

	fmt.Printf("✅ Pedido criado com sucesso: %+v\n", response)
}
