package Usecase

import "awesomeProject/entities"

type TransactionUseCaseI interface {
	CreateTransaction(transaction *entities.Transaction) error
}
