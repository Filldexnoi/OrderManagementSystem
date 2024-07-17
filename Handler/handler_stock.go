package Handler

import (
	"awesomeProject/Usecase"
	"awesomeProject/observability/logs"
	"awesomeProject/payload"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type StockHandler struct {
	UseCase Usecase.StockUseCaseI
}

func NewStockHandler(useCase Usecase.StockUseCaseI) *StockHandler {
	return &StockHandler{UseCase: useCase}
}

func (h *StockHandler) CreateStock(c *fiber.Ctx) error {
	stock := new(payload.RequestStock)
	if err := c.BodyParser(stock); err != nil {
		logs.LogError("Failed to bodyParser stock request", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	stockEntity, err := h.UseCase.CreateStock(stock.ToStock())
	if err != nil {
		logs.LogError("Failed to create stock", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	logs.LogInfo("Stock created successfully", logrus.Fields{"product_id": stockEntity.ProductId})
	ResStock := payload.StockToStockRes(stockEntity)
	return c.Status(fiber.StatusCreated).JSON(ResStock)
}

func (h *StockHandler) GetAllQtyProducts(c *fiber.Ctx) error {
	stocks, err := h.UseCase.GetQtyAllProduct()
	if err != nil {
		logs.LogError("Failed to get all stocks", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	logs.LogInfo("Get all stocks successfully", nil)
	var ResStocks []*payload.RespondStock
	for _, stock := range stocks {
		ResStocks = append(ResStocks, payload.StockToStockRes(stock))
	}
	return c.Status(fiber.StatusOK).JSON(ResStocks)
}

func (h *StockHandler) GetQtyProductByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		logs.LogError("Failed to params int product id", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}
	stock, err := h.UseCase.GetQtyByIDProduct(uint(id))
	if err != nil {
		logs.LogError("Failed to get stock by product id", logrus.Fields{"error": err.Error(), "product_id": id})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	logs.LogInfo("Get stock by product ID successfully", logrus.Fields{"product_id": stock.ProductId})
	ResStock := payload.StockToStockRes(stock)
	return c.Status(fiber.StatusOK).JSON(ResStock)
}

func (h *StockHandler) UpdateStock(c *fiber.Ctx) error {
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
	stockEntity, err := h.UseCase.UpdateStock(stock.ToStock(), uint(id))
	if err != nil {
		logs.LogError("Failed to update stock", logrus.Fields{"error": err.Error(), "product_id": stock.ProductId})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	logs.LogInfo("Stock updated successfully", logrus.Fields{"product_id": stockEntity.ProductId})
	ResStock := payload.StockToStockRes(stockEntity)
	return c.Status(fiber.StatusOK).JSON(ResStock)
}

func (h *StockHandler) DeleteStock(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		logs.LogError("Failed to params int product id", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}
	stockEntity, err := h.UseCase.DeleteStock(uint(id))
	if err != nil {
		logs.LogError("Failed to delete stock", logrus.Fields{"error": err.Error(), "product_id": id})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	logs.LogInfo("Stock deleted successfully", logrus.Fields{"product_id": stockEntity.ProductId})
	ResStock := payload.StockToStockRes(stockEntity)
	return c.Status(fiber.StatusOK).JSON(ResStock)
}
