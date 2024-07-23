package Handler

import (
	"awesomeProject/Usecase"
	"awesomeProject/observability/logs"
	"awesomeProject/payload"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

type StockHandler struct {
	UseCase Usecase.StockUseCaseI
}

func NewStockHandler(useCase Usecase.StockUseCaseI) *StockHandler {
	return &StockHandler{UseCase: useCase}
}

func (h *StockHandler) CreateStock(c *fiber.Ctx) error {
	ctx, sp := otel.Tracer("stock").Start(c.UserContext(), "stockCreateHandler")
	defer sp.End()
	stock := new(payload.RequestStock)
	if err := c.BodyParser(stock); err != nil {
		logs.LogError("Failed to bodyParser stock request", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	stockEntity, err := h.UseCase.CreateStock(ctx, stock.ToStock())
	if err != nil {
		logs.LogError("Failed to create stock", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	logs.LogInfo("Stock created successfully", logrus.Fields{"product_id": stockEntity.ProductId})
	ResStock := payload.StockToStockRes(stockEntity)
	h.SetSubAttributesWithJson(ResStock, sp)
	return c.Status(fiber.StatusCreated).JSON(ResStock)
}

func (h *StockHandler) GetAllQtyProducts(c *fiber.Ctx) error {
	ctx, sp := otel.Tracer("stock").Start(c.UserContext(), "stockGetAllHandler")
	defer sp.End()
	stocks, err := h.UseCase.GetQtyAllProduct(ctx)
	if err != nil {
		logs.LogError("Failed to get all stocks", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	logs.LogInfo("Get all stocks successfully", nil)
	var ResStocks []*payload.RespondStock
	for _, stock := range stocks {
		ResStocks = append(ResStocks, payload.StockToStockRes(stock))
	}
	h.SetSubAttributesWithJson(ResStocks, sp)
	return c.Status(fiber.StatusOK).JSON(ResStocks)
}

func (h *StockHandler) GetQtyProductByID(c *fiber.Ctx) error {
	ctx, sp := otel.Tracer("stock").Start(c.UserContext(), "stockGetByIDHandler")
	defer sp.End()
	id, err := c.ParamsInt("id")
	if err != nil {
		logs.LogError("Failed to params int product id", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}
	stock, err := h.UseCase.GetQtyByIDProduct(ctx, uint(id))
	if err != nil {
		logs.LogError("Failed to get stock by product id", logrus.Fields{"error": err.Error(), "product_id": id})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	logs.LogInfo("Get stock by product ID successfully", logrus.Fields{"product_id": stock.ProductId})
	ResStock := payload.StockToStockRes(stock)
	h.SetSubAttributesWithJson(ResStock, sp)
	return c.Status(fiber.StatusOK).JSON(ResStock)
}

func (h *StockHandler) UpdateStock(c *fiber.Ctx) error {
	ctx, sp := otel.Tracer("stock").Start(c.UserContext(), "stockUpdateHandler")
	defer sp.End()
	id, err := c.ParamsInt("id")
	if err != nil {
		logs.LogError("Failed to params int product id", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}
	stock := new(payload.RequestStock)
	if err := c.BodyParser(stock); err != nil {
		logs.LogError("Failed to bodyParser stock request", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	stockEntity, err := h.UseCase.UpdateStock(ctx, stock.ToStock(), uint(id))
	if err != nil {
		logs.LogError("Failed to update stock", logrus.Fields{"error": err.Error(), "product_id": stock.ProductId})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	logs.LogInfo("Stock updated successfully", logrus.Fields{"product_id": stockEntity.ProductId})
	ResStock := payload.StockToStockRes(stockEntity)
	h.SetSubAttributesWithJson(ResStock, sp)
	return c.Status(fiber.StatusOK).JSON(ResStock)
}

func (h *StockHandler) DeleteStock(c *fiber.Ctx) error {
	ctx, sp := otel.Tracer("stock").Start(c.UserContext(), "stockDeleteHandler")
	defer sp.End()
	id, err := c.ParamsInt("id")
	if err != nil {
		logs.LogError("Failed to params int product id", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}
	stockEntity, err := h.UseCase.DeleteStock(ctx, uint(id))
	if err != nil {
		logs.LogError("Failed to delete stock", logrus.Fields{"error": err.Error(), "product_id": id})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	logs.LogInfo("Stock deleted successfully", logrus.Fields{"product_id": stockEntity.ProductId})
	ResStock := payload.StockToStockRes(stockEntity)
	h.SetSubAttributesWithJson(ResStock, sp)
	return c.Status(fiber.StatusOK).JSON(ResStock)
}

func (h *StockHandler) SetSubAttributesWithJson(obj any, sp trace.Span) {
	if stocks, ok := obj.([]*payload.RespondStock); ok {
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
	} else if stock, ok := obj.(*payload.RespondStock); ok {
		sp.SetAttributes(
			attribute.Int("ProductID", int(stock.ProductId)),
			attribute.Int("ProductQuantity", int(stock.Quantity)),
		)
	} else {
		sp.RecordError(errors.New("invalid type" + reflect.TypeOf(obj).String()))
	}
}
