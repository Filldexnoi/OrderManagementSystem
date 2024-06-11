package models

import (
	"time"
)

type Product struct {
	ProductId        int       `gorm:"column:product_id;primaryKey;autoIncrement"`
	ProductTypes     string    `gorm:"column:product_types"`
	ProductName      string    `gorm:"column:product_name"`
	ProductPrice     float64   `gorm:"column:product_price"`
	ProductCreatedAt time.Time `gorm:"column:product_created_at"`
	ProductUpdatedAt time.Time `gorm:"column:product_updated_at"`
}

func (Product) TableName() string {
	return "products"
}
