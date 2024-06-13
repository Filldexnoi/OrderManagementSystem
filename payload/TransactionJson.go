package payload

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

//func (t *IncomingTransaction) ToEntityTransaction() entities.Transaction {
//	return entities.Transaction{
//		OrderAddress: t.Address,
//		Items:        t.Items,
//	}
//}
