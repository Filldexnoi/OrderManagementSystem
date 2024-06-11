package Usecase

import (
	"awesomeProject/adpter"
	"awesomeProject/entities"
)

type ProductUseCaseI interface {
	CreateProduct(product *entities.Product) error
	GetAllProducts() ([]adpter.ProductBrowserOutput, error)
	GetByIDProduct(id uint) (entities.Product, error)
	UpdateProduct(product *entities.Product, id uint) error
	DeleteProduct(id uint) error
}
