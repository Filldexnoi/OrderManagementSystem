package Usecase

import (
	"awesomeProject/Repo"
	"awesomeProject/entities"
)

type ProductUseCase struct {
	repo Repo.ProductRepoI
}

func NewProductUseCase(repo Repo.ProductRepoI) ProductUseCaseI {
	return &ProductUseCase{repo: repo}
}

func (s *ProductUseCase) CreateProduct(product *entities.Product) (*entities.Product, error) {
	createdProduct, err := s.repo.SaveCreateProduct(product)
	if err != nil {
		return nil, err
	}
	return createdProduct, nil
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

func (s *ProductUseCase) UpdateProduct(product *entities.Product, id uint) (*entities.Product, error) {
	product.ProductId = id
	updatedProduct, err := s.repo.SaveUpdateProduct(product)
	if err != nil {
		return nil, err
	}
	return updatedProduct, nil
}

func (s *ProductUseCase) DeleteProduct(id uint) (*entities.Product, error) {
	deletedProduct, err := s.repo.SaveDeleteProduct(id)
	if err != nil {
		return nil, err
	}
	return deletedProduct, nil
}
