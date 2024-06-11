package entities

import (
	"time"
)

type Product struct {
	ProductId        uint
	ProductTypes     string
	ProductName      string
	ProductPrice     float64
	ProductCreatedAt time.Time
	ProductUpdatedAt time.Time
}
