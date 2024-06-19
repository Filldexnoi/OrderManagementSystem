package models

import (
	"awesomeProject/entities"
	"time"
)

type Transaction struct {
	TransactionID uint   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	OrderAddress  string `gorm:"column:order_address"`
	Items         []Item
	TotalPrice    float64   `gorm:"column:total_price"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}

type Item struct {
	TransactionID uint `gorm:"column:transaction_id;primaryKey"`
	ProductID     uint `gorm:"column:product_id;primaryKey"`
	Quantity      uint `gorm:"column:quantity"`
	Product       *Product
}

func (Transaction) TableName() string {
	return "transactions"
}

func (Item) TableName() string {
	return "items"
}

func (t Transaction) ToTransaction() entities.Transaction {
	return entities.Transaction{
		TransactionId: t.TransactionID,
		OrderAddress:  t.OrderAddress,
		Items:         t.ToItems(),
		TotalPrice:    t.TotalPrice,
	}
}

func (i Item) ToItem() entities.Item {
	return entities.Item{
		ProductId: i.Product.ProductId,
		Quantity:  i.Quantity,
		Price:     i.Product.ProductPrice,
	}
}

func (t Transaction) ToItems() []entities.Item {
	var items []entities.Item
	for _, item := range t.Items {
		items = append(items, item.ToItem())
	}
	return items
}

func TransactionToGormTransaction(transaction entities.Transaction) Transaction {
	return Transaction{
		TransactionID: transaction.TransactionId,
		OrderAddress:  transaction.OrderAddress,
		Items:         ItemToGormItem(transaction),
		TotalPrice:    transaction.TotalPrice,
	}
}

func ItemToGormItem(transaction entities.Transaction) []Item {
	var items []Item
	for _, item := range transaction.Items {
		items = append(items, Item{
			TransactionID: transaction.TransactionId,
			ProductID:     item.ProductId,
			Quantity:      item.Quantity,
		})
	}
	return items
}
