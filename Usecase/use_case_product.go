package Usecase

import (
	"awesomeProject/Repo"
	"awesomeProject/entities"
	"awesomeProject/payload"
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

func (s *ProductUseCase) GetAllProducts() ([]payload.OutgoingProduct, error) {
	return s.repo.SaveGetAllProduct()
}

func (s *ProductUseCase) GetByIDProduct(id uint) (entities.Product, error) {
	return s.repo.SaveGetByIDProduct(id)
}

func (s *ProductUseCase) UpdateProduct(product *entities.Product, id uint) error {
	return s.repo.SaveUpdateProduct(product, id)
}

func (s *ProductUseCase) DeleteProduct(id uint) error {
	return s.repo.SaveDeleteProduct(id)
}
