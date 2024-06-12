package Usecase

import (
	"awesomeProject/adpter"
	"awesomeProject/entities"
)

type StockUseCaseI interface {
	CreateStock(stock *entities.Stock) error
	GetQtyAllProduct() ([]adpter.CRStockJson, error)
	GetQtyByIDProduct(id uint) (entities.Stock, error)
	UpdateStock(stock *entities.Stock, id uint) error
	DeleteStock(id uint) error
}
