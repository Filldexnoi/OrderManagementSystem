package Repo

import (
	"awesomeProject/entities"
	"github.com/stretchr/testify/mock"
)

type ProductRepoMock struct {
	mock.Mock
}

func (m *ProductRepoMock) SaveCreateProduct(product entities.Product) (*entities.Product, error) {
	arguments := m.Called(product)
	if arguments.Get(0) != nil {
		return arguments.Get(0).(*entities.Product), arguments.Error(1)
	}
	return nil, arguments.Error(1)
}

func (m *ProductRepoMock) SaveGetAllProduct() ([]*entities.Product, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*entities.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *ProductRepoMock) SaveGetByIDProduct(id uint) (*entities.Product, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *ProductRepoMock) SaveUpdateProduct(product entities.Product) (*entities.Product, error) {
	args := m.Called(product)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *ProductRepoMock) SaveDeleteProduct(id uint) (*entities.Product, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *ProductRepoMock) GetPriceProducts(transaction *entities.Transaction) (*entities.Transaction, error) {
	args := m.Called(transaction)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Transaction), args.Error(1)
	}
	return nil, args.Error(1)
}
