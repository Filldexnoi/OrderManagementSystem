package Usecase

import "awesomeProject/entities"

type TransactionUseCaseI interface {
	CreateTransaction(transaction entities.Transaction) error
	GetOrder(id uint) (entities.Order, error)
}
