package repository

import (
	"payment-service/internal/domain"

	"gorm.io/gorm"
)

// PaymentRepository define a interface para operações de persistência de pagamentos
type PaymentRepository interface {
	Create(payment *domain.Payment) error
	GetByOrderID(orderID uint) (*domain.Payment, error)
	List() ([]domain.Payment, error)
	GetByID(id uint) (*domain.Payment, error)
}

// paymentRepository implementa PaymentRepository
type paymentRepository struct {
	db *gorm.DB
}

// NewPaymentRepository cria uma nova instância do repositório de pagamentos
func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{
		db: db,
	}
}

// Create cria um novo pagamento no banco de dados
func (r *paymentRepository) Create(payment *domain.Payment) error {
	if err := r.db.Create(payment).Error; err != nil {
		return err
	}
	return nil
}

// GetByOrderID retorna um pagamento específico pelo OrderID
func (r *paymentRepository) GetByOrderID(orderID uint) (*domain.Payment, error) {
	var payment domain.Payment
	if err := r.db.Where("order_id = ?", orderID).First(&payment).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

// List retorna todos os pagamentos do banco de dados
func (r *paymentRepository) List() ([]domain.Payment, error) {
	var payments []domain.Payment
	if err := r.db.Order("created_at DESC").Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

// GetByID retorna um pagamento específico pelo ID
func (r *paymentRepository) GetByID(id uint) (*domain.Payment, error) {
	var payment domain.Payment
	if err := r.db.First(&payment, id).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}
