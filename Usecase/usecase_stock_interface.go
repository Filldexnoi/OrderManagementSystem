package Usecase

import (
	"awesomeProject/payload"
)

type StockUseCaseI interface {
	CreateStock(stock *payload.RequestStock) error
	GetQtyAllProduct() ([]*payload.RespondStock, error)
	GetQtyByIDProduct(id uint) (*payload.RespondStock, error)
	UpdateStock(stock *payload.RequestStock, id uint) error
	DeleteStock(id uint) error
}
