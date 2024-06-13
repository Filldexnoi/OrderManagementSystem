package Usecase

import (
	"awesomeProject/Repo"
	"awesomeProject/entities"
)

type TransactionUseCase struct {
	Repo Repo.TransactionRepoI
}

func NewTransactionUseCase(repo Repo.TransactionRepoI) TransactionUseCaseI {
	return &TransactionUseCase{Repo: repo}
}

func (u *TransactionUseCase) CreateTransaction(transaction *entities.Transaction) error {
	return u.Repo.SaveCreateTransaction(transaction)
}
