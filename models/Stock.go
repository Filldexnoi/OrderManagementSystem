package models

import (
	"awesomeProject/entities"
	"time"
)

type Stock struct {
	ProductID uint      `gorm:"primary_key"`
	Quantity  uint      `gorm:"column:quantity"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Product   *Product
}

func (Stock) TableName() string {
	return "stocks"
}

func (s Stock) ToStock() entities.Stock {
	return entities.Stock{
		ProductId: s.ProductID,
		Quantity:  s.Quantity,
	}
}

func StockToGormStock(s entities.Stock) Stock {
	return Stock{
		ProductID: s.ProductId,
		Quantity:  s.Quantity,
	}
}
