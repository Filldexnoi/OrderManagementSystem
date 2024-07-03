package Usecase

import (
	"awesomeProject/entities"
	"github.com/google/uuid"
)

type OrderUseCaseI interface {
	CreateOrder(order *entities.Order) error
	UpdateStatusOrder(o *entities.Order, id uuid.UUID) error
	GetAllOrders() ([]*entities.Order, error)
}
