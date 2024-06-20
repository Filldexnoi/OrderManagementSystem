package Usecase

import (
	"awesomeProject/payload"
)

type OrderUseCaseI interface {
	CreateOrder(order *payload.RequestOrder) error
}
