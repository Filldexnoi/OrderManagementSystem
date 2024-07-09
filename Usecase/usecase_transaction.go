package Usecase

import (
	"awesomeProject/Repo"
	"awesomeProject/entities"
	"errors"
)

type TransactionUseCase struct {
	TransactionRepo Repo.TransactionRepoI
	ProductRepo     Repo.ProductRepoI
}

func NewTransactionUseCase(TRepo Repo.TransactionRepoI, PRepo Repo.ProductRepoI) TransactionUseCaseI {
	return &TransactionUseCase{
		TransactionRepo: TRepo,
		ProductRepo:     PRepo,
	}
}

func (u *TransactionUseCase) CreateTransaction(Transaction *entities.Transaction) (*entities.Transaction, error) {
	if !Transaction.IsValidCountry(Transaction.OrderAddress) {
		return nil, errors.New("dont have this country")
	}
	seen := make(map[uint]bool)
	for _, item := range Transaction.Items {
		if seen[item.ProductId] {
			return nil, errors.New("duplicate product_id")
		}
		seen[item.ProductId] = true
	}
	transaction, err := u.ProductRepo.GetPriceProducts(Transaction)
	if err != nil {
		return nil, err
	}
	transaction.CalPrice()
	createdTransaction, err := u.TransactionRepo.SaveCreateTransaction(transaction)
	if err != nil {
		return nil, err
	}
	return createdTransaction, nil
}

func (u *TransactionUseCase) GetAllTransaction() ([]*entities.Transaction, error) {
	transactions, err := u.TransactionRepo.SaveGetAllTransaction()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
