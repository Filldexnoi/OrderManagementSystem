package Repo

import (
	"awesomeProject/entities"
	"context"
	"github.com/stretchr/testify/mock"
)

type ProductRepoMock struct {
	mock.Mock
}

func (m *ProductRepoMock) SaveCreateProduct(ctx context.Context, product entities.Product) (*entities.Product, error) {
	arguments := m.Called(ctx, product)
	if arguments.Get(0) != nil {
		return arguments.Get(0).(*entities.Product), arguments.Error(1)
	}
	return nil, arguments.Error(1)
}

func (m *ProductRepoMock) SaveGetAllProduct(ctx context.Context) ([]*entities.Product, error) {
	args := m.Called(ctx)
	if args.Get(0) != nil {
		return args.Get(0).([]*entities.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *ProductRepoMock) SaveGetByIDProduct(ctx context.Context, id uint) (*entities.Product, error) {
	args := m.Called(ctx, id)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *ProductRepoMock) SaveUpdateProduct(ctx context.Context, product entities.Product) (*entities.Product, error) {
	args := m.Called(ctx, product)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *ProductRepoMock) SaveDeleteProduct(ctx context.Context, id uint) (*entities.Product, error) {
	args := m.Called(ctx, id)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *ProductRepoMock) GetPriceProducts(ctx context.Context, transaction *entities.Transaction) (*entities.Transaction, error) {
	args := m.Called(ctx, transaction)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Transaction), args.Error(1)
	}
	return nil, args.Error(1)
}
