package Usecase

import (
	"awesomeProject/payload"
)

type TransactionUseCaseI interface {
	CreateTransaction(transaction *payload.RequestTransaction) error
}
