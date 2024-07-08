package Repo

import (
	"awesomeProject/entities"
	"awesomeProject/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Stock struct {
	db *gorm.DB
}

func NewStock(db *gorm.DB) StockRepoI {
	return &Stock{db: db}
}

func (r *Stock) SaveCreateStock(stock *entities.Stock) (*entities.Stock, error) {
	createStock := models.StockToGormStock(stock)
	err := r.db.Create(&createStock).Error
	if err != nil {
		return nil, err
	}
	stockEntity := createStock.ToStock()
	return stockEntity, nil
}

func (r *Stock) SaveUpdateStock(stock *entities.Stock) (*entities.Stock, error) {
	updateStock := models.StockToGormStock(stock)
	err := r.db.Model(&models.Stock{}).Where("product_id = ?", updateStock.ProductID).Updates(&updateStock).Error
	if err != nil {
		return nil, err
	}
	stockEntity := updateStock.ToStock()
	return stockEntity, nil
}

func (r *Stock) SaveDeleteStock(id uint) (*entities.Stock, error) {
	var deleteStock models.Stock
	if err := r.db.First(&deleteStock, id).Error; err != nil {
		return nil, err
	}

	if err := r.db.Delete(&deleteStock, id).Error; err != nil {
		return nil, err
	}
	stockEntity := deleteStock.ToStock()
	return stockEntity, nil
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
