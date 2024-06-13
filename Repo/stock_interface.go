package Repo

import (
	"awesomeProject/entities"
	"awesomeProject/payload"
)

type StockRepoI interface {
	SaveCreateStock(stock *entities.Stock) error
	SaveGetQtyAllProduct() ([]payload.OutgoingStock, error)
	SaveGetQtyByIDProduct(id uint) (entities.Stock, error)
	SaveUpdateStock(stock *entities.Stock, id uint) error
	SaveDeleteStock(id uint) error
}
