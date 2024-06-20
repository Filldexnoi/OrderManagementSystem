package Repo

import (
	"awesomeProject/entities"
	"awesomeProject/models"
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
