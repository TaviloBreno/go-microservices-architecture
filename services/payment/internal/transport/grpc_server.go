package transport

import (
	"context"

	"payment-service/internal/service"
	pb "payment-service/proto"
)

// PaymentGRPCServer implementa o servidor gRPC do payment service
type PaymentGRPCServer struct {
	pb.UnimplementedPaymentServiceServer
	paymentService service.PaymentService
}

// NewPaymentGRPCServer cria uma nova instância do servidor gRPC
func NewPaymentGRPCServer(paymentService service.PaymentService) *PaymentGRPCServer {
	return &PaymentGRPCServer{
		paymentService: paymentService,
	}
}

// GetPaymentStatus retorna o status do pagamento para um pedido específico
func (s *PaymentGRPCServer) GetPaymentStatus(ctx context.Context, req *pb.PaymentStatusRequest) (*pb.PaymentStatusResponse, error) {
	payment, err := s.paymentService.GetPaymentByOrderID(uint(req.GetOrderId()))
	if err != nil {
		return &pb.PaymentStatusResponse{
			OrderId: req.GetOrderId(),
			Status:  "NOT_FOUND",
			Amount:  0,
		}, nil
	}

	return &pb.PaymentStatusResponse{
		OrderId: uint32(payment.OrderID),
		Status:  payment.Status,
		Amount:  payment.Amount,
	}, nil
}
