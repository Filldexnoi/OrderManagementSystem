package entities

import "time"

type Order struct {
	OrderId      uint
	OrderAddress string
	Items        []Item
	Status       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Item struct {
	ProductId uint
	Quantity  uint
	Price     float64
}
