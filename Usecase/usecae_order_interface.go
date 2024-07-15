package Usecase

import (
	"awesomeProject/entities"
	"github.com/google/uuid"
)

type OrderUseCaseI interface {
	CreateOrder(order *entities.Order) (*entities.Order, error)
	UpdateStatusOrder(o *entities.Order, id uuid.UUID) (*entities.Order, error)
	GetAllOrders() ([]*entities.Order, error)
}
