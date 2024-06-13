package Repo

import (
	"awesomeProject/entities"
	"awesomeProject/payload"
)

type ProductRepoI interface {
	SaveCreateProduct(product *entities.Product) error
	SaveGetAllProduct() ([]payload.OutgoingProduct, error)
	SaveGetByIDProduct(id uint) (entities.Product, error)
	SaveUpdateProduct(product *entities.Product, id uint) error
	SaveDeleteProduct(id uint) error
}
