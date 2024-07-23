package Repo

import (
	"awesomeProject/entities"
	"awesomeProject/models"
	"context"
	"errors"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
	"reflect"
)

type Stock struct {
	db *gorm.DB
}

func NewStock(db *gorm.DB) StockRepoI {
	return &Stock{db: db}
}

func (r *Stock) SaveCreateStock(ctx context.Context, stock entities.Stock) (*entities.Stock, error) {
	_, sp := otel.Tracer("stock").Start(ctx, "StockCreateRepository")
	defer sp.End()
	createStock := models.StockToGormStock(stock)
	err := r.db.Create(&createStock).Error
	if err != nil {
		return nil, err
	}
	stockEntity := createStock.ToStock()
	r.SetSubAttributesWithJson(stockEntity, sp)
	return stockEntity, nil
}

func (r *Stock) SaveUpdateStock(ctx context.Context, stock entities.Stock) (*entities.Stock, error) {
	_, sp := otel.Tracer("stock").Start(ctx, "StockUpdateRepository")
	defer sp.End()
	updateStock := models.StockToGormStock(stock)
	err := r.db.Model(&models.Stock{}).Where("product_id = ?", updateStock.ProductID).Updates(&updateStock).Error
	if err != nil {
		return nil, err
	}
	stockEntity := updateStock.ToStock()
	r.SetSubAttributesWithJson(stockEntity, sp)
	return stockEntity, nil
}

func (r *Stock) SaveDeleteStock(ctx context.Context, id uint) (*entities.Stock, error) {
	_, sp := otel.Tracer("stock").Start(ctx, "StockDeleteRepository")
	defer sp.End()
	var deleteStock models.Stock
	if err := r.db.First(&deleteStock, id).Error; err != nil {
		return nil, err
	}

	if err := r.db.Delete(&deleteStock, id).Error; err != nil {
		return nil, err
	}
	stockEntity := deleteStock.ToStock()
	r.SetSubAttributesWithJson(stockEntity, sp)
	return stockEntity, nil
}

func (r *Stock) SaveGetQtyAllProduct(ctx context.Context) ([]*entities.Stock, error) {
	_, sp := otel.Tracer("stock").Start(ctx, "StockGetAllRepository")
	defer sp.End()
	var stocksGorm []*models.Stock
	err := r.db.Find(&stocksGorm).Error
	var stocks []*entities.Stock
	for _, stock := range stocksGorm {
		stocks = append(stocks, stock.ToStock())
	}
	r.SetSubAttributesWithJson(stocks, sp)
	return stocks, err
}

func (r *Stock) SaveGetQtyByIDProduct(ctx context.Context, id uint) (*entities.Stock, error) {
	_, sp := otel.Tracer("stock").Start(ctx, "StockGetByIdRepository")
	defer sp.End()
	var stockGorm *models.Stock
	err := r.db.First(&stockGorm, id).Error
	stock := stockGorm.ToStock()
	r.SetSubAttributesWithJson(stock, sp)
	return stock, err
}

func (r *Stock) CheckStockToCreateOrder(ctx context.Context, transaction *entities.Transaction) error {
	_, sp := otel.Tracer("order").Start(ctx, "CheckStockToCreateOrderRepository")
	defer sp.End()
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, item := range transaction.Items {
			result := tx.Model(&models.Stock{}).Where("product_id = ? AND quantity >= ?", item.ProductId, item.Quantity).
				Update("quantity", gorm.Expr("quantity - ?", item.Quantity))
			if result.Error != nil {
				return result.Error
			}

			if result.RowsAffected == 0 {
				return errors.New(fmt.Sprintf("ไม่สามารถลดจำนวนสินค้าได้เนื่องจากสินค้า ID %d มีจำนวนในสต็อกไม่เพียงพอ", item.ProductId))
			}
		}
		r.SetTransactionSubAttributes(transaction, sp)
		return nil
	})
}

func (r *Stock) SetTransactionSubAttributes(transaction *entities.Transaction, sp trace.Span) {
	var itemID = make([]int, len(transaction.Items))
	var itemQuantity = make([]int, len(transaction.Items))

	for _, item := range transaction.Items {
		itemID = append(itemID, int(item.ProductId))
		itemQuantity = append(itemQuantity, int(item.Quantity))
	}
	sp.SetAttributes(
		attribute.String("TransactionID", transaction.TransactionId.String()),
		attribute.String("TransactionOrderAddress", transaction.OrderAddress),
		attribute.IntSlice("ItemID", itemID),
		attribute.IntSlice("ItemQuantity", itemQuantity),
	)
  
func (r *Stock) SetSubAttributesWithJson(obj any, sp trace.Span) {
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
