package Repo

import (
	"awesomeProject/entities"
	"awesomeProject/models"
	"errors"
	"gorm.io/gorm"
	"time"
)

type TransactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) TransactionRepoI {
	return &TransactionRepo{db: db}
}

func (r *TransactionRepo) SaveCreateTransaction(transaction *entities.Transaction) error {
	var product models.Product
	for i, item := range transaction.Items {
		err := r.db.Model(&models.Product{}).Select("product_price").Where("product_id = ?", item.ProductId).First(&product).Error
		if err != nil {
			return errors.New("product not found")
		}
		transaction.Items[i].Price = product.ProductPrice
	}
	transaction.CalPrice()
	transactionSave := models.Transaction{
		OrderAddress: transaction.OrderAddress,
		TotalPrice:   transaction.TotalPrice,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err := r.db.Create(&transactionSave).Error
	if err != nil {
		return errors.New("could not save transaction")
	}
	for _, item := range transaction.Items {
		dbItem := models.Item{
			TransactionId: transactionSave.TransactionId,
			Product:       models.Product{ProductId: item.ProductId},
			Quantity:      item.Quantity,
		}
		transactionSave.Items = append(transactionSave.Items, dbItem)
	}
	return r.db.Save(&transactionSave).Error
}
