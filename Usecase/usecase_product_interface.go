package Usecase

import (
	"awesomeProject/entities"
)

type ProductUseCaseI interface {
	CreateProduct(product *entities.Product) error
	GetAllProducts() ([]entities.Product, error)
	GetByIDProduct(id uint) (entities.Product, error)
	UpdateProduct(product *entities.Product) error
	DeleteProduct(id uint) error
}
