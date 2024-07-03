package Handler

import (
	"awesomeProject/Usecase"
	"awesomeProject/payload"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	orderPayload := new(payload.RequestCreateOrder)
	if err := c.BodyParser(orderPayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err := h.UseCase.CreateOrder(orderPayload.ToOrder())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(orderPayload)
}
func (h *OrderHandler) UpdateOrderStatus(c *fiber.Ctx) error {
	idP := c.Params("id")
	id, err := uuid.Parse(idP)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid order ID"})
	}
	order := new(payload.RequestUpdateStatusOrder)
	if err := c.BodyParser(order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err = h.UseCase.UpdateStatusOrder(order.ToOrder(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(order)
}

func (h *OrderHandler) GetAllOrders(c *fiber.Ctx) error {
	orders, err := h.UseCase.GetAllOrders()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	var ResOrders []*payload.ResponseOrder
	for _, order := range orders {
		ResOrders = append(ResOrders, payload.OrderToOrderRespond(order))
	}
	return c.JSON(ResOrders)
}
