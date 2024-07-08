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
	order, err := h.UseCase.CreateOrder(orderPayload.ToOrder())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	ResOrder := payload.OrderToOrderRespond(order)
	return c.Status(fiber.StatusCreated).JSON(ResOrder)
}
func (h *OrderHandler) UpdateOrderStatus(c *fiber.Ctx) error {
	idP := c.Params("id")
	id, err := uuid.Parse(idP)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid order ID"})
	}
	ReqOrder := new(payload.RequestUpdateStatusOrder)
	if err := c.BodyParser(ReqOrder); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	order, err := h.UseCase.UpdateStatusOrder(ReqOrder.ToOrder(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	ResOrder := payload.OrderToOrderRespond(order)
	return c.JSON(ResOrder)
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
