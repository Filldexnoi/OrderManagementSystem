package payload

import "awesomeProject/entities"

type OutgoingStock struct {
	ProductId      uint `json:"product_id"`
	QuantitySizeS  uint `json:"quantity_size_s"`
	QuantitySizeM  uint `json:"quantity_size_m"`
	QuantitySizeL  uint `json:"quantity_size_l"`
	QuantitySizeXL uint `json:"quantity_size_xl"`
}

type IncomingStockJson struct {
	QuantitySizeS  uint `json:"quantity_size_s"`
	QuantitySizeM  uint `json:"quantity_size_m"`
	QuantitySizeL  uint `json:"quantity_size_l"`
	QuantitySizeXL uint `json:"quantity_size_xl"`
}

func (s *OutgoingStock) TableName() string {
	return "stocks"
}

func (s *IncomingStockJson) ToStockEntity() *entities.Stock {
	return &entities.Stock{
		QuantitySizeS:  s.QuantitySizeS,
		QuantitySizeM:  s.QuantitySizeM,
		QuantitySizeL:  s.QuantitySizeL,
		QuantitySizeXL: s.QuantitySizeXL,
	}
}

func (s *OutgoingStock) ToStockEntity() *entities.Stock {
	return &entities.Stock{
		ProductId:      s.ProductId,
		QuantitySizeS:  s.QuantitySizeS,
		QuantitySizeM:  s.QuantitySizeM,
		QuantitySizeL:  s.QuantitySizeL,
		QuantitySizeXL: s.QuantitySizeXL,
	}
}

func ToStockJson(p entities.Stock) OutgoingStock {
	return OutgoingStock{
		ProductId:      p.ProductId,
		QuantitySizeS:  p.QuantitySizeS,
		QuantitySizeM:  p.QuantitySizeM,
		QuantitySizeL:  p.QuantitySizeL,
		QuantitySizeXL: p.QuantitySizeXL,
	}
}

func IncomingStock(s *entities.Stock) IncomingStockJson {
	return IncomingStockJson{
		QuantitySizeS:  s.QuantitySizeS,
		QuantitySizeM:  s.QuantitySizeM,
		QuantitySizeL:  s.QuantitySizeL,
		QuantitySizeXL: s.QuantitySizeXL,
	}
}
