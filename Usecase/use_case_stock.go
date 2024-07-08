package Usecase

import (
	"awesomeProject/Repo"
	"awesomeProject/entities"
	"errors"
)

type StockUseCase struct {
	repo Repo.StockRepoI
}

func NewStockUseCase(repo Repo.StockRepoI) StockUseCaseI {
	return &StockUseCase{repo: repo}
}

func (s *StockUseCase) CreateStock(stock *entities.Stock) (*entities.Stock, error) {
	createStock, err := s.repo.SaveCreateStock(stock)
	if err != nil {
		return nil, err
	}
	if !stock.IsEqualCreatedStock(createStock) {
		return nil, errors.New("stock is not equal to createdStock")
	}
	return createStock, nil
}

func (s *StockUseCase) UpdateStock(stock *entities.Stock, id uint) (*entities.Stock, error) {
	stock.ProductId = id
	updatedStock, err := s.repo.SaveUpdateStock(stock)
	if err != nil {
		return nil, err
	}
	if !stock.IsEqualUpdatedStock(updatedStock) {
		return nil, errors.New("stock is not equal to updatedStock")
	}
	return updatedStock, nil
}

func (s *StockUseCase) DeleteStock(id uint) (*entities.Stock, error) {
	deletedStock, err := s.repo.SaveDeleteStock(id)
	if err != nil {
		return nil, err
	}
	if deletedStock.ProductId != id {
		return nil, errors.New("id is not equal to deletedStock id")
	}
	return deletedStock, nil
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
	if stock.ProductId != id {
		return nil, errors.New("id is not equal to stock id")
	}
	return stock, err
}
