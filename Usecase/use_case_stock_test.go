package Usecase

import (
	"awesomeProject/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockStockRepo struct {
	SaveCreateStockFunc         func(stock *entities.Stock) error
	SaveGetQtyAllProductFunc    func() ([]*entities.Stock, error)
	SaveGetQtyByIDProductFunc   func(id uint) (*entities.Stock, error)
	SaveUpdateStockFunc         func(stock *entities.Stock) error
	SaveDeleteStockFunc         func(id uint) error
	CheckStockToCreateOrderFunc func(transaction *entities.Transaction) error
}

func (m *mockStockRepo) SaveCreateStock(stock *entities.Stock) error {
	return m.SaveCreateStockFunc(stock)
}
func (m *mockStockRepo) SaveGetQtyAllProduct() ([]*entities.Stock, error) {
	return m.SaveGetQtyAllProductFunc()
}
func (m *mockStockRepo) SaveGetQtyByIDProduct(id uint) (*entities.Stock, error) {
	return m.SaveGetQtyByIDProductFunc(id)
}
func (m *mockStockRepo) SaveUpdateStock(stock *entities.Stock) error {
	return m.SaveUpdateStockFunc(stock)
}
func (m *mockStockRepo) SaveDeleteStock(id uint) error {
	return m.SaveDeleteStockFunc(id)
}
func (m *mockStockRepo) CheckStockToCreateOrder(transaction *entities.Transaction) error {
	return m.CheckStockToCreateOrderFunc(transaction)
}

func TestStockUseCase_CreateStock(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := &mockStockRepo{
			SaveCreateStockFunc: func(stock *entities.Stock) error {
				return nil
			},
		}
		service := NewStockUseCase(repo)

		err := service.CreateStock(
			&entities.Stock{
				ProductId: 1,
				Quantity:  5,
			})
		assert.NoError(t, err)
	})
}

func TestStockUseCase_GetQtyAllProduct(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := &mockStockRepo{
			SaveGetQtyAllProductFunc: func() ([]*entities.Stock, error) {
				return []*entities.Stock{}, nil
			},
		}
		service := NewStockUseCase(repo)
		_, err := service.GetQtyAllProduct()
		assert.NoError(t, err)
	})
}

func TestStockUseCase_GetQtyByIDProduct(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := &mockStockRepo{
			SaveGetQtyByIDProductFunc: func(id uint) (*entities.Stock, error) {
				return &entities.Stock{}, nil
			},
		}
		service := NewStockUseCase(repo)
		_, err := service.GetQtyByIDProduct(5)
		assert.NoError(t, err)
	})
}

func TestStockUseCase_UpdateStock(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := &mockStockRepo{
			SaveUpdateStockFunc: func(stock *entities.Stock) error {
				return nil
			},
		}
		service := NewStockUseCase(repo)
		err := service.UpdateStock(
			&entities.Stock{
				Quantity: 10,
			}, 5)
		assert.NoError(t, err)
	})
}

func TestStockUseCase_DeleteStock(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := &mockStockRepo{
			SaveDeleteStockFunc: func(id uint) error {
				return nil
			},
		}
		service := NewStockUseCase(repo)
		err := service.DeleteStock(5)
		assert.NoError(t, err)
	})
}
