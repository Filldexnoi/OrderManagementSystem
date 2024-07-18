package Repo

import (
	"awesomeProject/entities"
	"context"
)

type StockRepoI interface {
	SaveCreateStock(ctx context.Context, stock entities.Stock) (*entities.Stock, error)
	SaveGetQtyAllProduct(ctx context.Context) ([]*entities.Stock, error)
	SaveGetQtyByIDProduct(ctx context.Context, id uint) (*entities.Stock, error)
	SaveUpdateStock(ctx context.Context, stock entities.Stock) (*entities.Stock, error)
	SaveDeleteStock(ctx context.Context, id uint) (*entities.Stock, error)
	CheckStockToCreateOrder(transaction *entities.Transaction) error
}
