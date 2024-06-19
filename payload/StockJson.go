package payload

import "awesomeProject/entities"

type OutgoingStock struct {
	ProductId uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

type IncomingStockJson struct {
	ProductId uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

func (s *OutgoingStock) TableName() string {
	return "stocks"
}
func (s *IncomingStockJson) TableName() string {
	return "stocks"
}

func (s *IncomingStockJson) ToStockEntity() *entities.Stock {
	return &entities.Stock{
		ProductId: s.ProductId,
		Quantity:  s.Quantity,
	}
}

func (s *OutgoingStock) ToStockEntity() *entities.Stock {
	return &entities.Stock{
		ProductId: s.ProductId,
		Quantity:  s.Quantity,
	}
}

func ToStockJson(p entities.Stock) OutgoingStock {
	return OutgoingStock{
		ProductId: p.ProductId,
		Quantity:  p.Quantity,
	}
}

func IncomingStock(s *entities.Stock) IncomingStockJson {
	return IncomingStockJson{
		ProductId: s.ProductId,
		Quantity:  s.Quantity,
	}
}
