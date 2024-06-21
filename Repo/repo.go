package Repo

import "gorm.io/gorm"

type GormRepo struct {
	ProductRepo     ProductRepoI
	StockRepo       StockRepoI
	TransactionRepo TransactionRepoI
	OrderRepo       OrderRepoI
}

func NewGormRepo(db *gorm.DB) *GormRepo {
	return &GormRepo{
		ProductRepo:     NewProductRepo(db),
		StockRepo:       NewStock(db),
		TransactionRepo: NewTransactionRepo(db),
		OrderRepo:       NewOrderRepo(db),
	}
}
