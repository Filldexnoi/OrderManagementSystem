package Repo

import (
	"awesomeProject/entities"
	"gorm.io/gorm"
)

type Stock struct {
	db *gorm.DB
}

func NewStock(db *gorm.DB) StockRepoI {
	return &Stock{db: db}
}

func (r *Stock) SaveCreateStock(stock *entities.Stock) error {
	return r.db.Create(stock).Error
}

func (r *Stock) SaveUpdateStock(stock *entities.Stock) error {
	return r.db.Model(&entities.Stock{}).Where("product_id = ?", stock.ProductId).Updates(stock).Error
}

func (r *Stock) SaveDeleteStock(id uint) error {
	return r.db.Delete(&entities.Stock{}, id).Error
}

func (r *Stock) SaveGetQtyAllProduct() ([]entities.Stock, error) {
	var stocks []entities.Stock
	err := r.db.Find(&stocks).Error
	return stocks, err
}

func (r *Stock) SaveGetQtyByIDProduct(id uint) (entities.Stock, error) {
	var stock entities.Stock
	err := r.db.First(&stock, id).Error
	return stock, err
}
