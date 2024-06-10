package models

import (
	"time"
)

type Product struct {
	ProductId        uint      `gorm:"column:product_id;primary_key;auto_increment"`
	ProductTypes     string    `gorm:"column:product_types"`
	ProductName      string    `gorm:"column:product_name"`
	ProductPrice     float64   `gorm:"column:product_price"`
	ProductCreatedAt time.Time `gorm:"column:product_created_at"`
	ProductUpdatedAt time.Time `gorm:"column:product_updated_at"`
}
