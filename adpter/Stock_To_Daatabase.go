package adpter

import "awesomeProject/entities"

type UpdateStockData struct {
	QuantitySizeS  uint
	QuantitySizeM  uint
	QuantitySizeL  uint
	QuantitySizeXL uint
}

type CreateStockData struct {
	ProductId      uint
	QuantitySizeS  uint
	QuantitySizeM  uint
	QuantitySizeL  uint
	QuantitySizeXL uint
}

func (s *UpdateStockData) TableName() string {
	return "stocks"
}

func (s *CreateStockData) TableName() string {
	return "stocks"
}
func StockToUpdateStockData(p *entities.Stock) *UpdateStockData {
	return &UpdateStockData{
		QuantitySizeS:  p.QuantitySizeS,
		QuantitySizeM:  p.QuantitySizeM,
		QuantitySizeL:  p.QuantitySizeL,
		QuantitySizeXL: p.QuantitySizeXL,
	}
}

func StockToCreateStockData(p *entities.Stock) *CreateStockData {
	return &CreateStockData{
		ProductId:      p.ProductId,
		QuantitySizeS:  p.QuantitySizeS,
		QuantitySizeM:  p.QuantitySizeM,
		QuantitySizeL:  p.QuantitySizeL,
		QuantitySizeXL: p.QuantitySizeXL,
	}
}
