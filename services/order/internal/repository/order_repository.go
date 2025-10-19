package repository

import (
	"github.com/seu-usuario/go-microservices-architecture/services/order/internal/domain"
	"gorm.io/gorm"
)

// OrderRepository define a interface para operações de persistência de pedidos
type OrderRepository interface {
	Create(order *domain.Order) error
	List() ([]domain.Order, error)
	GetByID(id uint) (*domain.Order, error)
}

// orderRepository implementa OrderRepository
type orderRepository struct {
	db *gorm.DB
}

// NewOrderRepository cria uma nova instância do repositório de pedidos
func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

// Create cria um novo pedido no banco de dados
func (r *orderRepository) Create(order *domain.Order) error {
	if err := r.db.Create(order).Error; err != nil {
		return err
	}
	return nil
}

// List retorna todos os pedidos do banco de dados
func (r *orderRepository) List() ([]domain.Order, error) {
	var orders []domain.Order
	if err := r.db.Order("created_at DESC").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// GetByID retorna um pedido específico pelo ID
func (r *orderRepository) GetByID(id uint) (*domain.Order, error) {
	var order domain.Order
	if err := r.db.First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
