package payload

import (
	"awesomeProject/entities"
)

type IncomingProduct struct {
	ProductTypes string  `json:"product_types"`
	ProductName  string  `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
}

type OutgoingProduct struct {
	ProductId    uint    `json:"product_id"`
	ProductTypes string  `json:"product_types"`
	ProductName  string  `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
}

func (p *OutgoingProduct) TableName() string {
	return "products"
}
func (p *IncomingProduct) TableName() string {
	return "products"
}
func (p *IncomingProduct) ToProduct() *entities.Product {
	return &entities.Product{
		ProductTypes: p.ProductTypes,
		ProductName:  p.ProductName,
		ProductPrice: p.ProductPrice,
	}
}

func ToInComingProduct(p *entities.Product) *IncomingProduct {
	return &IncomingProduct{
		ProductTypes: p.ProductTypes,
		ProductName:  p.ProductName,
		ProductPrice: p.ProductPrice,
	}
}
func ToOutgoingProduct(p entities.Product) OutgoingProduct {
	return OutgoingProduct{
		ProductId:    p.ProductId,
		ProductTypes: p.ProductTypes,
		ProductName:  p.ProductName,
		ProductPrice: p.ProductPrice,
	}
}
