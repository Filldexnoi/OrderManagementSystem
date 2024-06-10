package Repo

import (
	"awesomeProject/entities"
	"gorm.io/gorm"
)

type Product struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepoI {
	return &Product{db: db}
}

func (r *Product) SaveCreateProduct(product *entities.Product) error {
	return r.db.Create(product).Error
}

func (r *Product) SaveUpdateProduct(product *entities.Product) error {
	return r.db.Model(&entities.Product{}).Where("product_id = ?", product.ProductId).Updates(product).Error
}

func (r *Product) SaveDeleteProduct(id uint) error {
	return r.db.Delete(&entities.Product{}, id).Error
}

func (r *Product) SaveGetAllProduct() ([]entities.Product, error) {
	var products []entities.Product
	err := r.db.Find(&products).Error
	return products, err
}
func (r *Product) SaveGetByIDProduct(id uint) (entities.Product, error) {
	var product entities.Product
	err := r.db.First(&product, id).Error
	return product, err
}
