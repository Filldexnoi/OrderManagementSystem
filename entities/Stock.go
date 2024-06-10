package entities

import "time"

type Stock struct {
	ProductId      uint
	QuantitySizeS  uint
	QuantitySizeM  uint
	QuantitySizeL  uint
	QuantitySizeXL uint
	StockCreatedAt time.Time
	StockUpdatedAt time.Time
}
