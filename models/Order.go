package models

import (
	"time"
)

type Order struct {
	TransactionID uint `gorm:"column:order_id;primaryKey"`
	Transaction   Transaction
	Status        string    `gorm:"column:status"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
}
