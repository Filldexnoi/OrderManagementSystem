package Repo

import (
	"awesomeProject/entities"
	"context"
	"github.com/google/uuid"
)

type OrderRepoI interface {
	SaveCreateOrder(ctx context.Context, order entities.Order) (*entities.Order, error)
	GetOrderForUpdateStatus(ctx context.Context, id uuid.UUID) (*entities.Order, error)
	SaveUpdateStatusOrder(ctx context.Context, o entities.Order) (*entities.Order, error)
	SaveGetAllOrders(ctx context.Context) ([]*entities.Order, error)
}
