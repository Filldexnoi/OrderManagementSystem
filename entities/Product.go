package entities

type Product struct {
	ProductId    uint
	ProductTypes string
	ProductName  string
	ProductPrice float64
}

func (p *Product) IsEqualCreatedProduct(createdProduct *Product) bool {
	if p.ProductName != createdProduct.ProductName || p.ProductTypes != createdProduct.ProductTypes || p.ProductPrice != createdProduct.ProductPrice {
		return false
	}
	return true
}

func (p *Product) IsEqualUpdatedProduct(updatedProduct *Product) bool {
	if p.ProductId != updatedProduct.ProductId || p.ProductName != updatedProduct.ProductName || p.ProductTypes !=
		updatedProduct.ProductTypes || p.ProductPrice != updatedProduct.ProductPrice {
		return false
	}
	return true
}
