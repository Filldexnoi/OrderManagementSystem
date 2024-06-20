package Usecase

import (
	"awesomeProject/Repo"
	"awesomeProject/payload"
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

func (u *OrderUseCase) CreateOrder(o *payload.RequestCreateOrder) error {
	orderEntity := o.ToOrder()
	transaction, err := u.TransactionRepo.GetTransactionToCreateOrder(orderEntity)
	if err != nil {
		return err
	}
	err = u.StockRepo.CheckStockToCreateOrder(transaction)
	if err != nil {
		return err
	}
	order, err := orderEntity.InitStatus()
	if err != nil {
		return err
	}
	return u.OrderRepo.SaveCreateOrder(order)
}

func (u *OrderUseCase) UpdateStatusOrder(o *payload.RequestUpdateStatusOrder, id uuid.UUID) error {
	order, err := u.OrderRepo.GetOrderForUpdateStatus(id)
	if err != nil {
		return err
	}
	newStatusOrder, err := order.ChangeStatus(o.Status)
	if err != nil {
		return err
	}
	return u.OrderRepo.SaveUpdateStatusOrder(newStatusOrder)
}

func (u *OrderUseCase) GetAllOrders() ([]*payload.ResponseOrder, error) {
	orders, err := u.OrderRepo.SaveGetAllOrders()
	if err != nil {
		return nil, err
	}
	var ResOrders []*payload.ResponseOrder
	for _, order := range orders {
		ResOrders = append(ResOrders, payload.OrderToOrderRespond(order))
	}
	return ResOrders, nil
}
