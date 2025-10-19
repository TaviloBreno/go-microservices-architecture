package repository

import (
	"notification-service/internal/domain"

	"gorm.io/gorm"
)

// NotificationRepository define os métodos de acesso a dados para notificações
type NotificationRepository interface {
	Create(notification *domain.Notification) error
	GetByID(id uint) (*domain.Notification, error)
	GetByOrderID(orderID uint) ([]*domain.Notification, error)
	GetByPaymentID(paymentID uint) ([]*domain.Notification, error)
	GetAll() ([]*domain.Notification, error)
	Update(notification *domain.Notification) error
	Delete(id uint) error
}

// NotificationRepositoryImpl implementa NotificationRepository usando GORM
type NotificationRepositoryImpl struct {
	db *gorm.DB
}

// NewNotificationRepository cria uma nova instância do repositório
func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &NotificationRepositoryImpl{
		db: db,
	}
}

// Create cria uma nova notificação no banco de dados
func (r *NotificationRepositoryImpl) Create(notification *domain.Notification) error {
	return r.db.Create(notification).Error
}

// GetByID busca uma notificação por ID
func (r *NotificationRepositoryImpl) GetByID(id uint) (*domain.Notification, error) {
	var notification domain.Notification
	err := r.db.First(&notification, id).Error
	if err != nil {
		return nil, err
	}
	return &notification, nil
}

// GetByOrderID busca notificações por Order ID
func (r *NotificationRepositoryImpl) GetByOrderID(orderID uint) ([]*domain.Notification, error) {
	var notifications []*domain.Notification
	err := r.db.Where("order_id = ?", orderID).Find(&notifications).Error
	return notifications, err
}

// GetByPaymentID busca notificações por Payment ID
func (r *NotificationRepositoryImpl) GetByPaymentID(paymentID uint) ([]*domain.Notification, error) {
	var notifications []*domain.Notification
	err := r.db.Where("payment_id = ?", paymentID).Find(&notifications).Error
	return notifications, err
}

// GetAll busca todas as notificações
func (r *NotificationRepositoryImpl) GetAll() ([]*domain.Notification, error) {
	var notifications []*domain.Notification
	err := r.db.Order("created_at DESC").Find(&notifications).Error
	return notifications, err
}

// Update atualiza uma notificação existente
func (r *NotificationRepositoryImpl) Update(notification *domain.Notification) error {
	return r.db.Save(notification).Error
}

// Delete remove uma notificação pelo ID
func (r *NotificationRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&domain.Notification{}, id).Error
}
