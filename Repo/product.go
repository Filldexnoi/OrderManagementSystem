package Repo

import (
	"awesomeProject/entities"
	"awesomeProject/models"
	"context"
	"errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
	"reflect"
)

type ProductDB struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepoI {
	return &ProductDB{db: db}
}

func (r *ProductDB) SaveCreateProduct(ctx context.Context, product entities.Product) (*entities.Product, error) {
	_, sp := otel.Tracer("product").Start(ctx, "productCreateRepository")
	defer sp.End()
	createProduct := models.ProductToGormProduct(product)
	err := r.db.Create(&createProduct).Error
	if err != nil {
		return nil, err
	}
	productEntity := createProduct.ToProduct()
	r.SetProductSubAttributes(productEntity, sp)
	return productEntity, nil
}

func (r *ProductDB) SaveUpdateProduct(ctx context.Context, product entities.Product) (*entities.Product, error) {
	_, sp := otel.Tracer("product").Start(ctx, "productUpdateRepository")
	defer sp.End()
	updateProduct := models.ProductToGormProduct(product)
	err := r.db.Model(&models.Product{}).Where("product_id = ?", product.ProductId).Updates(&updateProduct).Error
	if err != nil {
		return nil, err
	}
	productEntity := updateProduct.ToProduct()
	r.SetProductSubAttributes(productEntity, sp)
	return productEntity, nil
}

func (r *ProductDB) SaveDeleteProduct(ctx context.Context, id uint) (*entities.Product, error) {
	_, sp := otel.Tracer("product").Start(ctx, "productDeleteRepository")
	defer sp.End()
	var deletedProduct models.Product

	if err := r.db.First(&deletedProduct, id).Error; err != nil {
		return nil, err
	}

	if err := r.db.Delete(&deletedProduct, id).Error; err != nil {
		return nil, err
	}
	productEntity := deletedProduct.ToProduct()
	r.SetProductSubAttributes(productEntity, sp)
	return productEntity, nil
}

func (r *ProductDB) SaveGetAllProduct(ctx context.Context) ([]*entities.Product, error) {
	_, sp := otel.Tracer("product").Start(ctx, "productGetAllRepository")
	defer sp.End()
	var productsGorm []models.Product
	err := r.db.Find(&productsGorm).Error
	var product []*entities.Product
	for _, p := range productsGorm {
		product = append(product, p.ToProduct())
	}
	r.SetProductSubAttributes(product, sp)
	return product, err
}
func (r *ProductDB) SaveGetByIDProduct(ctx context.Context, id uint) (*entities.Product, error) {
	_, sp := otel.Tracer("product").Start(ctx, "productGetByIDRepository")
	defer sp.End()
	var productGorm models.Product
	err := r.db.First(&productGorm, id).Error
	product := productGorm.ToProduct()
	r.SetProductSubAttributes(product, sp)
	return product, err
}

func (r *ProductDB) GetPriceProducts(ctx context.Context, transaction *entities.Transaction) (*entities.Transaction, error) {
	_, sp := otel.Tracer("product").Start(ctx, "productGetPriceRepository")
	defer sp.End()
	var product models.Product
	for i, item := range transaction.Items {
		err := r.db.Model(&models.Product{}).Select("product_price").Where("product_id = ?", item.ProductId).First(&product).Error
		if err != nil {
			return nil, errors.New("product not found")
		}
		transaction.Items[i].Price = product.ProductPrice
	}
	r.SetTransactionSubAttributes(transaction, sp)
	return transaction, nil
}

func (r *ProductDB) SetProductSubAttributes(productData any, sp trace.Span) {
	if products, ok := productData.([]*entities.Product); ok {
		productIDs := make([]int, len(products))
		productNames := make([]string, len(products))
		productTypes := make([]string, len(products))
		productPrices := make([]float64, len(products))

		for _, product := range products {
			productIDs = append(productIDs, int(product.ProductId))
			productNames = append(productNames, product.ProductName)
			productTypes = append(productTypes, product.ProductTypes)
			productPrices = append(productPrices, product.ProductPrice)
		}

		sp.SetAttributes(
			attribute.IntSlice("ProductID", productIDs),
			attribute.StringSlice("ProductName", productNames),
			attribute.StringSlice("ProductTypes", productTypes),
			attribute.Float64Slice("ProductPrice", productPrices),
		)
	} else if product, ok := productData.(*entities.Product); ok {
		sp.SetAttributes(
			attribute.Int("ProductID", int(product.ProductId)),
			attribute.String("ProductName", product.ProductName),
			attribute.String("ProductTypes", product.ProductTypes),
			attribute.Float64("ProductPrice", product.ProductPrice),
		)
	} else {
		sp.RecordError(errors.New("invalid type" + reflect.TypeOf(productData).String()))
	}
}

func (r *ProductDB) SetTransactionSubAttributes(transaction *entities.Transaction, sp trace.Span) {
	productIDs := make([]int, len(transaction.Items))
	productQuantity := make([]int, len(transaction.Items))
	productPrices := make([]float64, len(transaction.Items))

	for _, item := range transaction.Items {
		productIDs = append(productIDs, int(item.ProductId))
		productQuantity = append(productQuantity, int(item.Quantity))
		productPrices = append(productPrices, item.Price)
	}
	sp.SetAttributes(
		attribute.String("TransactionID", transaction.TransactionId.String()),
		attribute.String("TransactionOrderAddress", transaction.OrderAddress),
		attribute.IntSlice("ProductID", productIDs),
		attribute.IntSlice("ProductQuantity", productQuantity),
		attribute.Float64Slice("ProductPrice", productPrices),
		attribute.Float64("TransactionTotalPrice", transaction.TotalPrice),
	)
}
