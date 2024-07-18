package Repo

import (
	"awesomeProject/entities"
	"context"
)

type TransactionRepoI interface {
	SaveCreateTransaction(transaction *entities.Transaction) (*entities.Transaction, error)
	SaveGetAllTransaction() ([]*entities.Transaction, error)
	GetTransactionToCreateOrder(ctx context.Context, order entities.Order) (*entities.Transaction, error)
}
