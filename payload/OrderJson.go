package payload

import "awesomeProject/entities"

type RequestOrder struct {
	TransactionId uint `json:"transaction_id"`
}

func (p *RequestOrder) ToEntityOrder() entities.Order {
	return entities.Order{
		TransactionId: p.TransactionId,
	}
}
