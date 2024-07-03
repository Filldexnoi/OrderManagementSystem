package Usecase

import (
	"awesomeProject/entities"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockOrderRepo struct {
	SaveCreateOrderFunc         func(order *entities.Order) error
	GetOrderForUpdateStatusFunc func(id uuid.UUID) (*entities.Order, error)
	SaveUpdateStatusOrderFunc   func(o *entities.Order) error
	SaveGetAllOrdersFunc        func() ([]*entities.Order, error)
}

func (m *mockOrderRepo) SaveCreateOrder(order *entities.Order) error {
	return m.SaveCreateOrderFunc(order)
}

func (m *mockOrderRepo) GetOrderForUpdateStatus(id uuid.UUID) (*entities.Order, error) {
	return m.GetOrderForUpdateStatusFunc(id)
}

func (m *mockOrderRepo) SaveUpdateStatusOrder(order *entities.Order) error {
	return m.SaveUpdateStatusOrderFunc(order)
}

func (m *mockOrderRepo) SaveGetAllOrders() ([]*entities.Order, error) {
	return m.SaveGetAllOrdersFunc()
}

func TestOrderUseCase_CreateOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ORepo := &mockOrderRepo{
			SaveCreateOrderFunc: func(order *entities.Order) error {
				return nil
			},
		}
		TRepo := &mockTransactionRepo{
			GetTransactionToCreateOrderFunc: func(order *entities.Order) (*entities.Transaction, error) {
				return &entities.Transaction{}, nil
			},
		}
		SRepo := &mockStockRepo{
			CheckStockToCreateOrderFunc: func(transaction *entities.Transaction) error {
				return nil
			},
		}
		service := NewOrderUseCase(ORepo, SRepo, TRepo)
		err := service.CreateOrder(
			&entities.Order{
				TransactionId: uuid.Nil,
			},
		)
		assert.NoError(t, err)
	})

	t.Run("Stock not enough", func(t *testing.T) {
		ORepo := &mockOrderRepo{
			SaveCreateOrderFunc: func(order *entities.Order) error {
				return nil
			},
		}
		TRepo := &mockTransactionRepo{
			GetTransactionToCreateOrderFunc: func(order *entities.Order) (*entities.Transaction, error) {
				return &entities.Transaction{}, nil
			},
		}
		SRepo := &mockStockRepo{
			CheckStockToCreateOrderFunc: func(transaction *entities.Transaction) error {
				return errors.New("stock not enough")
			},
		}
		service := NewOrderUseCase(ORepo, SRepo, TRepo)
		err := service.CreateOrder(
			&entities.Order{
				TransactionId: uuid.Nil,
			},
		)
		assert.Error(t, err)
		assert.EqualError(t, err, "stock not enough")
	})

	t.Run("Dont have Transaction", func(t *testing.T) {
		ORepo := &mockOrderRepo{
			SaveCreateOrderFunc: func(order *entities.Order) error {
				return nil
			},
		}
		TRepo := &mockTransactionRepo{
			GetTransactionToCreateOrderFunc: func(order *entities.Order) (*entities.Transaction, error) {
				return nil, errors.New("dont have transaction")
			},
		}
		SRepo := &mockStockRepo{
			CheckStockToCreateOrderFunc: func(transaction *entities.Transaction) error {
				return nil
			},
		}
		service := NewOrderUseCase(ORepo, SRepo, TRepo)
		err := service.CreateOrder(
			&entities.Order{
				TransactionId: uuid.Nil,
			},
		)
		assert.Error(t, err)
		assert.EqualError(t, err, "dont have transaction")
	})

	t.Run("Invalid status to init", func(t *testing.T) {
		ORepo := &mockOrderRepo{
			SaveCreateOrderFunc: func(order *entities.Order) error {
				return nil
			},
		}
		TRepo := &mockTransactionRepo{
			GetTransactionToCreateOrderFunc: func(order *entities.Order) (*entities.Transaction, error) {
				return &entities.Transaction{}, nil
			},
		}
		SRepo := &mockStockRepo{
			CheckStockToCreateOrderFunc: func(transaction *entities.Transaction) error {
				return nil
			},
		}
		service := NewOrderUseCase(ORepo, SRepo, TRepo)
		err := service.CreateOrder(
			&entities.Order{
				TransactionId: uuid.Nil,
				Status:        "Processing",
			})
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid order status: from Processing to New")
	})

	t.Run("Cannot save create order", func(t *testing.T) {
		ORepo := &mockOrderRepo{
			SaveCreateOrderFunc: func(order *entities.Order) error {
				return errors.New("cannot save create order")
			},
		}
		TRepo := &mockTransactionRepo{
			GetTransactionToCreateOrderFunc: func(order *entities.Order) (*entities.Transaction, error) {
				return &entities.Transaction{}, nil
			},
		}
		SRepo := &mockStockRepo{
			CheckStockToCreateOrderFunc: func(transaction *entities.Transaction) error {
				return nil
			},
		}
		service := NewOrderUseCase(ORepo, SRepo, TRepo)
		err := service.CreateOrder(
			&entities.Order{
				TransactionId: uuid.Nil,
			})
		assert.Error(t, err)
		assert.EqualError(t, err, "cannot save create order")
	})
}

func TestOrderUseCase_UpdateStatus(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ORepo := &mockOrderRepo{
			GetOrderForUpdateStatusFunc: func(id uuid.UUID) (*entities.Order, error) {
				return &entities.Order{
					Status: "New",
				}, nil
			},
			SaveUpdateStatusOrderFunc: func(o *entities.Order) error {
				return nil
			},
		}
		TRepo := &mockTransactionRepo{}
		SRepo := &mockStockRepo{}

		service := NewOrderUseCase(ORepo, SRepo, TRepo)
		order := &entities.Order{
			Status: "Paid",
		}
		err := service.UpdateStatusOrder(order, uuid.Nil)
		assert.NoError(t, err)
	})

	t.Run("Cannot save update status", func(t *testing.T) {
		ORepo := &mockOrderRepo{
			GetOrderForUpdateStatusFunc: func(id uuid.UUID) (*entities.Order, error) {
				return &entities.Order{Status: "New"}, nil
			},
			SaveUpdateStatusOrderFunc: func(o *entities.Order) error {
				return errors.New("cannot save update status")
			},
		}
		TRepo := &mockTransactionRepo{}
		SRepo := &mockStockRepo{}
		service := NewOrderUseCase(ORepo, SRepo, TRepo)
		order := &entities.Order{
			Status: "Paid",
		}
		err := service.UpdateStatusOrder(order, uuid.Nil)
		assert.Error(t, err)
		assert.EqualError(t, err, "cannot save update status")
	})

	t.Run("Cannot get order", func(t *testing.T) {
		ORepo := &mockOrderRepo{
			GetOrderForUpdateStatusFunc: func(id uuid.UUID) (*entities.Order, error) {
				return nil, errors.New("cannot get order")
			},
			SaveUpdateStatusOrderFunc: func(o *entities.Order) error {
				return nil
			},
		}
		TRepo := &mockTransactionRepo{}
		SRepo := &mockStockRepo{}
		service := NewOrderUseCase(ORepo, SRepo, TRepo)
		order := &entities.Order{
			Status: "Paid",
		}
		err := service.UpdateStatusOrder(order, uuid.Nil)
		assert.Error(t, err)
		assert.EqualError(t, err, "cannot get order")
	})

	t.Run("Cannot change status", func(t *testing.T) {
		ORepo := &mockOrderRepo{
			GetOrderForUpdateStatusFunc: func(id uuid.UUID) (*entities.Order, error) {
				return &entities.Order{Status: "New"}, nil
			},
			SaveUpdateStatusOrderFunc: func(o *entities.Order) error {
				return nil
			},
		}
		TRepo := &mockTransactionRepo{}
		SRepo := &mockStockRepo{}
		service := NewOrderUseCase(ORepo, SRepo, TRepo)
		order := &entities.Order{
			Status: "Processing",
		}
		err := service.UpdateStatusOrder(order, uuid.Nil)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid o status: from New to Processing")
	})
}

func TestOrderUseCase_GetAllOrders(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ORepo := &mockOrderRepo{
			SaveGetAllOrdersFunc: func() ([]*entities.Order, error) {
				return []*entities.Order{}, nil
			},
		}
		TRepo := &mockTransactionRepo{}
		SRepo := &mockStockRepo{}
		service := NewOrderUseCase(ORepo, SRepo, TRepo)
		_, err := service.GetAllOrders()
		assert.NoError(t, err)
	})

	t.Run("Cannot get all orders", func(t *testing.T) {
		ORepo := &mockOrderRepo{
			SaveGetAllOrdersFunc: func() ([]*entities.Order, error) {
				return nil, errors.New("cannot get all orders")
			},
		}
		TRepo := &mockTransactionRepo{}
		SRepo := &mockStockRepo{}
		service := NewOrderUseCase(ORepo, SRepo, TRepo)
		_, err := service.GetAllOrders()
		assert.Error(t, err)
		assert.EqualError(t, err, "cannot get all orders")
	})
}
