package entities

import "time"

type Order struct {
	OrderId       uint
	TransactionId uint
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
