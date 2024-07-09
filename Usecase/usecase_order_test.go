package Usecase

import (
	"awesomeProject/Repo"
	"awesomeProject/entities"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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
	order := &entities.Order{
		OrderId:       uuid.New(),
		TransactionId: transaction.TransactionId,
	}

	orderInvalidStatusInit := &entities.Order{
		OrderId:       uuid.New(),
		TransactionId: transaction.TransactionId,
		Status:        "Processing",
	}

	t.Run("success", func(t *testing.T) {
		mockTransactionRepo.On("GetTransactionToCreateOrder", order).Return(transaction, nil).Once()
		mockStockRepo.On("CheckStockToCreateOrder", transaction).Return(nil).Once()
		mockOrderRepo.On("SaveCreateOrder", order).Return(order, nil).Once()
		result, err := service.CreateOrder(order)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, order.TransactionId, result.TransactionId)
		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("Stock not enough", func(t *testing.T) {
		mockTransactionRepo.On("GetTransactionToCreateOrder", order).Return(transaction, nil).Once()
		mockStockRepo.On("CheckStockToCreateOrder", transaction).Return(errors.New("stock not enough")).Once()
		result, err := service.CreateOrder(order)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "stock not enough", err.Error())
		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("cannot get transaction", func(t *testing.T) {
		mockTransactionRepo.On("GetTransactionToCreateOrder", order).Return(nil, errors.New("cannot get transaction")).Once()
		result, err := service.CreateOrder(order)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "cannot get transaction", err.Error())
		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("Invalid status to init", func(t *testing.T) {
		mockTransactionRepo.On("GetTransactionToCreateOrder", orderInvalidStatusInit).Return(transaction, nil).Once()
		mockStockRepo.On("CheckStockToCreateOrder", transaction).Return(nil).Once()
		result, err := service.CreateOrder(orderInvalidStatusInit)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "invalid order status: from Processing to New", err.Error())
		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("Cannot save create order", func(t *testing.T) {
		mockTransactionRepo.On("GetTransactionToCreateOrder", order).Return(transaction, nil).Once()
		mockStockRepo.On("CheckStockToCreateOrder", transaction).Return(nil).Once()
		mockOrderRepo.On("SaveCreateOrder", order).Return(nil, errors.New("cannot save create order")).Once()
		result, err := service.CreateOrder(order)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "cannot save create order", err.Error())
		mockOrderRepo.AssertExpectations(t)
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
	order := &entities.Order{
		OrderId:       getOrder.OrderId,
		TransactionId: getOrder.TransactionId,
		Status:        "Paid",
	}
	orderInvalidStatus := &entities.Order{
		OrderId:       getOrder.OrderId,
		TransactionId: getOrder.TransactionId,
		Status:        "Processing",
	}

	t.Run("success", func(t *testing.T) {
		mockOrderRepo.On("GetOrderForUpdateStatus", getOrder.OrderId).Return(getOrder, nil)
		mockOrderRepo.On("SaveUpdateStatusOrder", order).Return(order, nil).Once()
		result, err := service.UpdateStatusOrder(order, order.OrderId)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, order, result)
		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("cannot save update status", func(t *testing.T) {
		mockOrderRepo.On("GetOrderForUpdateStatus", getOrder.OrderId).Return(getOrder, nil)
		mockOrderRepo.On("SaveUpdateStatusOrder", order).Return(nil, errors.New("cannot save update status")).Once()
		result, err := service.UpdateStatusOrder(order, order.OrderId)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "cannot save update status", err.Error())
		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("cannot get order", func(t *testing.T) {
		mockOrderRepo.On("GetOrderForUpdateStatus", order.OrderId).Return(nil, errors.New("cannot get order")).Once()
		result, err := service.UpdateStatusOrder(order, order.OrderId)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "cannot get order", err.Error())
		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("cannot change status", func(t *testing.T) {
		mockOrderRepo.On("GetOrderForUpdateStatus", getOrder.OrderId).Return(getOrder, nil).Once()
		result, err := service.UpdateStatusOrder(orderInvalidStatus, orderInvalidStatus.OrderId)
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
		mockOrderRepo.On("SaveGetAllOrders").Return(orders, nil).Once()
		result, err := service.GetAllOrders()
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, orders, result)
		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("cannot get all orders", func(t *testing.T) {
		mockOrderRepo.On("SaveGetAllOrders").Return(nil, errors.New("cannot get all orders")).Once()
		result, err := service.GetAllOrders()
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "cannot get all orders", err.Error())
		mockOrderRepo.AssertExpectations(t)
	})
}
