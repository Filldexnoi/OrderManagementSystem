package Repo

import (
	"awesomeProject/entities"
	"awesomeProject/models"
	"awesomeProject/payload"
	"gorm.io/gorm"
)

type ProductDB struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepoI {
	return &ProductDB{db: db}
}

func (r *ProductDB) SaveCreateProduct(product *entities.Product) error {
	return r.db.Create(payload.ToInComingProduct(product)).Error
}

func (r *ProductDB) SaveUpdateProduct(product *entities.Product, id uint) error {
	return r.db.Model(&models.Product{}).Where("product_id = ?", id).Updates(payload.ToInComingProduct(product)).Error
}

func (r *ProductDB) SaveDeleteProduct(id uint) error {
	return r.db.Delete(&entities.Product{}, id).Error
}

func (r *ProductDB) SaveGetAllProduct() ([]payload.OutgoingProduct, error) {
	var products []payload.OutgoingProduct
	err := r.db.Find(&products).Error
	return products, err
}
func (r *ProductDB) SaveGetByIDProduct(id uint) (entities.Product, error) {
	var product entities.Product
	err := r.db.First(&product, id).Error
	return product, err
}
