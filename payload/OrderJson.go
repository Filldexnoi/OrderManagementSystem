package payload

import (
	"awesomeProject/entities"
	"github.com/google/uuid"
)

type RequestOrder struct {
	TransactionId uuid.UUID `json:"transaction_id"`
}

type ResponseOrder struct {
	OrderId       uuid.UUID `json:"order_id"`
	TransactionId uuid.UUID `json:"transaction_id"`
	Status        string    `json:"status"`
}

func (p *RequestOrder) ToOrder() *entities.Order {
	return &entities.Order{
		TransactionId: p.TransactionId,
	}
}
