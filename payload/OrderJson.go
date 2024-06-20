package payload

import (
	"awesomeProject/entities"
	"github.com/google/uuid"
)

type RequestCreateOrder struct {
	TransactionId uuid.UUID `json:"transaction_id"`
}

type RequestUpdateStatusOrder struct {
	Status string `json:"status"`
}

type ResponseOrder struct {
	OrderId       uuid.UUID `json:"order_id"`
	TransactionId uuid.UUID `json:"transaction_id"`
	Status        string    `json:"status"`
}

func (p *RequestCreateOrder) ToOrder() *entities.Order {
	return &entities.Order{
		TransactionId: p.TransactionId,
	}
}

func (p *RequestUpdateStatusOrder) ToOrder() *entities.Order {
	return &entities.Order{
		Status: p.Status,
	}
}

func OrderToOrderRespond(order *entities.Order) *ResponseOrder {
	return &ResponseOrder{
		OrderId:       order.OrderId,
		TransactionId: order.TransactionId,
		Status:        order.Status,
	}
}
