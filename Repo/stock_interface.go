package Repo

import (
	"awesomeProject/entities"
)

type StockRepoI interface {
	SaveCreateStock(stock *entities.Stock) error
	SaveGetQtyAllProduct() ([]*entities.Stock, error)
	SaveGetQtyByIDProduct(id uint) (*entities.Stock, error)
	SaveUpdateStock(stock *entities.Stock) error
	SaveDeleteStock(id uint) error
}
