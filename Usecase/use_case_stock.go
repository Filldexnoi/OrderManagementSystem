package Usecase

import (
	"awesomeProject/Repo"
	"awesomeProject/payload"
)

type StockUseCase struct {
	repo Repo.StockRepoI
}

func NewStockUseCase(repo Repo.StockRepoI) StockUseCaseI {
	return &StockUseCase{repo: repo}
}

func (s *StockUseCase) CreateStock(stock *payload.RequestStock) error {
	stockEntity := stock.ToStock()
	return s.repo.SaveCreateStock(stockEntity)
}

func (s *StockUseCase) UpdateStock(stock *payload.RequestStock, id uint) error {
	stockEntity := stock.ToStock()
	stockEntity.ProductId = id
	return s.repo.SaveUpdateStock(stockEntity)
}

func (s *StockUseCase) DeleteStock(id uint) error {
	return s.repo.SaveDeleteStock(id)
}

func (s *StockUseCase) GetQtyAllProduct() ([]*payload.RespondStock, error) {
	stocks, err := s.repo.SaveGetQtyAllProduct()
	if err != nil {
		return nil, err
	}
	var ResStocks []*payload.RespondStock
	for _, stock := range stocks {
		ResStocks = append(ResStocks, payload.StockToStockRes(stock))
	}
	return ResStocks, nil
}

func (s *StockUseCase) GetQtyByIDProduct(id uint) (*payload.RespondStock, error) {
	stock, err := s.repo.SaveGetQtyByIDProduct(id)
	if err != nil {
		return nil, err
	}
	ResStock := payload.StockToStockRes(stock)
	return ResStock, err
}
