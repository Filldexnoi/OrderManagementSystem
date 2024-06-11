package Repo

import (
	"awesomeProject/adpter"
	"awesomeProject/entities"
)

type ProductRepoI interface {
	SaveCreateProduct(product *entities.Product) error
	SaveGetAllProduct() ([]adpter.ProductBrowserOutput, error)
	SaveGetByIDProduct(id uint) (entities.Product, error)
	SaveUpdateProduct(product *entities.Product, id uint) error
	SaveDeleteProduct(id uint) error
}
