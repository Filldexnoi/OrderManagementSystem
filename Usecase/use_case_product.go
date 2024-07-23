package Usecase

import (
	"awesomeProject/Repo"
	"awesomeProject/entities"
	"context"
	"errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type ProductUseCase struct {
	repo Repo.ProductRepoI
}

func NewProductUseCase(repo Repo.ProductRepoI) ProductUseCaseI {
	return &ProductUseCase{repo: repo}
}

func (s *ProductUseCase) CreateProduct(ctx context.Context, product entities.Product) (*entities.Product, error) {
	ctx, sp := otel.Tracer("product").Start(ctx, "productCreateUseCase")
	defer sp.End()

	createdProduct, err := s.repo.SaveCreateProduct(ctx, product)
	if err != nil {
		return nil, err
	}
	s.SetProductSubAttributes(createdProduct, sp)
	return createdProduct, nil
}

func (s *ProductUseCase) GetAllProducts(ctx context.Context) ([]*entities.Product, error) {
	ctx, sp := otel.Tracer("product").Start(ctx, "productGetAllUseCase")
	defer sp.End()
	products, err := s.repo.SaveGetAllProduct(ctx)
	if err != nil {
		return nil, err
	}
	s.SetProductSubAttributes(products, sp)
	return products, nil
}

func (s *ProductUseCase) GetByIDProduct(ctx context.Context, id uint) (*entities.Product, error) {
	ctx, sp := otel.Tracer("product").Start(ctx, "productGetByIDUseCase")
	defer sp.End()
	product, err := s.repo.SaveGetByIDProduct(ctx, id)
	if err != nil {
		return nil, err
	}
	s.SetProductSubAttributes(product, sp)
	return product, err
}

func (s *ProductUseCase) UpdateProduct(ctx context.Context, product entities.Product, id uint) (*entities.Product, error) {
	ctx, sp := otel.Tracer("product").Start(ctx, "productUpdateUseCase")
	defer sp.End()
	product.ProductId = id
	updatedProduct, err := s.repo.SaveUpdateProduct(ctx, product)
	if err != nil {
		return nil, err
	}
	s.SetProductSubAttributes(updatedProduct, sp)
	return updatedProduct, nil
}

func (s *ProductUseCase) DeleteProduct(ctx context.Context, id uint) (*entities.Product, error) {
	ctx, sp := otel.Tracer("product").Start(ctx, "productDeleteUseCase")
	defer sp.End()
	deletedProduct, err := s.repo.SaveDeleteProduct(ctx, id)
	if err != nil {
		return nil, err
	}
	s.SetProductSubAttributes(deletedProduct, sp)
	return deletedProduct, nil
}

func (s *ProductUseCase) SetProductSubAttributes(productData any, sp trace.Span) {
	if products, ok := productData.([]*entities.Product); ok {
		productIDs := make([]int, len(products))
		productNames := make([]string, len(products))
		productTypes := make([]string, len(products))
		productPrices := make([]float64, len(products))

		for _, product := range products {
			productIDs = append(productIDs, int(product.ProductId))
			productNames = append(productNames, product.ProductName)
			productTypes = append(productTypes, product.ProductTypes)
			productPrices = append(productPrices, product.ProductPrice)
		}

		sp.SetAttributes(
			attribute.IntSlice("ProductID", productIDs),
			attribute.StringSlice("ProductName", productNames),
			attribute.StringSlice("ProductTypes", productTypes),
			attribute.Float64Slice("ProductPrice", productPrices),
		)
	} else if product, ok := productData.(*entities.Product); ok {
		sp.SetAttributes(
			attribute.Int("ProductID", int(product.ProductId)),
			attribute.String("ProductName", product.ProductName),
			attribute.String("ProductTypes", product.ProductTypes),
			attribute.Float64("ProductPrice", product.ProductPrice),
		)
	} else {
		sp.RecordError(errors.New("invalid type"))
	}
}
