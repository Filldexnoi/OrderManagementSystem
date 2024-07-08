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

func (s *ProductUseCase) CreateProduct(product *entities.Product) (*entities.Product, error) {
	createdProduct, err := s.repo.SaveCreateProduct(product)
	if err != nil {
		return nil, err
	}
	if !product.IsEqualCreatedProduct(createdProduct) {
		return nil, errors.New("product is not equal to createdProduct")
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
	if product.ProductId != id {
		return nil, errors.New("productID is not equal to id")
	}
	return product, err
}

func (s *ProductUseCase) UpdateProduct(product *entities.Product, id uint) (*entities.Product, error) {
	product.ProductId = id
	updatedProduct, err := s.repo.SaveUpdateProduct(product)
	if err != nil {
		return nil, err
	}
	if !product.IsEqualUpdatedProduct(updatedProduct) {
		return nil, errors.New("product is not equal to updatedProduct")
	}
	return updatedProduct, nil
}

func (s *ProductUseCase) DeleteProduct(id uint) (*entities.Product, error) {
	deletedProduct, err := s.repo.SaveDeleteProduct(id)
	if err != nil {
		return nil, err
	}
	if deletedProduct.ProductId != id {
		return nil, errors.New("deleted productID is not equal to id")
	}
	return deletedProduct, nil
}
