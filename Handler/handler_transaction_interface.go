package Handler

import "awesomeProject/entities"

type TransactionHandlerI interface {
	CreateTransaction(order entities.Order) error
}
