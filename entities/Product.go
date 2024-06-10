package entities

import (
	"awesomeProject/Handler"
	"time"
)

type Product struct {
	ProductId        uint
	ProductTypes     string
	ProductName      string
	ProductPrice     float64
	ProductCreatedAt time.Time
	ProductUpdatedAt time.Time
}

func (p *Product) JsonToProduct(Json Handler.ProductJson) Product {
	return Product{
		ProductTypes: Json.ProductTypes,
		ProductName:  Json.ProductName,
		ProductPrice: Json.ProductPrice,
	}
}
