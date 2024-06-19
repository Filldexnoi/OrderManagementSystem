package Repo

import (
	"awesomeProject/entities"
	"awesomeProject/models"
	"errors"
	"gorm.io/gorm"
	"time"
)

type ProductDB struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepoI {
	return &ProductDB{db: db}
}

func (r *ProductDB) SaveCreateProduct(product *entities.Product) error {
	gormProduct := models.ProductToGormProduct(product)
	gormProduct.ProductCreatedAt = time.Now()
	err := r.db.Create(gormProduct).Error
	if err != nil {
		return errors.New("cannot create product")
	}
	return nil
}

func (r *ProductDB) SaveUpdateProduct(product *entities.Product) error {
	gormProduct := models.ProductToGormProduct(product)
	gormProduct.ProductUpdatedAt = time.Now()
	return r.db.Model(&models.Product{}).Where("product_id = ?", product.ProductId).Updates(gormProduct).Error
}

func (r *ProductDB) SaveDeleteProduct(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}

func (r *ProductDB) SaveGetAllProduct() ([]*entities.Product, error) {
	var productsGorm []models.Product
	err := r.db.Find(&productsGorm).Error
	var product []*entities.Product
	for _, p := range productsGorm {
		product = append(product, p.ToProduct())
	}
	return product, err
}
func (r *ProductDB) SaveGetByIDProduct(id uint) (*entities.Product, error) {
	var productGorm models.Product
	err := r.db.First(&productGorm, id).Error
	product := productGorm.ToProduct()
	return product, err
}
