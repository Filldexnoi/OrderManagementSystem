package Repo

import (
	"awesomeProject/entities"
	"awesomeProject/models"
	"errors"
	"gorm.io/gorm"
)

type ProductDB struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepoI {
	return &ProductDB{db: db}
}

func (r *ProductDB) SaveCreateProduct(product entities.Product) (*entities.Product, error) {
	createProduct := models.ProductToGormProduct(product)
	err := r.db.Create(&createProduct).Error
	if err != nil {
		return nil, err
	}
	productEntity := createProduct.ToProduct()
	return productEntity, nil
}

func (r *ProductDB) SaveUpdateProduct(product entities.Product) (*entities.Product, error) {
	updateProduct := models.ProductToGormProduct(product)
	err := r.db.Model(&models.Product{}).Where("product_id = ?", product.ProductId).Updates(&updateProduct).Error
	if err != nil {
		return nil, err
	}
	productEntity := updateProduct.ToProduct()
	return productEntity, nil
}

func (r *ProductDB) SaveDeleteProduct(id uint) (*entities.Product, error) {
	var deletedProduct models.Product

	if err := r.db.First(&deletedProduct, id).Error; err != nil {
		return nil, err
	}

	if err := r.db.Delete(&deletedProduct, id).Error; err != nil {
		return nil, err
	}
	productEntity := deletedProduct.ToProduct()
	return productEntity, nil
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

func (r *ProductDB) GetPriceProducts(transaction *entities.Transaction) (*entities.Transaction, error) {
	var product models.Product
	for i, item := range transaction.Items {
		err := r.db.Model(&models.Product{}).Select("product_price").Where("product_id = ?", item.ProductId).First(&product).Error
		if err != nil {
			return nil, errors.New("product not found")
		}
		transaction.Items[i].Price = product.ProductPrice
	}
	return transaction, nil
}
