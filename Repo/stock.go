package Repo

import (
	"awesomeProject/entities"
	"awesomeProject/models"
	"gorm.io/gorm"
	"time"
)

type Stock struct {
	db *gorm.DB
}

func NewStock(db *gorm.DB) StockRepoI {
	return &Stock{db: db}
}

func (r *Stock) SaveCreateStock(stock *entities.Stock) error {
	stockGorm := models.StockToGormStock(stock)
	stockGorm.CreatedAt = time.Now()
	return r.db.Create(stockGorm).Error
}

func (r *Stock) SaveUpdateStock(stock *entities.Stock) error {
	stockGorm := models.StockToGormStock(stock)
	stockGorm.UpdatedAt = time.Now()
	return r.db.Model(&models.Stock{}).Where("product_id = ?", stockGorm.ProductID).Updates(stockGorm).Error
}

func (r *Stock) SaveDeleteStock(id uint) error {
	return r.db.Delete(&models.Stock{}, id).Error
}

func (r *Stock) SaveGetQtyAllProduct() ([]*entities.Stock, error) {
	var stocksGorm []*models.Stock
	err := r.db.Find(&stocksGorm).Error
	var stocks []*entities.Stock
	for _, stock := range stocksGorm {
		stocks = append(stocks, stock.ToStock())
	}
	return stocks, err
}

func (r *Stock) SaveGetQtyByIDProduct(id uint) (*entities.Stock, error) {
	var stockGorm *models.Stock
	err := r.db.First(&stockGorm, id).Error
	stock := stockGorm.ToStock()
	return stock, err
}
