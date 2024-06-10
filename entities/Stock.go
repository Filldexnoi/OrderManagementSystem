package entities

import "time"

type Stock struct {
	ProductId      uint      `json:"product_id"`
	QuantitySizeS  uint      `json:"quantity_size_s"`
	QuantitySizeM  uint      `json:"quantity_size_m"`
	QuantitySizeL  uint      `json:"quantity_size_l"`
	QuantitySizeXL uint      `json:"quantity_size_x_l"`
	StockCreatedAt time.Time `json:"stock_created_at"`
	StockUpdatedAt time.Time `json:"stock_updated_at"`
}
