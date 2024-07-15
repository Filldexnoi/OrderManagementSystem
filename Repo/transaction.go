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

func (r *TransactionRepo) SaveCreateTransaction(transaction *entities.Transaction) (*entities.Transaction, error) {
	createTransaction := models.TransactionToGormTransaction(transaction)
	err := r.db.Create(&createTransaction).Error
	if err != nil {
		return nil, err
	}
	transactionEntity := createTransaction.ToTransaction()
	return transactionEntity, nil
}

func (r *TransactionRepo) SaveGetAllTransaction() ([]*entities.Transaction, error) {
	var TransactionsGorm []models.Transaction
	err := r.db.Model(&models.Transaction{}).Preload("Items.Product").Find(&TransactionsGorm).Error
	if err != nil {
		return nil, err
	}
	var transaction []*entities.Transaction
	for _, t := range TransactionsGorm {
		transaction = append(transaction, t.ToTransaction())
	}
	return transaction, nil
}

func (r *TransactionRepo) GetTransactionToCreateOrder(order *entities.Order) (*entities.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Model(&models.Transaction{}).Where("transaction_id = ?", order.TransactionId).
		Preload("Items.Product").First(&transaction).Error
	if err != nil {
		return nil, err
	}
	transactionEntity := transaction.ToTransaction()
	return transactionEntity, nil
}
