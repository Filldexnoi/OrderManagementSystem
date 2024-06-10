package models

import (
	"time"
)

type Stock struct {
	ProductId      uint `gorm:"primary_key"`
	QuantitySizeS  uint
	QuantitySizeM  uint
	QuantitySizeL  uint
	QuantitySizeXL uint
	StockCreatedAt time.Time
	StockUpdatedAt time.Time
	Product        Product `gorm:"foreignKey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
