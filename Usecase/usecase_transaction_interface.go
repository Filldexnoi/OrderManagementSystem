package Usecase

import (
	"awesomeProject/entities"
	"context"
)

type TransactionUseCaseI interface {
	CreateTransaction(ctx context.Context, transaction entities.Transaction) (*entities.Transaction, error)
	GetAllTransaction(ctx context.Context) ([]*entities.Transaction, error)
}
