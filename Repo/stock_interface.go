package Repo

import (
	"awesomeProject/entities"
	"context"
)

type StockRepoI interface {
	SaveCreateStock(stock entities.Stock) (*entities.Stock, error)
	SaveGetQtyAllProduct() ([]*entities.Stock, error)
	SaveGetQtyByIDProduct(id uint) (*entities.Stock, error)
	SaveUpdateStock(stock entities.Stock) (*entities.Stock, error)
	SaveDeleteStock(id uint) (*entities.Stock, error)
	CheckStockToCreateOrder(ctx context.Context, transaction *entities.Transaction) error
}
