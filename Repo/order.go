package Repo

import (
	"awesomeProject/entities"
	"awesomeProject/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepo struct {
	DB *gorm.DB
}

func NewOrderRepo(db *gorm.DB) OrderRepoI {
	return &OrderRepo{DB: db}
}

func (r *OrderRepo) SaveCreateOrder(order *entities.Order) error {
	orderGorm := models.OrderToGormOrder(order)
	return r.DB.Create(&orderGorm).Error
}

func (r *OrderRepo) GetOrderForUpdateStatus(id uuid.UUID) (*entities.Order, error) {
	var order models.Order
	err := r.DB.Model(&models.Order{}).Where("order_id = ?", id).First(&order).Error
	if err != nil {
		return nil, err
	}
	orderEntity := order.ToOrder()
	return orderEntity, nil
}

func (r *OrderRepo) SaveUpdateStatusOrder(o *entities.Order) error {
	orderGorm := models.OrderToGormOrder(o)
	return r.DB.Save(&orderGorm).Error
}

func (r *OrderRepo) SaveGetAllOrders() ([]*entities.Order, error) {
	var ordersGorm []*models.Order
	err := r.DB.Find(&ordersGorm).Error
	if err != nil {
		return nil, err
	}
	var orders []*entities.Order
	for _, order := range ordersGorm {
		orders = append(orders, order.ToOrder())
	}
	return orders, nil
}
