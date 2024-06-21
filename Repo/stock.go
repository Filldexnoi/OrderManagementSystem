package Repo

import (
	"awesomeProject/entities"
	"awesomeProject/models"
	"errors"
	"fmt"
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

func (r *Stock) CheckStockToCreateOrder(transaction *entities.Transaction) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, item := range transaction.Items {
			result := tx.Model(&models.Stock{}).Where("product_id = ? AND quantity >= ?", item.ProductId, item.Quantity).
				Update("quantity", gorm.Expr("quantity - ?", item.Quantity))
			if result.Error != nil {
				return result.Error
			}

			if result.RowsAffected == 0 {
				return errors.New(fmt.Sprintf("ไม่สามารถลดจำนวนสินค้าได้เนื่องจากสินค้า ID %d มีจำนวนในสต็อกไม่เพียงพอ", item.ProductId))
			}

		}
		return nil
	})
}
