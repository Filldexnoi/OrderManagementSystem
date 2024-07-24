package Usecase

import (
	"awesomeProject/Repo"
	"awesomeProject/entities"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestOrderUseCase_CreateOrder(t *testing.T) {
	mockOrderRepo := new(Repo.OrderRepoMock)
	mockTransactionRepo := new(Repo.TransactionRepoMock)
	mockStockRepo := new(Repo.StockRepoMock)
	service := NewOrderUseCase(mockOrderRepo, mockStockRepo, mockTransactionRepo)
	transaction := &entities.Transaction{
		TransactionId: uuid.New(),
		OrderAddress:  "th",
		Items: []entities.Item{
			{ProductId: 1, Quantity: 5, Price: 100},
			{ProductId: 2, Quantity: 3, Price: 200},
		},
		TotalPrice: 1100,
	}
	order := entities.Order{
		OrderId:       uuid.New(),
		TransactionId: transaction.TransactionId,
	}
	orderInitStatus := entities.Order{
		OrderId:       order.OrderId,
		TransactionId: transaction.TransactionId,
		Status:        "New",
	}

	orderInvalidStatusInit := entities.Order{
		OrderId:       uuid.New(),
		TransactionId: transaction.TransactionId,
		Status:        "Processing",
	}

	t.Run("success", func(t *testing.T) {
		mockTransactionRepo.On("GetTransactionToCreateOrder", mock.Anything, order).Return(transaction, nil).Once()
		mockStockRepo.On("CheckStockToCreateOrder", mock.Anything, transaction).Return(nil).Once()
		mockOrderRepo.On("SaveCreateOrder", mock.Anything, orderInitStatus).Return(&orderInitStatus, nil).Once()
		ctx := context.Background()
		result, err := service.CreateOrder(ctx, order)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, &orderInitStatus, result)
		mockOrderRepo.AssertExpectations(t)
		mockTransactionRepo.AssertExpectations(t)
		mockStockRepo.AssertExpectations(t)
	})

	t.Run("Stock not enough", func(t *testing.T) {
		mockTransactionRepo.On("GetTransactionToCreateOrder", mock.Anything, order).Return(transaction, nil).Once()
		mockStockRepo.On("CheckStockToCreateOrder", mock.Anything, transaction).Return(errors.New("stock not enough")).Once()
		ctx := context.Background()
		result, err := service.CreateOrder(ctx, order)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "stock not enough", err.Error())
		mockTransactionRepo.AssertExpectations(t)
		mockStockRepo.AssertExpectations(t)
	})

	t.Run("cannot get transaction", func(t *testing.T) {
		mockTransactionRepo.On("GetTransactionToCreateOrder", mock.Anything, order).Return(nil, errors.New("cannot get transaction")).Once()
		ctx := context.Background()
		result, err := service.CreateOrder(ctx, order)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "cannot get transaction", err.Error())
		mockTransactionRepo.AssertExpectations(t)
	})

	t.Run("Invalid status to init", func(t *testing.T) {
		mockTransactionRepo.On("GetTransactionToCreateOrder", mock.Anything, orderInvalidStatusInit).Return(transaction, nil).Once()
		mockStockRepo.On("CheckStockToCreateOrder", mock.Anything, transaction).Return(nil).Once()
		ctx := context.Background()
		result, err := service.CreateOrder(ctx, orderInvalidStatusInit)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "invalid order status: from Processing to New", err.Error())
		mockTransactionRepo.AssertExpectations(t)
		mockStockRepo.AssertExpectations(t)
	})

	t.Run("Cannot save create order", func(t *testing.T) {
		mockTransactionRepo.On("GetTransactionToCreateOrder", mock.Anything, order).Return(transaction, nil).Once()
		mockStockRepo.On("CheckStockToCreateOrder", mock.Anything, transaction).Return(nil).Once()
		mockOrderRepo.On("SaveCreateOrder", mock.Anything, orderInitStatus).Return(nil, errors.New("cannot save create order")).Once()
		ctx := context.Background()
		result, err := service.CreateOrder(ctx, order)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "cannot save create order", err.Error())
		mockOrderRepo.AssertExpectations(t)
		mockTransactionRepo.AssertExpectations(t)
		mockStockRepo.AssertExpectations(t)
	})
}

func TestOrderUseCase_UpdateStatus(t *testing.T) {
	mockOrderRepo := new(Repo.OrderRepoMock)
	mockTransactionRepo := new(Repo.TransactionRepoMock)
	mockStockRepo := new(Repo.StockRepoMock)
	service := NewOrderUseCase(mockOrderRepo, mockStockRepo, mockTransactionRepo)
	getOrder := &entities.Order{
		OrderId:       uuid.New(),
		TransactionId: uuid.New(),
		Status:        "New",
	}
	order := entities.Order{
		OrderId:       getOrder.OrderId,
		TransactionId: getOrder.TransactionId,
		Status:        "Paid",
	}
	orderInvalidStatus := entities.Order{
		OrderId:       getOrder.OrderId,
		TransactionId: getOrder.TransactionId,
		Status:        "Processing",
	}

	t.Run("success", func(t *testing.T) {
		mockOrderRepo.On("GetOrderForUpdateStatus", mock.Anything, getOrder.OrderId).Return(getOrder, nil).Once()
		mockOrderRepo.On("SaveUpdateStatusOrder", mock.Anything, order).Return(&order, nil).Once()
		ctx := context.Background()
		result, err := service.UpdateStatusOrder(ctx, order, order.OrderId)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, &order, result)
		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("cannot save update status", func(t *testing.T) {
		mockOrderRepo.On("GetOrderForUpdateStatus", mock.Anything, getOrder.OrderId).Return(getOrder, nil).Once()
		mockOrderRepo.On("SaveUpdateStatusOrder", mock.Anything, order).Return(nil, errors.New("cannot save update status")).Once()
		ctx := context.Background()
		result, err := service.UpdateStatusOrder(ctx, order, order.OrderId)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "cannot save update status", err.Error())
		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("cannot get order", func(t *testing.T) {
		mockOrderRepo.On("GetOrderForUpdateStatus", mock.Anything, order.OrderId).Return(nil, errors.New("cannot get order")).Once()
		ctx := context.Background()
		result, err := service.UpdateStatusOrder(ctx, order, order.OrderId)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "cannot get order", err.Error())
		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("cannot change status", func(t *testing.T) {
		mockOrderRepo.On("GetOrderForUpdateStatus", mock.Anything, getOrder.OrderId).Return(getOrder, nil).Once()
		ctx := context.Background()
		result, err := service.UpdateStatusOrder(ctx, orderInvalidStatus, orderInvalidStatus.OrderId)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "invalid o status: from New to Processing", err.Error())
		mockOrderRepo.AssertExpectations(t)
	})
}

func TestOrderUseCase_GetAllOrders(t *testing.T) {
	mockOrderRepo := new(Repo.OrderRepoMock)
	mockTransactionRepo := new(Repo.TransactionRepoMock)
	mockStockRepo := new(Repo.StockRepoMock)
	service := NewOrderUseCase(mockOrderRepo, mockStockRepo, mockTransactionRepo)
	orders := []*entities.Order{
		{OrderId: uuid.New(), TransactionId: uuid.New(), Status: "New"},
		{OrderId: uuid.New(), TransactionId: uuid.New(), Status: "Paid"},
	}
	t.Run("success", func(t *testing.T) {
		mockOrderRepo.On("SaveGetAllOrders", mock.Anything).Return(orders, nil).Once()
		ctx := context.Background()
		result, err := service.GetAllOrders(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, orders, result)
		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("cannot get all orders", func(t *testing.T) {
		mockOrderRepo.On("SaveGetAllOrders", mock.Anything).Return(nil, errors.New("cannot get all orders")).Once()
		ctx := context.Background()
		result, err := service.GetAllOrders(ctx)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "cannot get all orders", err.Error())
		mockOrderRepo.AssertExpectations(t)
	})
}
