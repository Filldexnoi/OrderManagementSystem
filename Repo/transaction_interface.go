package Repo

import (
	"awesomeProject/entities"
	"context"
)

type TransactionRepoI interface {
	SaveCreateTransaction(context context.Context, transaction *entities.Transaction) (*entities.Transaction, error)
	SaveGetAllTransaction(context context.Context) ([]*entities.Transaction, error)
	GetTransactionToCreateOrder(ctx context.Context, order entities.Order) (*entities.Transaction, error)
}
