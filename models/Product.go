package models

import (
	"time"
)

type Product struct {
	ProductId        uint `gorm:"primaryKey"`
	ProductTypes     string
	ProductName      string
	ProductPrice     float64
	ProductCreatedAt time.Time
	ProductUpdatedAt time.Time
}
