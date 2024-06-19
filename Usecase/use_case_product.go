package Usecase

import (
	"awesomeProject/Repo"
	"awesomeProject/payload"
)

type ProductUseCase struct {
	repo Repo.ProductRepoI
}

func NewProductUseCase(repo Repo.ProductRepoI) ProductUseCaseI {
	return &ProductUseCase{repo: repo}
}

func (s *ProductUseCase) CreateProduct(product *payload.RequestProduct) error {
	productEntity := product.ToProduct()
	return s.repo.SaveCreateProduct(productEntity)
}

func (s *ProductUseCase) GetAllProducts() ([]*payload.RespondProduct, error) {

	products, err := s.repo.SaveGetAllProduct()
	if err != nil {
		return nil, err
	}
	var ResProduct []*payload.RespondProduct
	for _, product := range products {
		ResProduct = append(ResProduct, payload.ProductToRespondProduct(product))
	}
	return ResProduct, nil
}

func (s *ProductUseCase) GetByIDProduct(id uint) (*payload.RespondProduct, error) {
	product, err := s.repo.SaveGetByIDProduct(id)
	if err != nil {
		return nil, err
	}
	ResProduct := payload.ProductToRespondProduct(product)
	return ResProduct, err
}

func (s *ProductUseCase) UpdateProduct(product *payload.RequestProduct, id uint) error {
	productEntity := product.ToProduct()
	productEntity.ProductId = id
	return s.repo.SaveUpdateProduct(productEntity)
}

func (s *ProductUseCase) DeleteProduct(id uint) error {
	return s.repo.SaveDeleteProduct(id)
}
