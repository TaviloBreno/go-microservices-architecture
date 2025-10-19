package domain

import "time"

// Order representa um pedido no sistema
type Order struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Customer  string    `json:"customer" gorm:"not null"`
	ProductID uint      `json:"product_id" gorm:"not null"`
	Quantity  int       `json:"quantity" gorm:"not null"`
	Price     float64   `json:"price" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// TableName retorna o nome da tabela no banco de dados
func (Order) TableName() string {
	return "orders"
}
