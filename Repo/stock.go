package Repo

import (
	"awesomeProject/entities"
	"awesomeProject/payload"
	"errors"
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

	return r.db.Create(payload.IncomingStock(stock)).Error
}

func (r *Stock) SaveUpdateStock(stock *entities.Stock, id uint) error {
	err := r.db.First(&entities.Product{}, id).Error
	if err != nil {
		return errors.New("dont have Stock this product")
	}
	return r.db.Model(&entities.Stock{}).Where("product_id = ?", id).Updates(payload.IncomingStock(stock)).Error
}

func (r *Stock) SaveDeleteStock(id uint) error {
	return r.db.Delete(&entities.Stock{}, id).Error
}

func (r *Stock) SaveGetQtyAllProduct() ([]payload.OutgoingStock, error) {
	var stocks []payload.OutgoingStock
	err := r.db.Find(&stocks).Error
	return stocks, err
}

func (r *Stock) SaveGetQtyByIDProduct(id uint) (entities.Stock, error) {
	var stock entities.Stock
	err := r.db.First(&stock, id).Error
	return stock, err
}
