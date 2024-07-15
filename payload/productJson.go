package payload

import (
	"awesomeProject/entities"
)

type RequestProduct struct {
	ProductTypes string  `json:"product_types"`
	ProductName  string  `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
}

type RespondProduct struct {
	ProductId    uint    `json:"product_id"`
	ProductTypes string  `json:"product_types"`
	ProductName  string  `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
}

func (p *RequestProduct) ToProduct() entities.Product {
	return entities.Product{
		ProductTypes: p.ProductTypes,
		ProductName:  p.ProductName,
		ProductPrice: p.ProductPrice,
	}
}

func ProductToRespondProduct(entity entities.Product) RespondProduct {
	return RespondProduct{
		ProductId:    entity.ProductId,
		ProductTypes: entity.ProductTypes,
		ProductName:  entity.ProductName,
		ProductPrice: entity.ProductPrice,
	}
}
