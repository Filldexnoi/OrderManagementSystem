package entities

type Stock struct {
	ProductId uint
	Quantity  uint
}

func (s *Stock) IsEqualCreatedStock(createdStock *Stock) bool {
	if s.ProductId != createdStock.ProductId || s.Quantity != createdStock.Quantity {
		return false
	}
	return true
}

func (s *Stock) IsEqualUpdatedStock(updatedStock *Stock) bool {
	if s.ProductId != updatedStock.ProductId || s.Quantity != updatedStock.Quantity {
		return false
	}
	return true
}
