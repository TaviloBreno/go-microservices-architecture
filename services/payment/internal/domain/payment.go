package domain

import "time"

// Payment representa um pagamento no sistema
type Payment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	OrderID   uint      `json:"order_id" gorm:"not null;uniqueIndex"`
	Status    string    `json:"status" gorm:"not null"`
	Amount    float64   `json:"amount" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// TableName retorna o nome da tabela no banco de dados
func (Payment) TableName() string {
	return "payments"
}

// PaymentStatus define os possíveis status de pagamento
const (
	StatusApproved = "approved"
	StatusFailed   = "failed"
	StatusPending  = "pending"
)
