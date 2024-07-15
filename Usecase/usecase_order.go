package Usecase

import (
	"awesomeProject/Repo"
	"awesomeProject/entities"
	"github.com/google/uuid"
)

type OrderUseCase struct {
	OrderRepo       Repo.OrderRepoI
	StockRepo       Repo.StockRepoI
	TransactionRepo Repo.TransactionRepoI
}

func NewOrderUseCase(o Repo.OrderRepoI, s Repo.StockRepoI, t Repo.TransactionRepoI) OrderUseCaseI {
	return &OrderUseCase{
		OrderRepo:       o,
		StockRepo:       s,
		TransactionRepo: t,
	}
}

func (u *OrderUseCase) CreateOrder(o entities.Order) (*entities.Order, error) {
	transaction, err := u.TransactionRepo.GetTransactionToCreateOrder(o)
	if err != nil {
		return nil, err
	}
	err = u.StockRepo.CheckStockToCreateOrder(transaction)
	if err != nil {
		return nil, err
	}
	order, err := o.InitStatus()
	if err != nil {
		return nil, err
	}
	createdOrder, err := u.OrderRepo.SaveCreateOrder(order)
	if err != nil {
		return nil, err
	}
	return createdOrder, nil
}

func (u *OrderUseCase) UpdateStatusOrder(o entities.Order, id uuid.UUID) (*entities.Order, error) {
	order, err := u.OrderRepo.GetOrderForUpdateStatus(id)
	if err != nil {
		return nil, err
	}
	newStatusOrder, err := order.ChangeStatus(o.Status)
	if err != nil {
		return nil, err
	}
	updatedOrder, err := u.OrderRepo.SaveUpdateStatusOrder(newStatusOrder)
	if err != nil {
		return nil, err
	}
	return updatedOrder, nil
}

func (u *OrderUseCase) GetAllOrders() ([]*entities.Order, error) {
	orders, err := u.OrderRepo.SaveGetAllOrders()
	if err != nil {
		return nil, err
	}
	return orders, nil
}
