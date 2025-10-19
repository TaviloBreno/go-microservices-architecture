package grpc

import (
	"context"
	"log"
	"time"

	"github.com/seu-usuario/go-microservices-architecture/services/order/internal/service"
	"github.com/seu-usuario/go-microservices-architecture/services/order/proto"
)

// OrderGRPCServer implementa o servidor gRPC para pedidos
type OrderGRPCServer struct {
	proto.UnimplementedOrderServiceServer
	orderService service.OrderService
}

// NewOrderGRPCServer cria uma nova instância do servidor gRPC
func NewOrderGRPCServer(orderService service.OrderService) *OrderGRPCServer {
	return &OrderGRPCServer{
		orderService: orderService,
	}
}

// CreateOrder cria um novo pedido via gRPC
func (s *OrderGRPCServer) CreateOrder(ctx context.Context, req *proto.OrderRequest) (*proto.OrderResponse, error) {
	log.Printf("📝 Recebida requisição CreateOrder: Customer=%s, ProductID=%d, Quantity=%d, Price=%.2f",
		req.GetCustomer(), req.GetProductId(), req.GetQuantity(), req.GetPrice())

	// Chamar o serviço de negócio
	order, err := s.orderService.CreateOrder(
		req.GetCustomer(),
		uint(req.GetProductId()),
		int(req.GetQuantity()),
		req.GetPrice(),
	)
	if err != nil {
		log.Printf("❌ Erro ao criar pedido: %v", err)
		return nil, err
	}

	// Converter para resposta gRPC
	response := &proto.OrderResponse{
		Id:        uint32(order.ID),
		Customer:  order.Customer,
		ProductId: uint32(order.ProductID),
		Quantity:  int32(order.Quantity),
		Price:     order.Price,
		CreatedAt: order.CreatedAt.Format(time.RFC3339),
	}

	log.Printf("✅ Pedido criado com sucesso: ID=%d", order.ID)
	return response, nil
}

// ListOrders lista todos os pedidos via gRPC
func (s *OrderGRPCServer) ListOrders(ctx context.Context, req *proto.ListOrdersRequest) (*proto.ListOrdersResponse, error) {
	log.Println("📋 Recebida requisição ListOrders")

	// Chamar o serviço de negócio
	orders, err := s.orderService.ListOrders()
	if err != nil {
		log.Printf("❌ Erro ao listar pedidos: %v", err)
		return nil, err
	}

	// Converter para resposta gRPC
	var orderResponses []*proto.OrderResponse
	for _, order := range orders {
		orderResponse := &proto.OrderResponse{
			Id:        uint32(order.ID),
			Customer:  order.Customer,
			ProductId: uint32(order.ProductID),
			Quantity:  int32(order.Quantity),
			Price:     order.Price,
			CreatedAt: order.CreatedAt.Format(time.RFC3339),
		}
		orderResponses = append(orderResponses, orderResponse)
	}

	response := &proto.ListOrdersResponse{
		Orders: orderResponses,
	}

	log.Printf("✅ Listados %d pedidos", len(orders))
	return response, nil
}
