package Repo

import (
	"awesomeProject/entities"
)

type ProductRepoI interface {
	SaveCreateProduct(product entities.Product) (*entities.Product, error)
	SaveGetAllProduct() ([]*entities.Product, error)
	SaveGetByIDProduct(id uint) (*entities.Product, error)
	SaveUpdateProduct(product entities.Product) (*entities.Product, error)
	SaveDeleteProduct(id uint) (*entities.Product, error)
	GetPriceProducts(transaction *entities.Transaction) (*entities.Transaction, error)
}
