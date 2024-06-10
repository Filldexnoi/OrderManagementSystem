package Repo

import (
	"awesomeProject/entities"
	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) Repo {
	return &ProductRepo{db: db}
}

func (r *ProductRepo) SaveCreateProduct(product *entities.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepo) SaveUpdateProduct(product *entities.Product) error {
	return r.db.Model(&entities.Product{}).Where("product_id = ?", product.ProductId).Updates(product).Error
}

func (r *ProductRepo) SaveDeleteProduct(id uint) error {
	return r.db.Delete(&entities.Product{}, id).Error
}

func (r *ProductRepo) SaveGetAllProduct() ([]entities.Product, error) {
	var products []entities.Product
	err := r.db.Find(&products).Error
	return products, err
}
func (r *ProductRepo) SaveGetByIDProduct(id uint) (entities.Product, error) {
	var product entities.Product
	err := r.db.First(&product, id).Error
	return product, err
}
