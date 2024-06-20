package Handler

import (
	"awesomeProject/Usecase"
	"awesomeProject/payload"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	UseCase Usecase.OrderUseCaseI
}

func NewOrderHandler(orderUseCase Usecase.OrderUseCaseI) OrderHandlerI {
	return &OrderHandler{
		UseCase: orderUseCase,
	}
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	orderPayload := new(payload.RequestOrder)
	if err := c.BodyParser(orderPayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err := h.UseCase.CreateOrder(orderPayload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(orderPayload)
}
