package Repo

import "awesomeProject/entities"

type TransactionRepoI interface {
	SaveCreateTransaction(transaction *entities.Transaction) error
	SaveGetAllTransaction() ([]*entities.Transaction, error)
}
