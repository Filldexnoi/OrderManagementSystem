package models

import (
	"time"
)

type Order struct {
	Transaction Transaction `gorm:"column:transaction_id;primaryKey;"`
	Status      string      `gorm:"column:status"`
	CreatedAt   time.Time   `gorm:"column:created_at"`
	UpdatedAt   time.Time   `gorm:"column:updated_at"`
}
