package Usecase

import (
	"awesomeProject/payload"
)

type ProductUseCaseI interface {
	CreateProduct(product *payload.RequestProduct) error
	GetAllProducts() ([]*payload.RespondProduct, error)
	GetByIDProduct(id uint) (*payload.RespondProduct, error)
	UpdateProduct(product *payload.RequestProduct, id uint) error
	DeleteProduct(id uint) error
}
