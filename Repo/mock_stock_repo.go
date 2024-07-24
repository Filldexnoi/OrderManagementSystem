package Repo

import (
	"awesomeProject/entities"
	"context"
	"github.com/stretchr/testify/mock"
)

type StockRepoMock struct {
	mock.Mock
}

func (m *StockRepoMock) SaveCreateStock(ctx context.Context, stock entities.Stock) (*entities.Stock, error) {
	args := m.Called(ctx, stock)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Stock), nil
	}
	return nil, args.Error(1)
}

func (m *StockRepoMock) SaveGetQtyAllProduct(ctx context.Context) ([]*entities.Stock, error) {
	args := m.Called(ctx)
	if args.Get(0) != nil {
		return args.Get(0).([]*entities.Stock), nil
	}
	return nil, args.Error(1)
}

func (m *StockRepoMock) SaveGetQtyByIDProduct(ctx context.Context, id uint) (*entities.Stock, error) {
	args := m.Called(ctx, id)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Stock), nil
	}
	return nil, args.Error(1)
}

func (m *StockRepoMock) SaveUpdateStock(ctx context.Context, stock entities.Stock) (*entities.Stock, error) {
	args := m.Called(ctx, stock)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Stock), nil
	}
	return nil, args.Error(1)
}

func (m *StockRepoMock) SaveDeleteStock(ctx context.Context, id uint) (*entities.Stock, error) {
	args := m.Called(ctx, id)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Stock), nil
	}
	return nil, args.Error(1)
}

func (m *StockRepoMock) CheckStockToCreateOrder(ctx context.Context, transaction *entities.Transaction) error {
	args := m.Called(ctx, transaction)
	return args.Error(0)
}
