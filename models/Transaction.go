package models

type Transaction struct {
	TransactionId uint    `gorm:"column:transaction_id;primaryKey;AUTO_INCREMENT"`
	TotalPrice    float64 `gorm:"column:total_price;type:decimal(20,8);not null"`
}
