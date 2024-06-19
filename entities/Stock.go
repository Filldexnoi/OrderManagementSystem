package entities

import "time"

type Stock struct {
	ProductId      uint
	Quantity       uint
	StockCreatedAt time.Time
	StockUpdatedAt time.Time
}
