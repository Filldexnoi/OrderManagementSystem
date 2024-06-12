package Repo

import (
	"awesomeProject/entities"
	"gorm.io/gorm"
)

type TransactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) TransactionRepoI {
	return &TransactionRepo{db: db}
}

func (r *TransactionRepo) SaveCreateTransaction(transaction entities.Transaction) error {
	return r.db.Create(&transaction).Error
}

func (r *TransactionRepo) SaveGetOrder(id uint) (entities.Order, error) {
	order := entities.Order{}
	err := r.db.First(&order, id).Error
	return order, err
}
