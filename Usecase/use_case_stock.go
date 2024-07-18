package Usecase

import (
	"awesomeProject/Repo"
	"awesomeProject/entities"
	"context"
	"errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

type StockUseCase struct {
	repo Repo.StockRepoI
}

func NewStockUseCase(repo Repo.StockRepoI) StockUseCaseI {
	return &StockUseCase{repo: repo}
}

func (s *StockUseCase) CreateStock(ctx context.Context, stock entities.Stock) (*entities.Stock, error) {
	ctx, sp := otel.Tracer("stock").Start(ctx, "StockCreateUseCase")
	defer sp.End()
	createStock, err := s.repo.SaveCreateStock(ctx, stock)
	if err != nil {
		return nil, err
	}
	s.SetSubAttributesWithJson(createStock, sp)
	return createStock, nil
}

func (s *StockUseCase) UpdateStock(ctx context.Context, stock entities.Stock, id uint) (*entities.Stock, error) {
	ctx, sp := otel.Tracer("stock").Start(ctx, "StockUpdateUseCase")
	defer sp.End()
	stock.ProductId = id
	updatedStock, err := s.repo.SaveUpdateStock(ctx, stock)
	if err != nil {
		return nil, err
	}
	s.SetSubAttributesWithJson(updatedStock, sp)
	return updatedStock, nil
}

func (s *StockUseCase) DeleteStock(ctx context.Context, id uint) (*entities.Stock, error) {
	ctx, sp := otel.Tracer("stock").Start(ctx, "StockDeleteUseCase")
	defer sp.End()
	deletedStock, err := s.repo.SaveDeleteStock(ctx, id)
	if err != nil {
		return nil, err
	}
	s.SetSubAttributesWithJson(deletedStock, sp)
	return deletedStock, nil
}

func (s *StockUseCase) GetQtyAllProduct(ctx context.Context) ([]*entities.Stock, error) {
	ctx, sp := otel.Tracer("stock").Start(ctx, "StockGetAllUseCase")
	defer sp.End()
	stocks, err := s.repo.SaveGetQtyAllProduct(ctx)
	if err != nil {
		return nil, err
	}
	s.SetSubAttributesWithJson(stocks, sp)
	return stocks, nil
}

func (s *StockUseCase) GetQtyByIDProduct(ctx context.Context, id uint) (*entities.Stock, error) {
	ctx, sp := otel.Tracer("stock").Start(ctx, "StockGetByIdUseCase")
	defer sp.End()
	stock, err := s.repo.SaveGetQtyByIDProduct(ctx, id)
	if err != nil {
		return nil, err
	}
	s.SetSubAttributesWithJson(stock, sp)
	return stock, err
}

func (s *StockUseCase) SetSubAttributesWithJson(obj any, sp trace.Span) {
	if stocks, ok := obj.([]*entities.Stock); ok {
		var productID []int
		var productQuantity []int

		for _, stock := range stocks {
			productID = append(productID, int(stock.ProductId))
			productQuantity = append(productQuantity, int(stock.Quantity))
		}
		sp.SetAttributes(
			attribute.IntSlice("ProductID", productID),
			attribute.IntSlice("ProductQuantity", productQuantity),
		)
	} else if stock, ok := obj.(*entities.Stock); ok {
		sp.SetAttributes(
			attribute.Int("ProductID", int(stock.ProductId)),
			attribute.Int("ProductQuantity", int(stock.Quantity)),
		)
	} else {
		sp.RecordError(errors.New("invalid type" + reflect.TypeOf(obj).String()))
	}
}
