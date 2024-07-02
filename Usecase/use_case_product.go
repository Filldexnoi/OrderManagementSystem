package Usecase

import (
	"awesomeProject/Repo"
	"awesomeProject/entities"
	"errors"
)

type ProductUseCase struct {
	repo Repo.ProductRepoI
}

func NewProductUseCase(repo Repo.ProductRepoI) ProductUseCaseI {
	return &ProductUseCase{repo: repo}
}

func (s *ProductUseCase) CreateProduct(product *entities.Product) error {
	return s.repo.SaveCreateProduct(product)
}

func (s *ProductUseCase) GetAllProducts() ([]*entities.Product, error) {
	products, err := s.repo.SaveGetAllProduct()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductUseCase) GetByIDProduct(id uint) (*entities.Product, error) {
	product, err := s.repo.SaveGetByIDProduct(id)
	if err != nil {
		return nil, err
	}
	return product, err
}

func (s *ProductUseCase) UpdateProduct(product *entities.Product, id uint) error {
	product.ProductId = id
	if product.ProductPrice < 0 {
		return errors.New("product price must be not negative")
	}
	return s.repo.SaveUpdateProduct(product)
}

func (s *ProductUseCase) DeleteProduct(id uint) error {
	return s.repo.SaveDeleteProduct(id)
}
