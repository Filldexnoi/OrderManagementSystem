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

func TestTransactionUseCase_CreateTransaction(t *testing.T) {
	mockProductRepo := new(Repo.ProductRepoMock)
	mockTransactionRepo := new(Repo.TransactionRepoMock)
	service := NewTransactionUseCase(mockTransactionRepo, mockProductRepo)
	transaction := entities.Transaction{
		TransactionId: uuid.New(),
		OrderAddress:  "th",
		Items: []entities.Item{
			{ProductId: 1, Quantity: 5, Price: 100},
			{ProductId: 2, Quantity: 3, Price: 200},
		},
		TotalPrice: 1100,
	}
	transactionInvalidCountry := entities.Transaction{
		TransactionId: uuid.New(),
		OrderAddress:  "xx",
		Items: []entities.Item{
			{ProductId: 1, Quantity: 5, Price: 100},
			{ProductId: 2, Quantity: 3, Price: 200},
		},
		TotalPrice: 1100,
	}
	transactionDuplicateIdProduct := entities.Transaction{
		TransactionId: uuid.New(),
		OrderAddress:  "th",
		Items: []entities.Item{
			{ProductId: 1, Quantity: 5, Price: 100},
			{ProductId: 1, Quantity: 3, Price: 200},
		},
		TotalPrice: 1100,
	}
	t.Run("success", func(t *testing.T) {
		mockProductRepo.On("GetPriceProducts", mock.Anything, &transaction).Return(&transaction, nil).Once()
		mockTransactionRepo.On("SaveCreateTransaction", mock.Anything, &transaction).Return(&transaction, nil).Once()
		ctx := context.Background()
		result, err := service.CreateTransaction(ctx, transaction)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, &transaction, result)
		mockProductRepo.AssertExpectations(t)
		mockTransactionRepo.AssertExpectations(t)
	})

	t.Run("Invalid country", func(t *testing.T) {
		ctx := context.Background()
		result, err := service.CreateTransaction(ctx, transactionInvalidCountry)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "dont have this country", err.Error())
	})

	t.Run("Duplicate Id Product", func(t *testing.T) {
		ctx := context.Background()
		result, err := service.CreateTransaction(ctx, transactionDuplicateIdProduct)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "duplicate product_id", err.Error())
	})

	t.Run("Cannot get PriceProducts", func(t *testing.T) {
		mockProductRepo.On("GetPriceProducts", mock.Anything, &transaction).Return(nil, errors.New("cannot get price products")).Once()
		ctx := context.Background()
		result, err := service.CreateTransaction(ctx, transaction)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "cannot get price products", err.Error())
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("Cannot create transaction", func(t *testing.T) {
		mockProductRepo.On("GetPriceProducts", mock.Anything, &transaction).Return(&transaction, nil).Once()
		mockTransactionRepo.On("SaveCreateTransaction", mock.Anything, &transaction).Return(nil, errors.New("cannot create transaction")).Once()
		ctx := context.Background()
		result, err := service.CreateTransaction(ctx, transaction)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "cannot create transaction", err.Error())
		mockProductRepo.AssertExpectations(t)
		mockTransactionRepo.AssertExpectations(t)
	})
}

func TestTransactionUseCase_GetAllTransaction(t *testing.T) {
	mockProductRepo := new(Repo.ProductRepoMock)
	mockTransactionRepo := new(Repo.TransactionRepoMock)
	service := NewTransactionUseCase(mockTransactionRepo, mockProductRepo)
	transactions := []*entities.Transaction{
		{
			TransactionId: uuid.New(),
			OrderAddress:  "th",
			Items: []entities.Item{
				{ProductId: 1, Quantity: 5, Price: 100},
				{ProductId: 2, Quantity: 3, Price: 200},
			},
			TotalPrice: 1100,
		}, {
			TransactionId: uuid.New(),
			OrderAddress:  "en",
			Items: []entities.Item{
				{ProductId: 1, Quantity: 1, Price: 100},
				{ProductId: 2, Quantity: 1, Price: 200},
			},
			TotalPrice: 1100,
		},
	}
	t.Run("success", func(t *testing.T) {
		mockTransactionRepo.On("SaveGetAllTransaction", mock.Anything).Return(transactions, nil).Once()
		ctx := context.Background()
		result, err := service.GetAllTransaction(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, transactions, result)
		mockTransactionRepo.AssertExpectations(t)
	})

	t.Run("cannot get transactions", func(t *testing.T) {
		mockTransactionRepo.On("SaveGetAllTransaction", mock.Anything).Return(nil, errors.New("cannot get transactions")).Once()
		ctx := context.Background()
		result, err := service.GetAllTransaction(ctx)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "cannot get transactions", err.Error())
		mockTransactionRepo.AssertExpectations(t)
	})
}
