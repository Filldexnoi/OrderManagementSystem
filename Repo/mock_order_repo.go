package Repo

import (
	"awesomeProject/entities"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type OrderRepoMock struct {
	mock.Mock
}

func (m *OrderRepoMock) SaveCreateOrder(ctx context.Context, order entities.Order) (*entities.Order, error) {
	args := m.Called(ctx, order)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Order), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *OrderRepoMock) GetOrderForUpdateStatus(ctx context.Context, id uuid.UUID) (*entities.Order, error) {
	args := m.Called(ctx, id)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Order), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *OrderRepoMock) SaveUpdateStatusOrder(ctx context.Context, o entities.Order) (*entities.Order, error) {
	args := m.Called(ctx, o)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Order), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *OrderRepoMock) SaveGetAllOrders(ctx context.Context) ([]*entities.Order, error) {
	args := m.Called(ctx)
	if args.Get(0) != nil {
		return args.Get(0).([]*entities.Order), args.Error(1)
	}
	return nil, args.Error(1)
}
