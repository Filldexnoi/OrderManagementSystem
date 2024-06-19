package Usecase

import (
	"awesomeProject/Repo"
	"awesomeProject/payload"
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

func (u *TransactionUseCase) CreateTransaction(ReqTransaction *payload.RequestTransaction) error {
	transactionEntity := ReqTransaction.ToTransaction()
	if !transactionEntity.IsValidCountry(transactionEntity.OrderAddress) {
		return errors.New("dont have this country")
	}
	transaction, err := u.ProductRepo.GetPriceProducts(transactionEntity)
	if err != nil {
		return err
	}
	transaction.CalPrice()
	return u.TransactionRepo.SaveCreateTransaction(transaction)
}
