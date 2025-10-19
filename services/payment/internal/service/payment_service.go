package service

import (
	"errors"
	"log"
	"math/rand"

	"payment-service/internal/domain"
	"payment-service/internal/publisher"
	"payment-service/internal/repository"
)

// PaymentService define a interface para regras de negócio de pagamentos
type PaymentService interface {
	ProcessPayment(orderID uint, amount float64) (*domain.Payment, error)
	GetPaymentByOrderID(orderID uint) (*domain.Payment, error)
	ListPayments() ([]domain.Payment, error)
}

// paymentService implementa PaymentService
type paymentService struct {
	paymentRepo repository.PaymentRepository
	publisher   *publisher.PaymentPublisher
}

// NewPaymentService cria uma nova instância do serviço de pagamentos
func NewPaymentService(paymentRepo repository.PaymentRepository, pub *publisher.PaymentPublisher) PaymentService {
	return &paymentService{
		paymentRepo: paymentRepo,
		publisher:   pub,
	}
}

// ProcessPayment processa um pagamento para um pedido
func (s *paymentService) ProcessPayment(orderID uint, amount float64) (*domain.Payment, error) {
	// Validações de negócio
	if orderID == 0 {
		return nil, errors.New("ID do pedido é obrigatório")
	}

	if amount <= 0 {
		return nil, errors.New("valor do pagamento deve ser maior que zero")
	}

	// Verificar se já existe um pagamento para este pedido
	existingPayment, err := s.paymentRepo.GetByOrderID(orderID)
	if err == nil && existingPayment != nil {
		log.Printf("⚠️ Pagamento já existe para OrderID %d: Status=%s", orderID, existingPayment.Status)
		return existingPayment, nil
	}

	// Simular processamento do pagamento (90% de aprovação)
	var status string
	if rand.Float64() < 0.9 {
		status = domain.StatusApproved
		log.Printf("✅ Pagamento simulado - APROVADO para OrderID %d", orderID)
	} else {
		status = domain.StatusFailed
		log.Printf("❌ Pagamento simulado - FALHOU para OrderID %d", orderID)
	}

	// Criar o pagamento
	payment := &domain.Payment{
		OrderID: orderID,
		Status:  status,
		Amount:  amount,
	}

	// Salvar no repositório
	if err := s.paymentRepo.Create(payment); err != nil {
		log.Printf("Erro ao salvar pagamento: %v", err)
		return nil, errors.New("falha ao salvar pagamento")
	}

	// Publicar evento de pagamento processado
	if s.publisher != nil {
		err := s.publisher.PublishPaymentEvent(payment)
		if err != nil {
			log.Printf("⚠️ Erro ao publicar evento de pagamento: %v", err)
			// Não falhar o processamento do pagamento por causa do evento
		}
	}

	log.Printf("💰 Pagamento processado para OrderID: %d - Status: %s - Amount: %.2f",
		orderID, status, amount)

	return payment, nil
}

// GetPaymentByOrderID retorna o pagamento de um pedido específico
func (s *paymentService) GetPaymentByOrderID(orderID uint) (*domain.Payment, error) {
	if orderID == 0 {
		return nil, errors.New("ID do pedido é obrigatório")
	}

	payment, err := s.paymentRepo.GetByOrderID(orderID)
	if err != nil {
		log.Printf("Pagamento não encontrado para OrderID %d: %v", orderID, err)
		return nil, errors.New("pagamento não encontrado")
	}

	return payment, nil
}

// ListPayments retorna todos os pagamentos
func (s *paymentService) ListPayments() ([]domain.Payment, error) {
	payments, err := s.paymentRepo.List()
	if err != nil {
		log.Printf("Erro ao listar pagamentos: %v", err)
		return nil, errors.New("falha ao buscar pagamentos")
	}

	log.Printf("📋 Listados %d pagamentos", len(payments))
	return payments, nil
}
