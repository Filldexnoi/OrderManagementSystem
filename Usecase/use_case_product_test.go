package Usecase

import (
	"awesomeProject/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockProductRepo struct {
	SaveCreateProductFunc  func(product *entities.Product) error
	SaveGetAllProductFunc  func() ([]*entities.Product, error)
	SaveGetByIDProductFunc func(id uint) (*entities.Product, error)
	SaveUpdateProductFunc  func(product *entities.Product) error
	SaveDeleteProductFunc  func(id uint) error
	GetPriceProductsFunc   func(transaction *entities.Transaction) (*entities.Transaction, error)
}

func (m *mockProductRepo) SaveCreateProduct(product *entities.Product) error {
	return m.SaveCreateProductFunc(product)
}
func (m *mockProductRepo) SaveGetAllProduct() ([]*entities.Product, error) {
	return m.SaveGetAllProductFunc()
}
func (m *mockProductRepo) SaveGetByIDProduct(id uint) (*entities.Product, error) {
	return m.SaveGetByIDProductFunc(id)
}
func (m *mockProductRepo) SaveUpdateProduct(product *entities.Product) error {
	return m.SaveUpdateProductFunc(product)
}
func (m *mockProductRepo) SaveDeleteProduct(id uint) error {
	return m.SaveDeleteProductFunc(id)
}
func (m *mockProductRepo) GetPriceProducts(transaction *entities.Transaction) (*entities.Transaction, error) {
	return m.GetPriceProductsFunc(transaction)
}

func TestProductUseCase_CreateProduct(t *testing.T) {
	// Success case
	t.Run("success", func(t *testing.T) {
		repo := &mockProductRepo{
			SaveCreateProductFunc: func(product *entities.Product) error {
				// Simulate successful save
				return nil
			},
		}
		service := NewProductUseCase(repo)

		err := service.CreateProduct(
			&entities.Product{
				ProductTypes: "Shirt",
				ProductName:  "Long Shirt",
				ProductPrice: 999,
			})
		assert.NoError(t, err)
	})
}

func TestProductUseCase_GetAllProducts(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := &mockProductRepo{
			SaveGetAllProductFunc: func() ([]*entities.Product, error) {
				return []*entities.Product{}, nil
			},
		}
		service := NewProductUseCase(repo)

		_, err := service.GetAllProducts()
		assert.NoError(t, err)
	})
}

func TestProductUseCase_GetByIDProduct(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := &mockProductRepo{
			SaveGetByIDProductFunc: func(id uint) (*entities.Product, error) {
				return &entities.Product{}, nil
			},
		}
		service := NewProductUseCase(repo)

		_, err := service.GetByIDProduct(5)
		assert.NoError(t, err)
	})
}

func TestProductUseCase_UpdateProduct(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := &mockProductRepo{
			SaveUpdateProductFunc: func(product *entities.Product) error {
				return nil
			},
		}
		service := NewProductUseCase(repo)

		err := service.UpdateProduct(
			&entities.Product{
				ProductTypes: "Shirt",
				ProductName:  "Long Shirt",
				ProductPrice: 999,
			}, 5)
		assert.NoError(t, err)
	})

	t.Run("Price Product is negative", func(t *testing.T) {
		repo := &mockProductRepo{
			SaveUpdateProductFunc: func(product *entities.Product) error {
				return nil
			},
		}
		service := NewProductUseCase(repo)

		err := service.CreateProduct(
			&entities.Product{
				ProductTypes: "Shirt",
				ProductName:  "Long Shirt",
				ProductPrice: -999,
			})
		assert.Error(t, err)
		assert.Equal(t, "product price must be not negative", err.Error())
	})
}

func TestProductUseCase_DeleteProduct(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := &mockProductRepo{
			SaveDeleteProductFunc: func(id uint) error {
				return nil
			},
		}
		service := NewProductUseCase(repo)
		err := service.DeleteProduct(5)
		assert.NoError(t, err)
	})
}
