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

func (s *ProductUseCase) CreateProduct(product *entities.Product) error {
	return s.repo.SaveCreateProduct(product)
}

func (s *ProductUseCase) GetAllProducts() ([]entities.Product, error) {
	return s.repo.SaveGetAllProduct()
}

func (s *ProductUseCase) GetByIDProduct(id uint) (entities.Product, error) {
	return s.repo.SaveGetByIDProduct(id)
}

func (s *ProductUseCase) UpdateProduct(product *entities.Product) error {
	return s.repo.SaveUpdateProduct(product)
}

func (s *ProductUseCase) DeleteProduct(id uint) error {
	return s.repo.SaveDeleteProduct(id)
}
