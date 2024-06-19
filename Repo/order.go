package Repo

import (
	"awesomeProject/entities"
	"awesomeProject/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type OrderRepo struct {
	DB *gorm.DB
}

func NewOrderRepo(db *gorm.DB) OrderRepoI {
	return &OrderRepo{
		DB: db,
	}
}

func (r *OrderRepo) SaveCreateOrder(order entities.Order) error {
	var transaction models.Transaction
	err := r.DB.Model(&models.Transaction{}).Where("transaction_id = ?", order.TransactionId).First(&transaction).Error
	if err != nil {
		return errors.New("record not found")
	}
	for _, item := range transaction.Items {
		result := r.DB.Model(&models.Stock{}).Where("product_id = ? AND quantity >= ?", item.Product, item.Quantity).
			Update("quantity", gorm.Expr("quantity - ?", item.Quantity))
		if result.Error != nil {
			return result.Error
		}
		fmt.Printf("อัปเดตสต็อกสำเร็จ: ProductID = %d, QuantityToRemove = %d, RowsAffected = %d\n", item.Product, item.Quantity, result.RowsAffected)
	}
	return r.DB.Create(&order).Error
}
