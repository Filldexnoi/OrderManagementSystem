package adpter

import (
	"awesomeProject/entities"
)

type ProductBrowserInput struct {
	ProductTypes string  `json:"product_types"`
	ProductName  string  `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
}

type ProductBrowserOutput struct {
	ProductId    uint    `json:"product_id"`
	ProductTypes string  `json:"product_types"`
	ProductName  string  `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
}

func (p *ProductBrowserOutput) TableName() string {
	return "products"
}
func (p *ProductBrowserInput) ToProduct() *entities.Product {
	return &entities.Product{
		ProductTypes: p.ProductTypes,
		ProductName:  p.ProductName,
		ProductPrice: p.ProductPrice,
	}
}

func ToProduct(p entities.Product) ProductBrowserOutput {
	return ProductBrowserOutput{
		ProductId:    p.ProductId,
		ProductTypes: p.ProductTypes,
		ProductName:  p.ProductName,
		ProductPrice: p.ProductPrice,
	}
}
