package entities

import "time"

type Order struct {
	TransactionId uint
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
