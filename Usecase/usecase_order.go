package Usecase

import (
	"awesomeProject/Repo"
	"awesomeProject/entities"
	"context"
	"errors"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"reflect"
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

func (u *OrderUseCase) CreateOrder(ctx context.Context, o entities.Order) (*entities.Order, error) {
	ctx, sp := otel.Tracer("order").Start(ctx, "orderCreateUseCase")
	defer sp.End()
	transaction, err := u.TransactionRepo.GetTransactionToCreateOrder(ctx, o)
	if err != nil {
		return nil, err
	}
	err = u.StockRepo.CheckStockToCreateOrder(ctx, transaction)
	if err != nil {
		return nil, err
	}
	order, err := o.InitStatus()
	if err != nil {
		return nil, err
	}
	createdOrder, err := u.OrderRepo.SaveCreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}
	u.SetOrderSubAttributes(createdOrder, sp)
	return createdOrder, nil
}

func (u *OrderUseCase) UpdateStatusOrder(ctx context.Context, o entities.Order, id uuid.UUID) (*entities.Order, error) {
	ctx, sp := otel.Tracer("order").Start(ctx, "orderUpdateUseCase")
	defer sp.End()
	order, err := u.OrderRepo.GetOrderForUpdateStatus(ctx, id)
	if err != nil {
		return nil, err
	}
	newStatusOrder, err := order.ChangeStatus(o.Status)
	if err != nil {
		return nil, err
	}
	updatedOrder, err := u.OrderRepo.SaveUpdateStatusOrder(ctx, newStatusOrder)
	if err != nil {
		return nil, err
	}
	u.SetOrderSubAttributes(updatedOrder, sp)
	return updatedOrder, nil
}

func (u *OrderUseCase) GetAllOrders(ctx context.Context) ([]*entities.Order, error) {
	ctx, sp := otel.Tracer("order").Start(ctx, "orderGetAllUseCase")
	defer sp.End()
	orders, err := u.OrderRepo.SaveGetAllOrders(ctx)
	if err != nil {
		return nil, err
	}
	u.SetOrderSubAttributes(orders, sp)
	return orders, nil
}

func (u *OrderUseCase) SetOrderSubAttributes(obj any, sp trace.Span) {
	if orders, ok := obj.([]*entities.Order); ok {
		var orderID []string
		var transactionID []string
		var status []string

		for _, order := range orders {
			orderID = append(orderID, order.OrderId.String())
			transactionID = append(transactionID, order.TransactionId.String())
			status = append(status, order.Status)
		}
		sp.SetAttributes(
			attribute.StringSlice("orderID", orderID),
			attribute.StringSlice("transactionID", transactionID),
			attribute.StringSlice("status", status),
		)
	} else if order, ok := obj.(*entities.Order); ok {
		sp.SetAttributes(
			attribute.String("orderID", order.OrderId.String()),
			attribute.String("transactionID", order.TransactionId.String()),
			attribute.String("status", order.Status),
		)
	} else {
		sp.RecordError(errors.New("invalid type" + reflect.TypeOf(obj).String()))
	}
}
