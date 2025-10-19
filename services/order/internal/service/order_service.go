package service

import (
	"errors"
	"log"

	"github.com/seu-usuario/go-microservices-architecture/services/order/internal/domain"
	"github.com/seu-usuario/go-microservices-architecture/services/order/internal/messaging"
	"github.com/seu-usuario/go-microservices-architecture/services/order/internal/repository"
)

// OrderService define a interface para regras de negócio de pedidos
type OrderService interface {
	CreateOrder(customer string, productID uint, quantity int, price float64) (*domain.Order, error)
	ListOrders() ([]domain.Order, error)
	GetOrderByID(id uint) (*domain.Order, error)
}

// orderService implementa OrderService
type orderService struct {
	orderRepo     repository.OrderRepository
	orderPublisher *messaging.OrderPublisher
}

// NewOrderService cria uma nova instância do serviço de pedidos
func NewOrderService(orderRepo repository.OrderRepository, orderPublisher *messaging.OrderPublisher) OrderService {
	return &orderService{
		orderRepo:      orderRepo,
		orderPublisher: orderPublisher,
	}
}

// CreateOrder cria um novo pedido com validações de negócio
func (s *orderService) CreateOrder(customer string, productID uint, quantity int, price float64) (*domain.Order, error) {
	// Validações de negócio
	if customer == "" {
		return nil, errors.New("nome do cliente é obrigatório")
	}

	if productID == 0 {
		return nil, errors.New("ID do produto é obrigatório")
	}

	if quantity <= 0 {
		return nil, errors.New("quantidade deve ser maior que zero")
	}

	if price <= 0 {
		return nil, errors.New("preço deve ser maior que zero")
	}

	// Criar o pedido
	order := &domain.Order{
		Customer:  customer,
		ProductID: productID,
		Quantity:  quantity,
		Price:     price,
	}

	// Salvar no repositório
	if err := s.orderRepo.Create(order); err != nil {
		log.Printf("Erro ao criar pedido: %v", err)
		return nil, errors.New("falha ao salvar pedido")
	}

	log.Printf("✅ Pedido criado com sucesso no banco: ID=%d, Cliente=%s", order.ID, order.Customer)

	// Publicar evento no RabbitMQ
	if s.orderPublisher != nil {
		if err := s.orderPublisher.PublishOrderCreated(order); err != nil {
			// Log do erro mas não falha a criação do pedido
			log.Printf("⚠️ Erro ao publicar evento de pedido criado: %v", err)
		}
	} else {
		log.Println("⚠️ Publisher não configurado - evento não será enviado")
	}

	return order, nil
}

// ListOrders retorna todos os pedidos
func (s *orderService) ListOrders() ([]domain.Order, error) {
	orders, err := s.orderRepo.List()
	if err != nil {
		log.Printf("Erro ao listar pedidos: %v", err)
		return nil, errors.New("falha ao buscar pedidos")
	}

	log.Printf("📋 Listados %d pedidos", len(orders))
	return orders, nil
}

// GetOrderByID retorna um pedido específico pelo ID
func (s *orderService) GetOrderByID(id uint) (*domain.Order, error) {
	if id == 0 {
		return nil, errors.New("ID do pedido é obrigatório")
	}

	order, err := s.orderRepo.GetByID(id)
	if err != nil {
		log.Printf("Erro ao buscar pedido %d: %v", id, err)
		return nil, errors.New("pedido não encontrado")
	}

	return order, nil
}
