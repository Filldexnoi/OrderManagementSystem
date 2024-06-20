package models

import (
	"awesomeProject/entities"
	"time"
)

type Product struct {
	ProductId        uint      `gorm:"column:product_id;primaryKey;autoIncrement"`
	ProductTypes     string    `gorm:"column:product_types"`
	ProductName      string    `gorm:"column:product_name"`
	ProductPrice     float64   `gorm:"column:product_price"`
	ProductCreatedAt time.Time `gorm:"autoCreateTime"`
	ProductUpdatedAt time.Time `gorm:"autoUpdateTime"`
	Stock            Stock
	Item             Item
}

func (Product) TableName() string {
	return "products"
}

func ProductToGormProduct(p *entities.Product) *Product {
	return &Product{
		ProductId:    p.ProductId,
		ProductTypes: p.ProductTypes,
		ProductName:  p.ProductName,
		ProductPrice: p.ProductPrice,
	}
}

func (p Product) ToProduct() *entities.Product {
	return &entities.Product{
		ProductId:    p.ProductId,
		ProductTypes: p.ProductTypes,
		ProductName:  p.ProductName,
		ProductPrice: p.ProductPrice,
	}
}
