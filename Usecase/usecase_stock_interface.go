package Usecase

import (
	"awesomeProject/entities"
)

type StockUseCaseI interface {
	CreateStock(stock entities.Stock) (*entities.Stock, error)
	GetQtyAllProduct() ([]*entities.Stock, error)
	GetQtyByIDProduct(id uint) (*entities.Stock, error)
	UpdateStock(stock entities.Stock, id uint) (*entities.Stock, error)
	DeleteStock(id uint) (*entities.Stock, error)
}
