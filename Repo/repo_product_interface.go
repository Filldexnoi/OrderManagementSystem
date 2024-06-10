package Repo

import (
	"awesomeProject/entities"
)

type Repo interface {
	SaveCreateProduct(product *entities.Product) error
	SaveGetAllProduct() ([]entities.Product, error)
	SaveGetByIDProduct(id uint) (entities.Product, error)
	SaveUpdateProduct(product *entities.Product) error
	SaveDeleteProduct(id uint) error
}
