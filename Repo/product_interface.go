package Repo

import (
	"awesomeProject/entities"
	"context"
)

type ProductRepoI interface {
	SaveCreateProduct(ctx context.Context, product entities.Product) (*entities.Product, error)
	SaveGetAllProduct(ctx context.Context) ([]*entities.Product, error)
	SaveGetByIDProduct(ctx context.Context, id uint) (*entities.Product, error)
	SaveUpdateProduct(ctx context.Context, product entities.Product) (*entities.Product, error)
	SaveDeleteProduct(ctx context.Context, id uint) (*entities.Product, error)
	GetPriceProducts(ctx context.Context, transaction *entities.Transaction) (*entities.Transaction, error)
}
