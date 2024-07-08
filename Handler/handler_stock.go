package Handler

import (
	"awesomeProject/Usecase"
	"awesomeProject/payload"
	"github.com/gofiber/fiber/v2"
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	stockEntity, err := h.UseCase.CreateStock(stock.ToStock())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	ResStock := payload.StockToStockRes(stockEntity)
	return c.Status(fiber.StatusCreated).JSON(ResStock)
}

func (h *StockHandler) GetAllQtyProducts(c *fiber.Ctx) error {
	stocks, err := h.UseCase.GetQtyAllProduct()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	var ResStocks []*payload.RespondStock
	for _, stock := range stocks {
		ResStocks = append(ResStocks, payload.StockToStockRes(stock))
	}
	return c.JSON(ResStocks)
}

func (h *StockHandler) GetQtyProductByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}
	stock, err := h.UseCase.GetQtyByIDProduct(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	ResStock := payload.StockToStockRes(stock)
	return c.JSON(ResStock)
}

func (h *StockHandler) UpdateStock(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}
	stock := new(payload.RequestStock)
	if err := c.BodyParser(stock); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	stockEntity, err := h.UseCase.UpdateStock(stock.ToStock(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	ResStock := payload.StockToStockRes(stockEntity)
	return c.JSON(ResStock)
}

func (h *StockHandler) DeleteStock(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}
	stockEntity, err := h.UseCase.DeleteStock(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	ResStock := payload.StockToStockRes(stockEntity)
	return c.JSON(ResStock)
}
