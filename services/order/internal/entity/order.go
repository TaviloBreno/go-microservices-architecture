package entity

import "time"

type Order struct {
	ID         uint    `gorm:"primaryKey;autoIncrement"`
	CustomerID uint    `gorm:"not null"`
	Amount     float64 `gorm:"not null"`
	Status     string  `gorm:"size:20;default:'pending'"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
