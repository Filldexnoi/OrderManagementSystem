package models

import (
	"time"
)

type Stock struct {
	Product   Product `gorm:"column:product_id;primaryKey;"`
	Quantity  uint    `gorm:"column:quantity"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
