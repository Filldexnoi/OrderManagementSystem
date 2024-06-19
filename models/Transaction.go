package models

import "time"

type Transaction struct {
	TransactionId uint    `gorm:"column:transaction_id;primaryKey;autoIncrement"`
	OrderAddress  string  `gorm:"column:order_address"`
	Items         []Item  `gorm:"foreignKey:transaction_id"`
	TotalPrice    float64 `gorm:"column:total_price"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Item struct {
	TransactionId uint `gorm:"column:transaction_id;primaryKey"`
	Product       uint `gorm:"column:product_id;foreignKey:product_id;primaryKey"`
	Quantity      uint `gorm:"column:quantity"`
}

func (Transaction) TableName() string {
	return "transactions"
}

func (item Item) TableName() string {
	return "items"
}
