package Usecase

import (
	"awesomeProject/entities"
	"context"
)

type StockUseCaseI interface {
	CreateStock(ctx context.Context, stock entities.Stock) (*entities.Stock, error)
	GetQtyAllProduct(ctx context.Context) ([]*entities.Stock, error)
	GetQtyByIDProduct(ctx context.Context, id uint) (*entities.Stock, error)
	UpdateStock(ctx context.Context, stock entities.Stock, id uint) (*entities.Stock, error)
	DeleteStock(ctx context.Context, id uint) (*entities.Stock, error)
}
