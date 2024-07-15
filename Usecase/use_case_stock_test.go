package Usecase

import (
	"awesomeProject/Repo"
	"awesomeProject/entities"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStockUseCase_CreateStock(t *testing.T) {
	mockStockRepo := new(Repo.StockRepoMock)
	service := NewStockUseCase(mockStockRepo)

	stock := &entities.Stock{ProductId: 1, Quantity: 20}
	t.Run("successful create stock", func(t *testing.T) {
		mockStockRepo.On("SaveCreateStock", stock).Return(stock, nil).Once()
		result, err := service.CreateStock(stock)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, stock, result)
		mockStockRepo.AssertExpectations(t)
	})

	t.Run("fail : cannot createStock", func(t *testing.T) {
		mockStockRepo.On("SaveCreateStock", stock).Return(nil, errors.New("cannot create stock")).Once()
		result, err := service.CreateStock(stock)
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
		mockStockRepo.On("SaveGetQtyAllProduct").Return(stocks, nil).Once()
		results, err := service.GetQtyAllProduct()
		assert.NoError(t, err)
		assert.NotNil(t, results)
		assert.Equal(t, stocks, results)
		mockStockRepo.AssertExpectations(t)
	})

	t.Run("fail : cannot GetQtyAllProduct", func(t *testing.T) {
		mockStockRepo.On("SaveGetQtyAllProduct").Return(nil, errors.New("cannot get all stock")).Once()
		results, err := service.GetQtyAllProduct()
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
		mockStockRepo.On("SaveGetQtyByIDProduct", stock.ProductId).Return(stock, nil).Once()
		result, err := service.GetQtyByIDProduct(stock.ProductId)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, stock, result)
		mockStockRepo.AssertExpectations(t)
	})

	t.Run("fail : cannot GetQtyByIDProduct", func(t *testing.T) {
		mockStockRepo.On("SaveGetQtyByIDProduct", stock.ProductId).Return(nil, errors.New("cannot get stock by id")).Once()
		result, err := service.GetQtyByIDProduct(stock.ProductId)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "cannot get stock by id", err.Error())
		mockStockRepo.AssertExpectations(t)
	})
}

func TestStockUseCase_UpdateStock(t *testing.T) {
	mockStockRepo := new(Repo.StockRepoMock)
	service := NewStockUseCase(mockStockRepo)

	stock := &entities.Stock{ProductId: 1, Quantity: 20}
	t.Run("success", func(t *testing.T) {
		mockStockRepo.On("SaveUpdateStock", stock).Return(stock, nil).Once()
		result, err := service.UpdateStock(stock, stock.ProductId)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, stock, result)
		mockStockRepo.AssertExpectations(t)
	})

	t.Run("Cannot UpdateStock", func(t *testing.T) {
		mockStockRepo.On("SaveUpdateStock", stock).Return(nil, errors.New("cannot update stock")).Once()
		result, err := service.UpdateStock(stock, stock.ProductId)
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
		mockStockRepo.On("SaveDeleteStock", stock.ProductId).Return(stock, nil).Once()
		result, err := service.DeleteStock(stock.ProductId)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, stock, result)
		mockStockRepo.AssertExpectations(t)
	})
	t.Run("Cannot DeleteStock", func(t *testing.T) {
		mockStockRepo.On("SaveDeleteStock", stock.ProductId).Return(nil, errors.New("cannot delete stock")).Once()
		result, err := service.DeleteStock(stock.ProductId)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "cannot delete stock", err.Error())
		mockStockRepo.AssertExpectations(t)
	})
}
