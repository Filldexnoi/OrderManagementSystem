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
	stock := new(payload.OutgoingStock)
	if err := c.BodyParser(stock); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err := h.UseCase.CreateStock(stock.ToStockEntity())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(stock)
}

func (h *StockHandler) GetAllQtyProducts(c *fiber.Ctx) error {
	stocks, err := h.UseCase.GetQtyAllProduct()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(stocks)
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
	return c.JSON(payload.ToStockJson(stock))
}

func (h *StockHandler) UpdateStock(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}
	stock := new(payload.IncomingStockJson)
	if err := c.BodyParser(stock); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err = h.UseCase.UpdateStock(stock.ToStockEntity(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(stock)
}

func (h *StockHandler) DeleteStock(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}
	err = h.UseCase.DeleteStock(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
