package payload

import "awesomeProject/entities"

type IncomingOrder struct {
	TransactionId uint `json:"transaction_id"`
}

func (p *IncomingOrder) ToEntityOrder() entities.Order {
	return entities.Order{
		TransactionId: p.TransactionId,
	}
}
