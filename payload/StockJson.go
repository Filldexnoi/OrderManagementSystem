package payload

import "awesomeProject/entities"

type RespondStock struct {
	ProductId uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

type RequestStock struct {
	ProductId uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

func (s *RequestStock) ToStock() entities.Stock {
	return entities.Stock{
		ProductId: s.ProductId,
		Quantity:  s.Quantity,
	}
}

func StockToStockRes(s entities.Stock) RespondStock {
	return RespondStock{
		ProductId: s.ProductId,
		Quantity:  s.Quantity,
	}
}
