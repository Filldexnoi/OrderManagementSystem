package Repo

import (
	"awesomeProject/entities"
	"awesomeProject/models"
	"gorm.io/gorm"
)

type TransactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) TransactionRepoI {
	return &TransactionRepo{db: db}
}

func (r *TransactionRepo) SaveCreateTransaction(transaction *entities.Transaction) error {
	TransactionGorm := models.TransactionToGormTransaction(transaction)
	err := r.db.Create(&TransactionGorm).Error
	return err
}
