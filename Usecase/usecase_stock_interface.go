package Usecase

import (
	"awesomeProject/entities"
	"awesomeProject/payload"
)

type StockUseCaseI interface {
	CreateStock(stock *entities.Stock) error
	GetQtyAllProduct() ([]payload.OutgoingStock, error)
	GetQtyByIDProduct(id uint) (entities.Stock, error)
	UpdateStock(stock *entities.Stock, id uint) error
	DeleteStock(id uint) error
}
