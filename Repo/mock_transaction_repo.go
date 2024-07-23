package Repo

import (
	"awesomeProject/entities"
	"context"
	"github.com/stretchr/testify/mock"
)

type TransactionRepoMock struct {
	mock.Mock
}

func (m *TransactionRepoMock) SaveCreateTransaction(ctx context.Context, transaction *entities.Transaction) (*entities.Transaction, error) {
	args := m.Called(ctx, transaction)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Transaction), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *TransactionRepoMock) SaveGetAllTransaction(ctx context.Context) ([]*entities.Transaction, error) {
	args := m.Called(ctx)
	if args.Get(0) != nil {
		return args.Get(0).([]*entities.Transaction), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *TransactionRepoMock) GetTransactionToCreateOrder(ctx context.Context, order entities.Order) (*entities.Transaction, error) {
	args := m.Called(ctx, order)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Transaction), args.Error(1)
	}
	return nil, args.Error(1)
}
