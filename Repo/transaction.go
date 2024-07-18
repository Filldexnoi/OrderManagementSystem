package Repo

import (
	"awesomeProject/entities"
	"awesomeProject/models"
	"context"
	"errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
	"reflect"
)

type TransactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) TransactionRepoI {
	return &TransactionRepo{db: db}
}

func (r *TransactionRepo) SaveCreateTransaction(transaction *entities.Transaction) (*entities.Transaction, error) {
	createTransaction := models.TransactionToGormTransaction(transaction)
	err := r.db.Create(&createTransaction).Error
	if err != nil {
		return nil, err
	}
	transactionEntity := createTransaction.ToTransaction()
	return transactionEntity, nil
}

func (r *TransactionRepo) SaveGetAllTransaction() ([]*entities.Transaction, error) {
	var TransactionsGorm []models.Transaction
	err := r.db.Model(&models.Transaction{}).Preload("Items.Product").Find(&TransactionsGorm).Error
	if err != nil {
		return nil, err
	}
	var transaction []*entities.Transaction
	for _, t := range TransactionsGorm {
		transaction = append(transaction, t.ToTransaction())
	}
	return transaction, nil
}

func (r *TransactionRepo) GetTransactionToCreateOrder(ctx context.Context, order entities.Order) (*entities.Transaction, error) {
	_, sp := otel.Tracer("order").Start(ctx, "getTransactionToCreateOrderRepository")
	defer sp.End()
	var transaction models.Transaction
	err := r.db.Model(&models.Transaction{}).Where("transaction_id = ?", order.TransactionId).
		Preload("Items.Product").First(&transaction).Error
	if err != nil {
		return nil, err
	}
	transactionEntity := transaction.ToTransaction()
	r.SetTransactionSubAttributes(transactionEntity, sp)
	return transactionEntity, nil
}

func (r *TransactionRepo) SetTransactionSubAttributes(obj any, sp trace.Span) {
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
