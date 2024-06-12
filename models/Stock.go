package models

import (
	"time"
)

type Stock struct {
	ProductId      uint `gorm:"column:product_id;primaryKey;"`
	QuantitySizeS  uint `gorm:"column:quantity_size_s"`
	QuantitySizeM  uint `gorm:"column:quantity_size_m"`
	QuantitySizeL  uint `gorm:"column:quantity_size_l"`
	QuantitySizeXL uint `gorm:"column:quantity_size_xl"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
