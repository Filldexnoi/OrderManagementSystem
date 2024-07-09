package Repo

import (
	"awesomeProject/entities"
	"github.com/stretchr/testify/mock"
)

type TransactionRepoMock struct {
	mock.Mock
}

func (m *TransactionRepoMock) SaveCreateTransaction(transaction *entities.Transaction) (*entities.Transaction, error) {
	args := m.Called(transaction)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Transaction), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *TransactionRepoMock) SaveGetAllTransaction() ([]*entities.Transaction, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*entities.Transaction), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *TransactionRepoMock) GetTransactionToCreateOrder(order *entities.Order) (*entities.Transaction, error) {
	args := m.Called(order)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Transaction), args.Error(1)
	}
	return nil, args.Error(1)
}
