package adpter

import "awesomeProject/entities"

type ProductDatabase struct {
	ProductTypes string
	ProductName  string
	ProductPrice float64
}

func (p *ProductDatabase) TableName() string {
	return "products"
}
func ToProductDatabase(p *entities.Product) *ProductDatabase {
	return &ProductDatabase{
		ProductTypes: p.ProductTypes,
		ProductName:  p.ProductName,
		ProductPrice: p.ProductPrice,
	}
}
