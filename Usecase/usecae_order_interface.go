package Usecase

import (
	"awesomeProject/entities"
	"context"
	"github.com/google/uuid"
)

type OrderUseCaseI interface {
	CreateOrder(ctx context.Context, order entities.Order) (*entities.Order, error)
	UpdateStatusOrder(ctx context.Context, o entities.Order, id uuid.UUID) (*entities.Order, error)
	GetAllOrders(ctx context.Context) ([]*entities.Order, error)
}
