package Usecase

import (
	"awesomeProject/entities"
	"context"
)

type ProductUseCaseI interface {
	CreateProduct(ctx context.Context, product entities.Product) (*entities.Product, error)
	GetAllProducts(ctx context.Context) ([]*entities.Product, error)
	GetByIDProduct(ctx context.Context, id uint) (*entities.Product, error)
	UpdateProduct(ctx context.Context, product entities.Product, id uint) (*entities.Product, error)
	DeleteProduct(ctx context.Context, id uint) (*entities.Product, error)
}
