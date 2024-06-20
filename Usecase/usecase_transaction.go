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
	seen := make(map[uint]bool)
	for _, item := range transactionEntity.Items {
		if seen[item.ProductId] {
			return errors.New("duplicate product_id found")
		}
		seen[item.ProductId] = true
	}
	transaction, err := u.ProductRepo.GetPriceProducts(transactionEntity)
	if err != nil {
		return err
	}
	transaction.CalPrice()
	return u.TransactionRepo.SaveCreateTransaction(transaction)
}

func (u *TransactionUseCase) GetAllTransaction() ([]*payload.RespondTransaction, error) {
	transactions, err := u.TransactionRepo.SaveGetAllTransaction()
	if err != nil {
		return nil, err
	}
	var ResTransaction []*payload.RespondTransaction
	for _, transaction := range transactions {
		ResTransaction = append(ResTransaction, payload.TransactionToResTransaction(transaction))
	}
	return ResTransaction, nil
}
