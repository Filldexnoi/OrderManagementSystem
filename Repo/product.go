package Repo

import (
	"awesomeProject/entities"
	"gorm.io/gorm"
)

type ProductDB struct {
	db *gorm.DB
}

type Product struct {
	ProductTypes string
	ProductName  string
	ProductPrice float64
}

func ToProduct(p *entities.Product) *Product {
	return &Product{
		ProductTypes: p.ProductTypes,
		ProductName:  p.ProductName,
		ProductPrice: p.ProductPrice,
	}
}

func NewProductRepo(db *gorm.DB) ProductRepoI {
	return &ProductDB{db: db}
}

func (r *ProductDB) SaveCreateProduct(product *entities.Product) error {
	return r.db.Create(ToProduct(product)).Error
}

func (r *ProductDB) SaveUpdateProduct(product *entities.Product) error {
	return r.db.Model(&entities.Product{}).Where("product_id = ?", product.ProductId).Updates(product).Error
}

func (r *ProductDB) SaveDeleteProduct(id uint) error {
	return r.db.Delete(&entities.Product{}, id).Error
}

func (r *ProductDB) SaveGetAllProduct() ([]entities.Product, error) {
	var products []entities.Product
	err := r.db.Find(&products).Error
	return products, err
}
func (r *ProductDB) SaveGetByIDProduct(id uint) (entities.Product, error) {
	var product entities.Product
	err := r.db.First(&product, id).Error
	return product, err
}
