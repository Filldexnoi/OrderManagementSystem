package Usecase

import (
	"awesomeProject/Repo"
	"awesomeProject/entities"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductUseCase_CreateProduct(t *testing.T) {
	mockProductRepo := new(Repo.ProductRepoMock)
	service := NewProductUseCase(mockProductRepo)

	product := entities.Product{ProductId: 1, ProductName: "long shirt", ProductTypes: "shirt", ProductPrice: 500}

	t.Run("successful create", func(t *testing.T) {
		mockProductRepo.On("SaveCreateProduct", product).Return(&product, nil).Once()

		result, err := service.CreateProduct(product)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, &product, result)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("fail :cannot create product", func(t *testing.T) {
		mockProductRepo.On("SaveCreateProduct", product).Return(nil, errors.New("cannot create product")).Once()

		result, err := service.CreateProduct(product)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "cannot create product", err.Error())
		mockProductRepo.AssertExpectations(t)
	})
}

func TestProductUseCase_GetAllProducts(t *testing.T) {
	mockProductRepo := new(Repo.ProductRepoMock)
	service := NewProductUseCase(mockProductRepo)
	products := []*entities.Product{
		{ProductId: 1, ProductName: "long shirt", ProductTypes: "shirt", ProductPrice: 500},
		{ProductId: 2, ProductName: "short pant", ProductTypes: "pant", ProductPrice: 1000},
	}
	t.Run("successful get all products", func(t *testing.T) {
		mockProductRepo.On("SaveGetAllProduct").Return(products, nil).Once()
		result, err := service.GetAllProducts()
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, products, result)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("Cannot get all products", func(t *testing.T) {
		mockProductRepo.On("SaveGetAllProduct").Return(nil, errors.New("cannot get all products")).Once()
		result, err := service.GetAllProducts()
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "cannot get all products", err.Error())
		mockProductRepo.AssertExpectations(t)
	})
}

func TestProductUseCase_GetByIDProduct(t *testing.T) {
	mockProductRepo := new(Repo.ProductRepoMock)
	service := NewProductUseCase(mockProductRepo)
	product := &entities.Product{ProductId: 1, ProductName: "long shirt", ProductTypes: "shirt", ProductPrice: 500}
	id := uint(1)
	t.Run("successful get product by id", func(t *testing.T) {
		mockProductRepo.On("SaveGetByIDProduct", id).Return(product, nil).Once()
		result, err := service.GetByIDProduct(1)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, product, result)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("fail: Cannot get product by id", func(t *testing.T) {
		mockProductRepo.On("SaveGetByIDProduct", id).Return(nil, errors.New("cannot get product by id")).Once()
		result, err := service.GetByIDProduct(1)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "cannot get product by id")
		mockProductRepo.AssertExpectations(t)
	})
}

func TestProductUseCase_UpdateProduct(t *testing.T) {
	mockProductRepo := new(Repo.ProductRepoMock)
	service := NewProductUseCase(mockProductRepo)
	product := entities.Product{ProductId: 1, ProductName: "zzz", ProductTypes: "shirt", ProductPrice: 1000}
	id := uint(1)
	t.Run("successful update product", func(t *testing.T) {
		mockProductRepo.On("SaveUpdateProduct", product).Return(&product, nil).Once()
		result, err := service.UpdateProduct(product, id)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, result, &product)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("fail: cannot update product", func(t *testing.T) {
		mockProductRepo.On("SaveUpdateProduct", product).Return(nil, errors.New("cannot update product")).Once()
		result, err := service.UpdateProduct(product, id)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "cannot update product", err.Error())
		mockProductRepo.AssertExpectations(t)
	})
}

func TestProductUseCase_DeleteProduct(t *testing.T) {
	mockProductRepo := new(Repo.ProductRepoMock)
	service := NewProductUseCase(mockProductRepo)
	product := &entities.Product{ProductId: 1, ProductName: "zzz", ProductTypes: "shirt", ProductPrice: 1000}
	id := uint(1)
	t.Run("successful delete product", func(t *testing.T) {
		mockProductRepo.On("SaveDeleteProduct", id).Return(product, nil).Once()
		result, err := service.DeleteProduct(id)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, id, result.ProductId)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("fail : cannot delete product", func(t *testing.T) {
		mockProductRepo.On("SaveDeleteProduct", id).Return(nil, errors.New("cannot delete product")).Once()
		result, err := service.DeleteProduct(id)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "cannot delete product", err.Error())
		mockProductRepo.AssertExpectations(t)
	})
}
