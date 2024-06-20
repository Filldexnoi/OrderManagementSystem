package Repo

import (
	"awesomeProject/entities"
	"github.com/google/uuid"
)

type OrderRepoI interface {
	SaveCreateOrder(order *entities.Order) error
	GetOrderForUpdateStatus(id uuid.UUID) (*entities.Order, error)
	SaveUpdateStatusOrder(o *entities.Order) error
	SaveGetAllOrders() ([]*entities.Order, error)
}
