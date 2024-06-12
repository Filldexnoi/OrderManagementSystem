package Repo

import (
	"awesomeProject/adpter"
	"awesomeProject/entities"
)

type StockRepoI interface {
	SaveCreateStock(stock *entities.Stock) error
	SaveGetQtyAllProduct() ([]adpter.CRStockJson, error)
	SaveGetQtyByIDProduct(id uint) (entities.Stock, error)
	SaveUpdateStock(stock *entities.Stock, id uint) error
	SaveDeleteStock(id uint) error
}
