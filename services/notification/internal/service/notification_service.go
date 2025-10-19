package service

import (
	"fmt"
	"log"
	"time"

	"notification-service/internal/domain"
	"notification-service/internal/repository"
)

// NotificationService define os métodos de negócio para notificações
type NotificationService interface {
	ProcessPaymentNotification(paymentID uint, orderID uint, status string, amount float64) error
	GetAllNotifications() ([]*domain.Notification, error)
	GetNotificationsByOrderID(orderID uint) ([]*domain.Notification, error)
	GetNotificationsByPaymentID(paymentID uint) ([]*domain.Notification, error)
}

// NotificationServiceImpl implementa NotificationService
type NotificationServiceImpl struct {
	repo repository.NotificationRepository
}

// NewNotificationService cria uma nova instância do service
func NewNotificationService(repo repository.NotificationRepository) NotificationService {
	return &NotificationServiceImpl{
		repo: repo,
	}
}

// ProcessPaymentNotification processa uma notificação de pagamento
func (s *NotificationServiceImpl) ProcessPaymentNotification(paymentID uint, orderID uint, status string, amount float64) error {
	// Gerar mensagem baseada no status do pagamento
	var message string
	var notificationStatus string

	switch status {
	case "APPROVED":
		message = fmt.Sprintf("💬 Notificação: Pagamento aprovado para o pedido %d no valor de R$ %.2f - Enviada com sucesso!", orderID, amount)
		notificationStatus = domain.NotificationStatusSent
		log.Printf("✅ %s", message)
	case "REJECTED":
		message = fmt.Sprintf("❌ Notificação: Pagamento rejeitado para o pedido %d no valor de R$ %.2f", orderID, amount)
		notificationStatus = domain.NotificationStatusSent
		log.Printf("❌ %s", message)
	default:
		message = fmt.Sprintf("📝 Notificação: Status de pagamento '%s' para o pedido %d no valor de R$ %.2f", status, orderID, amount)
		notificationStatus = domain.NotificationStatusSent
		log.Printf("📝 %s", message)
	}

	// Criar notificação
	notification := &domain.Notification{
		PaymentID: paymentID,
		OrderID:   orderID,
		Message:   message,
		Status:    notificationStatus,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Salvar no banco de dados
	err := s.repo.Create(notification)
	if err != nil {
		log.Printf("❌ Erro ao salvar notificação no banco: %v", err)
		return fmt.Errorf("falha ao salvar notificação: %w", err)
	}

	log.Printf("✅ Notificação salva no banco de dados - ID: %d", notification.ID)

	// Simular envio da notificação (em um cenário real, aqui seria o envio por email, SMS, etc.)
	s.simulateNotificationSending(notification)

	return nil
}

// GetAllNotifications retorna todas as notificações
func (s *NotificationServiceImpl) GetAllNotifications() ([]*domain.Notification, error) {
	return s.repo.GetAll()
}

// GetNotificationsByOrderID retorna notificações por Order ID
func (s *NotificationServiceImpl) GetNotificationsByOrderID(orderID uint) ([]*domain.Notification, error) {
	return s.repo.GetByOrderID(orderID)
}

// GetNotificationsByPaymentID retorna notificações por Payment ID
func (s *NotificationServiceImpl) GetNotificationsByPaymentID(paymentID uint) ([]*domain.Notification, error) {
	return s.repo.GetByPaymentID(paymentID)
}

// simulateNotificationSending simula o envio da notificação
func (s *NotificationServiceImpl) simulateNotificationSending(notification *domain.Notification) {
	log.Printf("📧 Simulando envio de notificação...")
	log.Printf("📱 Canal: Email/SMS")
	log.Printf("📋 Conteúdo: %s", notification.Message)
	log.Printf("⏰ Enviado em: %s", notification.CreatedAt.Format("2006-01-02 15:04:05"))
	log.Printf("📨 Notificação enviada com sucesso!")
}
