package Usecase

import (
	"awesomeProject/Repo"
	"awesomeProject/entities"
)

type StockUseCase struct {
	repo Repo.StockRepoI
}

func NewStockUseCase(repo Repo.StockRepoI) StockUseCaseI {
	return &StockUseCase{repo: repo}
}

func (s *StockUseCase) CreateStock(stock *entities.Stock) error {
	return s.repo.SaveCreateStock(stock)
}

func (s *StockUseCase) UpdateStock(stock *entities.Stock, id uint) error {
	stock.ProductId = id
	return s.repo.SaveUpdateStock(stock)
}

func (s *StockUseCase) DeleteStock(id uint) error {
	return s.repo.SaveDeleteStock(id)
}

func (s *StockUseCase) GetQtyAllProduct() ([]*entities.Stock, error) {
	stocks, err := s.repo.SaveGetQtyAllProduct()
	if err != nil {
		return nil, err
	}
	return stocks, nil
}

func (s *StockUseCase) GetQtyByIDProduct(id uint) (*entities.Stock, error) {
	stock, err := s.repo.SaveGetQtyByIDProduct(id)
	if err != nil {
		return nil, err
	}
	return stock, err
}
