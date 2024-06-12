package Repo

import (
	"awesomeProject/adpter"
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
	err := r.db.First(&entities.Product{}, stock.ProductId).Error
	if err != nil {
		return err
	}
	return r.db.Create(adpter.StockToCreateStockData(stock)).Error
}

func (r *Stock) SaveUpdateStock(stock *entities.Stock, id uint) error {
	return r.db.Model(&entities.Stock{}).Where("product_id = ?", id).Updates(adpter.StockToUpdateStockData(stock)).Error
}

func (r *Stock) SaveDeleteStock(id uint) error {
	return r.db.Delete(&entities.Stock{}, id).Error
}

func (r *Stock) SaveGetQtyAllProduct() ([]adpter.CRStockJson, error) {
	var stocks []adpter.CRStockJson
	err := r.db.Find(&stocks).Error
	return stocks, err
}

func (r *Stock) SaveGetQtyByIDProduct(id uint) (entities.Stock, error) {
	var stock entities.Stock
	err := r.db.First(&stock, id).Error
	return stock, err
}
