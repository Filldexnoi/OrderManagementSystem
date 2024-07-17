package Usecase

import (
	"awesomeProject/Repo"
	"awesomeProject/entities"
	"context"
	"errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"reflect"
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

func (u *TransactionUseCase) CreateTransaction(ctx context.Context, Transaction entities.Transaction) (*entities.Transaction, error) {
	ctx, sp := otel.Tracer("transaction").Start(ctx, "transactionCreateUseCase")
	defer sp.End()

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
	transaction, err := u.ProductRepo.GetPriceProducts(ctx, &Transaction)
	if err != nil {
		return nil, err
	}
	transaction.CalPrice()
	createdTransaction, err := u.TransactionRepo.SaveCreateTransaction(ctx, transaction)
	if err != nil {
		return nil, err
	}
	u.SetTransactionSubAttributes(createdTransaction, sp)
	return createdTransaction, nil
}

func (u *TransactionUseCase) GetAllTransaction(ctx context.Context) ([]*entities.Transaction, error) {
	ctx, sp := otel.Tracer("transaction").Start(ctx, "transactionGetAllUseCase")
	defer sp.End()
	transactions, err := u.TransactionRepo.SaveGetAllTransaction(ctx)
	if err != nil {
		return nil, err
	}
	u.SetTransactionSubAttributes(transactions, sp)
	return transactions, nil
}

func (u *TransactionUseCase) SetTransactionSubAttributes(obj any, sp trace.Span) {
	if transactions, ok := obj.([]*entities.Transaction); ok {
		transactionID := make([]string, len(transactions))
		transactionOrderAddress := make([]string, len(transactions))
		transactionTotalPrice := make([]float64, len(transactions))

		for _, transaction := range transactions {
			transactionID = append(transactionID, transaction.TransactionId.String())
			transactionOrderAddress = append(transactionOrderAddress, transaction.OrderAddress)
			transactionTotalPrice = append(transactionTotalPrice, transaction.TotalPrice)
		}
		sp.SetAttributes(
			attribute.StringSlice("TransactionID", transactionID),
			attribute.StringSlice("TransactionOrderAddress", transactionOrderAddress),
			attribute.Float64Slice("TransactionTotalPrice", transactionTotalPrice),
		)
	} else if transaction, ok := obj.(*entities.Transaction); ok {
		sp.SetAttributes(
			attribute.String("TransactionID", transaction.TransactionId.String()),
			attribute.String("TransactionOrderAddress", transaction.OrderAddress),
			attribute.Float64("TransactionTotalPrice", transaction.TotalPrice),
		)
	} else {
		sp.RecordError(errors.New("invalid type" + reflect.TypeOf(obj).String()))
	}
}
