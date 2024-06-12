package adpter

import "awesomeProject/entities"

type CRStockJson struct {
	ProductId      uint `json:"product_id"`
	QuantitySizeS  uint `json:"quantity_size_s"`
	QuantitySizeM  uint `json:"quantity_size_m"`
	QuantitySizeL  uint `json:"quantity_size_l"`
	QuantitySizeXL uint `json:"quantity_size_xl"`
}

type UStockJson struct {
	QuantitySizeS  uint `json:"quantity_size_s"`
	QuantitySizeM  uint `json:"quantity_size_m"`
	QuantitySizeL  uint `json:"quantity_size_l"`
	QuantitySizeXL uint `json:"quantity_size_xl"`
}

func (s *CRStockJson) TableName() string {
	return "stocks"
}

func (s *UStockJson) ToStockEntity() *entities.Stock {
	return &entities.Stock{
		QuantitySizeS:  s.QuantitySizeS,
		QuantitySizeM:  s.QuantitySizeM,
		QuantitySizeL:  s.QuantitySizeL,
		QuantitySizeXL: s.QuantitySizeXL,
	}
}

func (s *CRStockJson) ToStockEntity() *entities.Stock {
	return &entities.Stock{
		ProductId:      s.ProductId,
		QuantitySizeS:  s.QuantitySizeS,
		QuantitySizeM:  s.QuantitySizeM,
		QuantitySizeL:  s.QuantitySizeL,
		QuantitySizeXL: s.QuantitySizeXL,
	}
}

func ToStockJson(p entities.Stock) CRStockJson {
	return CRStockJson{
		ProductId:      p.ProductId,
		QuantitySizeS:  p.QuantitySizeS,
		QuantitySizeM:  p.QuantitySizeM,
		QuantitySizeL:  p.QuantitySizeL,
		QuantitySizeXL: p.QuantitySizeXL,
	}
}
