package Repo

import (
	"awesomeProject/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type OrderRepoMock struct {
	mock.Mock
}

func (m *OrderRepoMock) SaveCreateOrder(order *entities.Order) (*entities.Order, error) {
	args := m.Called(order)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Order), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *OrderRepoMock) GetOrderForUpdateStatus(id uuid.UUID) (*entities.Order, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Order), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *OrderRepoMock) SaveUpdateStatusOrder(o *entities.Order) (*entities.Order, error) {
	args := m.Called(o)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Order), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *OrderRepoMock) SaveGetAllOrders() ([]*entities.Order, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*entities.Order), args.Error(1)
	}
	return nil, args.Error(1)
}
