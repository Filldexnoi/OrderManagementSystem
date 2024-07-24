package Usecase

import (
	"awesomeProject/Repo"
	"awesomeProject/entities"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestStockUseCase_CreateStock(t *testing.T) {
	mockStockRepo := new(Repo.StockRepoMock)
	service := NewStockUseCase(mockStockRepo)

	stock := entities.Stock{ProductId: 1, Quantity: 20}
	t.Run("successful create stock", func(t *testing.T) {
		mockStockRepo.On("SaveCreateStock", mock.Anything, stock).Return(&stock, nil).Once()
		ctx := context.Background()
		result, err := service.CreateStock(ctx, stock)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, &stock, result)
		mockStockRepo.AssertExpectations(t)
	})

	t.Run("fail : cannot createStock", func(t *testing.T) {
		mockStockRepo.On("SaveCreateStock", mock.Anything, stock).Return(nil, errors.New("cannot create stock")).Once()
		ctx := context.Background()
		result, err := service.CreateStock(ctx, stock)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "cannot create stock", err.Error())
		mockStockRepo.AssertExpectations(t)
	})
}

func TestStockUseCase_GetQtyAllProduct(t *testing.T) {
	mockStockRepo := new(Repo.StockRepoMock)
	service := NewStockUseCase(mockStockRepo)

	stocks := []*entities.Stock{
		{ProductId: 1, Quantity: 20},
		{ProductId: 2, Quantity: 30},
	}
	t.Run("successful get qty all product", func(t *testing.T) {
		mockStockRepo.On("SaveGetQtyAllProduct", mock.Anything).Return(stocks, nil).Once()
		ctx := context.Background()
		results, err := service.GetQtyAllProduct(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, results)
		assert.Equal(t, stocks, results)
		mockStockRepo.AssertExpectations(t)
	})

	t.Run("fail : cannot GetQtyAllProduct", func(t *testing.T) {
		mockStockRepo.On("SaveGetQtyAllProduct", mock.Anything).Return(nil, errors.New("cannot get all stock")).Once()
		ctx := context.Background()
		results, err := service.GetQtyAllProduct(ctx)
		assert.Error(t, err)
		assert.Nil(t, results)
		assert.Equal(t, "cannot get all stock", err.Error())
		mockStockRepo.AssertExpectations(t)
	})
}

func TestStockUseCase_GetQtyByIDProduct(t *testing.T) {
	mockStockRepo := new(Repo.StockRepoMock)
	service := NewStockUseCase(mockStockRepo)

	stock := &entities.Stock{ProductId: 1, Quantity: 20}
	t.Run("successful get stock by id", func(t *testing.T) {
		mockStockRepo.On("SaveGetQtyByIDProduct", mock.Anything, stock.ProductId).Return(stock, nil).Once()
		ctx := context.Background()
		result, err := service.GetQtyByIDProduct(ctx, stock.ProductId)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, stock, result)
		mockStockRepo.AssertExpectations(t)
	})

	t.Run("fail : cannot GetQtyByIDProduct", func(t *testing.T) {
		mockStockRepo.On("SaveGetQtyByIDProduct", mock.Anything, stock.ProductId).Return(nil, errors.New("cannot get stock by id")).Once()
		ctx := context.Background()
		result, err := service.GetQtyByIDProduct(ctx, stock.ProductId)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "cannot get stock by id", err.Error())
		mockStockRepo.AssertExpectations(t)
	})
}

func TestStockUseCase_UpdateStock(t *testing.T) {
	mockStockRepo := new(Repo.StockRepoMock)
	service := NewStockUseCase(mockStockRepo)

	stock := entities.Stock{ProductId: 1, Quantity: 20}
	t.Run("success", func(t *testing.T) {
		mockStockRepo.On("SaveUpdateStock", mock.Anything, stock).Return(&stock, nil).Once()
		ctx := context.Background()
		result, err := service.UpdateStock(ctx, stock, stock.ProductId)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, &stock, result)
		mockStockRepo.AssertExpectations(t)
	})

	t.Run("Cannot UpdateStock", func(t *testing.T) {
		mockStockRepo.On("SaveUpdateStock", mock.Anything, stock).Return(nil, errors.New("cannot update stock")).Once()
		ctx := context.Background()
		result, err := service.UpdateStock(ctx, stock, stock.ProductId)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "cannot update stock", err.Error())
		mockStockRepo.AssertExpectations(t)
	})
}

func TestStockUseCase_DeleteStock(t *testing.T) {
	mockStockRepo := new(Repo.StockRepoMock)
	service := NewStockUseCase(mockStockRepo)
	stock := &entities.Stock{ProductId: 1, Quantity: 20}
	t.Run("success", func(t *testing.T) {
		mockStockRepo.On("SaveDeleteStock", mock.Anything, stock.ProductId).Return(stock, nil).Once()
		ctx := context.Background()
		result, err := service.DeleteStock(ctx, stock.ProductId)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, stock, result)
		mockStockRepo.AssertExpectations(t)
	})
	t.Run("Cannot DeleteStock", func(t *testing.T) {
		mockStockRepo.On("SaveDeleteStock", mock.Anything, stock.ProductId).Return(nil, errors.New("cannot delete stock")).Once()
		ctx := context.Background()
		result, err := service.DeleteStock(ctx, stock.ProductId)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "cannot delete stock", err.Error())
		mockStockRepo.AssertExpectations(t)
	})
}
