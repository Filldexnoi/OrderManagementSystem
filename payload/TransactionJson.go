package payload

import "awesomeProject/entities"

type SaveTransactionData struct {
	Address    string
	Items      []Item
	TotalPrice float64
}

type Item struct {
	ProductId uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

type IncomingTransaction struct {
	Address string `json:"address"`
	Items   []Item `json:"items"`
}

type OutgoingTransaction struct {
	TransactionID uint    `json:"transaction_id"`
	Address       string  `json:"address"`
	Items         []Item  `json:"items"`
	TotalPrice    float64 `json:"total_price"`
}

func (t *SaveTransactionData) TableName() string {
	return "transactions"
}
func (t *IncomingTransaction) ToEntityTransaction() *entities.Transaction {
	items := make([]entities.Item, len(t.Items))
	for i, item := range t.Items {
		items[i] = entities.Item{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		}
	}
	return &entities.Transaction{
		OrderAddress: t.Address,
		Items:        items,
	}
}
