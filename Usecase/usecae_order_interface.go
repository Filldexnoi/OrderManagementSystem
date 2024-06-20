package Usecase

import (
	"awesomeProject/payload"
	"github.com/google/uuid"
)

type OrderUseCaseI interface {
	CreateOrder(order *payload.RequestCreateOrder) error
	UpdateStatusOrder(o *payload.RequestUpdateStatusOrder, id uuid.UUID) error
	GetAllOrders() ([]*payload.ResponseOrder, error)
}
