package Usecase

import (
	"awesomeProject/Repo"
	"awesomeProject/payload"
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

func (u *OrderUseCase) CreateOrder(order *payload.RequestOrder) error {
	orderEntity := order.ToOrder()
	transaction, err := u.TransactionRepo.GetTransactionToCreateOrder(orderEntity)
	if err != nil {
		return err
	}
	err = u.StockRepo.CheckStockToCreateOrder(transaction)
	if err != nil {
		return err
	}
	return u.OrderRepo.SaveCreateOrder(orderEntity)
}
