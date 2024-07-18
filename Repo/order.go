package Repo

import (
	"awesomeProject/entities"
	"awesomeProject/models"
	"context"
	"errors"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
	"reflect"
)

type OrderRepo struct {
	DB *gorm.DB
}

func NewOrderRepo(db *gorm.DB) OrderRepoI {
	return &OrderRepo{DB: db}
}

func (r *OrderRepo) SaveCreateOrder(ctx context.Context, order entities.Order) (*entities.Order, error) {
	_, sp := otel.Tracer("order").Start(ctx, "orderCreateRepository")
	defer sp.End()
	createdOrder := models.OrderToGormOrder(order)
	err := r.DB.Create(&createdOrder).Error
	if err != nil {
		return nil, err
	}
	orderEntity := createdOrder.ToOrder()
	r.SetOrderSubAttributes(order, sp)
	return orderEntity, nil
}

func (r *OrderRepo) GetOrderForUpdateStatus(ctx context.Context, id uuid.UUID) (*entities.Order, error) {
	_, sp := otel.Tracer("order").Start(ctx, "GetOrderForUpdateStatusRepository")
	defer sp.End()
	var order models.Order
	err := r.DB.Model(&models.Order{}).Where("order_id = ?", id).First(&order).Error
	if err != nil {
		return nil, err
	}
	orderEntity := order.ToOrder()
	r.SetOrderSubAttributes(order, sp)
	return orderEntity, nil
}

func (r *OrderRepo) SaveUpdateStatusOrder(ctx context.Context, o entities.Order) (*entities.Order, error) {
	_, sp := otel.Tracer("order").Start(ctx, "SaveUpdateStatusOrderRepository")
	defer sp.End()
	updatedOrder := models.OrderToGormOrder(o)
	err := r.DB.Save(&updatedOrder).Error
	if err != nil {
		return nil, err
	}
	orderEntity := updatedOrder.ToOrder()
	r.SetOrderSubAttributes(o, sp)
	return orderEntity, nil
}

func (r *OrderRepo) SaveGetAllOrders(ctx context.Context) ([]*entities.Order, error) {
	_, sp := otel.Tracer("order").Start(ctx, "SaveGetAllOrdersRepository")
	var ordersGorm []*models.Order
	err := r.DB.Find(&ordersGorm).Error
	if err != nil {
		return nil, err
	}
	var orders []*entities.Order
	for _, order := range ordersGorm {
		orders = append(orders, order.ToOrder())
	}
	r.SetOrderSubAttributes(orders, sp)
	return orders, nil
}

func (r *OrderRepo) SetOrderSubAttributes(obj any, sp trace.Span) {
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
