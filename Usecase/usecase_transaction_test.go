package Usecase

import (
	"awesomeProject/entities"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockTransactionRepo struct {
	SaveCreateTransactionFunc       func(transaction *entities.Transaction) error
	SaveGetAllTransactionFunc       func() ([]*entities.Transaction, error)
	GetTransactionToCreateOrderFunc func(order *entities.Order) (*entities.Transaction, error)
}

func (m *mockTransactionRepo) SaveCreateTransaction(transaction *entities.Transaction) error {
	return m.SaveCreateTransactionFunc(transaction)
}
func (m *mockTransactionRepo) SaveGetAllTransaction() ([]*entities.Transaction, error) {
	return m.SaveGetAllTransactionFunc()
}
func (m *mockTransactionRepo) GetTransactionToCreateOrder(order *entities.Order) (*entities.Transaction, error) {
	return m.GetTransactionToCreateOrderFunc(order)
}

func TestTransactionUseCase_CreateTransaction(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		TRepo := &mockTransactionRepo{
			SaveCreateTransactionFunc: func(transaction *entities.Transaction) error {
				return nil
			},
		}
		PRepo := &mockProductRepo{
			GetPriceProductsFunc: func(transaction *entities.Transaction) (*entities.Transaction, error) {
				return &entities.Transaction{}, nil
			},
		}
		service := NewTransactionUseCase(TRepo, PRepo)

		err := service.CreateTransaction(
			&entities.Transaction{
				OrderAddress: "th",
				Items:        []entities.Item{{ProductId: 1, Quantity: 5}, {ProductId: 2, Quantity: 3}},
			},
		)
		assert.NoError(t, err)
	})

	t.Run("Invalid country", func(t *testing.T) {
		TRepo := &mockTransactionRepo{
			SaveCreateTransactionFunc: func(transaction *entities.Transaction) error {
				return nil
			},
		}
		PRepo := &mockProductRepo{
			GetPriceProductsFunc: func(transaction *entities.Transaction) (*entities.Transaction, error) {
				return &entities.Transaction{}, nil
			},
		}
		service := NewTransactionUseCase(TRepo, PRepo)
		err := service.CreateTransaction(
			&entities.Transaction{
				OrderAddress: "xx",
				Items:        []entities.Item{{ProductId: 1, Quantity: 5}, {ProductId: 2, Quantity: 3}},
			})
		assert.Error(t, err)
		assert.Equal(t, "dont have this country", err.Error())
	})

	t.Run("Duplicate Id Product", func(t *testing.T) {
		TRepo := &mockTransactionRepo{
			SaveCreateTransactionFunc: func(transaction *entities.Transaction) error {
				return nil
			},
		}
		PRepo := &mockProductRepo{
			GetPriceProductsFunc: func(transaction *entities.Transaction) (*entities.Transaction, error) {
				return &entities.Transaction{}, nil
			},
		}
		service := NewTransactionUseCase(TRepo, PRepo)
		err := service.CreateTransaction(
			&entities.Transaction{
				OrderAddress: "th",
				Items:        []entities.Item{{ProductId: 1, Quantity: 5}, {ProductId: 1, Quantity: 3}},
			})
		assert.Error(t, err)
		assert.Equal(t, "duplicate product_id", err.Error())
	})

	t.Run("Cannot get PriceProducts", func(t *testing.T) {
		TRepo := &mockTransactionRepo{
			SaveCreateTransactionFunc: func(transaction *entities.Transaction) error {
				return nil
			},
		}
		PRepo := &mockProductRepo{
			GetPriceProductsFunc: func(transaction *entities.Transaction) (*entities.Transaction, error) {
				return nil, errors.New("cannot get price products")
			},
		}
		service := NewTransactionUseCase(TRepo, PRepo)
		err := service.CreateTransaction(&entities.Transaction{
			OrderAddress: "th",
			Items:        []entities.Item{{ProductId: 1, Quantity: 5}, {ProductId: 2, Quantity: 3}},
		})
		assert.Error(t, err)
		assert.Equal(t, "cannot get price products", err.Error())
	})

	t.Run("Cannot create transaction", func(t *testing.T) {
		TRepo := &mockTransactionRepo{
			SaveCreateTransactionFunc: func(transaction *entities.Transaction) error {
				return errors.New("cannot create transaction")
			},
		}
		PRepo := &mockProductRepo{
			GetPriceProductsFunc: func(transaction *entities.Transaction) (*entities.Transaction, error) {
				return &entities.Transaction{}, nil
			},
		}
		service := NewTransactionUseCase(TRepo, PRepo)
		err := service.CreateTransaction(&entities.Transaction{
			OrderAddress: "th",
			Items:        []entities.Item{{ProductId: 1, Quantity: 5}, {ProductId: 2, Quantity: 3}},
		})
		assert.Error(t, err)
		assert.Equal(t, "cannot create transaction", err.Error())
	})
}

func TestTransactionUseCase_GetAllTransaction(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		TRepo := &mockTransactionRepo{
			SaveGetAllTransactionFunc: func() ([]*entities.Transaction, error) {
				return []*entities.Transaction{}, nil
			},
		}
		PRepo := &mockProductRepo{}
		service := NewTransactionUseCase(TRepo, PRepo)
		_, err := service.GetAllTransaction()
		assert.NoError(t, err)
	})

	t.Run("Cannot get transactions", func(t *testing.T) {
		TRepo := &mockTransactionRepo{
			SaveGetAllTransactionFunc: func() ([]*entities.Transaction, error) {
				return nil, errors.New("cannot get transactions")
			},
		}
		PRepo := &mockProductRepo{}
		service := NewTransactionUseCase(TRepo, PRepo)
		_, err := service.GetAllTransaction()
		assert.Error(t, err)
		assert.Equal(t, "cannot get transactions", err.Error())
	})
}