package Repo

import (
	"awesomeProject/entities"
	"github.com/stretchr/testify/mock"
)

type StockRepoMock struct {
	mock.Mock
}

func (m *StockRepoMock) SaveCreateStock(stock entities.Stock) (*entities.Stock, error) {
	args := m.Called(stock)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Stock), nil
	}
	return nil, args.Error(1)
}

func (m *StockRepoMock) SaveGetQtyAllProduct() ([]*entities.Stock, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*entities.Stock), nil
	}
	return nil, args.Error(1)
}

func (m *StockRepoMock) SaveGetQtyByIDProduct(id uint) (*entities.Stock, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Stock), nil
	}
	return nil, args.Error(1)
}

func (m *StockRepoMock) SaveUpdateStock(stock entities.Stock) (*entities.Stock, error) {
	args := m.Called(stock)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Stock), nil
	}
	return nil, args.Error(1)
}

func (m *StockRepoMock) SaveDeleteStock(id uint) (*entities.Stock, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Stock), nil
	}
	return nil, args.Error(1)
}

func (m *StockRepoMock) CheckStockToCreateOrder(transaction *entities.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}
