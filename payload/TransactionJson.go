package payload

import (
	"awesomeProject/entities"
	"github.com/google/uuid"
)

type Item struct {
	ProductId uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

type RequestTransaction struct {
	Address string `json:"address"`
	Items   []Item `json:"items"`
}

type RespondTransaction struct {
	TransactionID uuid.UUID `json:"transaction_id"`
	Address       string    `json:"address"`
	Items         []Item    `json:"items"`
	TotalPrice    float64   `json:"total_price"`
}

func (t *RequestTransaction) ToTransaction() entities.Transaction {
	items := make([]entities.Item, len(t.Items))
	for i, item := range t.Items {
		items[i] = entities.Item{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		}
	}
	return entities.Transaction{
		OrderAddress: t.Address,
		Items:        items,
	}
}

func TransactionToResTransaction(transaction *entities.Transaction) *RespondTransaction {
	items := make([]Item, len(transaction.Items))
	for i, item := range transaction.Items {
		items[i] = Item{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		}
	}
	return &RespondTransaction{
		TransactionID: transaction.TransactionId,
		Address:       transaction.OrderAddress,
		Items:         items,
		TotalPrice:    transaction.TotalPrice,
	}
}
