package Usecase

import (
	"awesomeProject/entities"
	"awesomeProject/payload"
)

type ProductUseCaseI interface {
	CreateProduct(product *entities.Product) error
	GetAllProducts() ([]payload.OutgoingProduct, error)
	GetByIDProduct(id uint) (entities.Product, error)
	UpdateProduct(product *entities.Product, id uint) error
	DeleteProduct(id uint) error
}
