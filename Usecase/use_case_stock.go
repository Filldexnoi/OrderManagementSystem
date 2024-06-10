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

func (s *StockUseCase) UpdateStock(stock *entities.Stock) error {
	return s.repo.SaveUpdateStock(stock)
}

func (s *StockUseCase) DeleteStock(id uint) error {
	return s.repo.SaveDeleteStock(id)
}

func (s *StockUseCase) GetQtyAllProduct() ([]entities.Stock, error) {
	return s.repo.SaveGetQtyAllProduct()
}

func (s *StockUseCase) GetQtyByIDProduct(id uint) (entities.Stock, error) {
	return s.repo.SaveGetQtyByIDProduct(id)
}
