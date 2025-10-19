package domain

import (
	"time"

	"gorm.io/gorm"
)

// Notification representa uma notificação no sistema
type Notification struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	PaymentID uint           `gorm:"index" json:"payment_id"`
	OrderID   uint           `gorm:"index" json:"order_id"`
	Message   string         `gorm:"type:varchar(500);not null" json:"message"`
	Status    string         `gorm:"type:varchar(50);not null;default:'SENT'" json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Status possíveis para notificações
const (
	NotificationStatusSent    = "SENT"
	NotificationStatusPending = "PENDING"
	NotificationStatusFailed  = "FAILED"
)

// TableName define o nome da tabela no banco de dados
func (Notification) TableName() string {
	return "notifications"
}

// BeforeCreate hook executado antes de criar uma notificação
func (n *Notification) BeforeCreate(tx *gorm.DB) error {
	if n.Status == "" {
		n.Status = NotificationStatusSent
	}
	return nil
}
