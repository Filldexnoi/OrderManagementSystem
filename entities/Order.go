package entities

import "github.com/google/uuid"

type Order struct {
	OrderId       uuid.UUID
	TransactionId uuid.UUID
	Status        string
}
