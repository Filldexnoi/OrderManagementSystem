package Usecase

import (
	"awesomeProject/entities"
)

type TransactionUseCaseI interface {
	CreateTransaction(transaction entities.Transaction) (*entities.Transaction, error)
	GetAllTransaction() ([]*entities.Transaction, error)
}
