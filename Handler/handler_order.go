package Handler

import (
	"awesomeProject/Usecase"
	"awesomeProject/observability/logs"
	"awesomeProject/payload"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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
		logs.LogError("Failed to bodyParser order request", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	order, err := h.UseCase.CreateOrder(orderPayload.ToOrder())
	if err != nil {
		logs.LogError("Failed to create order", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	logs.LogInfo("Create order successfully", logrus.Fields{"order_id": order.OrderId})
	ResOrder := payload.OrderToOrderRespond(order)
	return c.Status(fiber.StatusCreated).JSON(ResOrder)
}
func (h *OrderHandler) UpdateOrderStatus(c *fiber.Ctx) error {
	idP := c.Params("id")
	id, err := uuid.Parse(idP)
	if err != nil {
		logs.LogError("Failed to uuid parse", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid order ID"})
	}
	ReqOrder := new(payload.RequestUpdateStatusOrder)
	if err := c.BodyParser(ReqOrder); err != nil {
		logs.LogError("Failed to bodyParser order request", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	order, err := h.UseCase.UpdateStatusOrder(ReqOrder.ToOrder(), id)
	if err != nil {
		logs.LogError("Failed to update order status", logrus.Fields{"error": err.Error(), "order_id": id})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	logs.LogInfo("Update order successfully", logrus.Fields{"order_id": order.OrderId})
	ResOrder := payload.OrderToOrderRespond(order)
	return c.Status(fiber.StatusOK).JSON(ResOrder)
}

func (h *OrderHandler) GetAllOrders(c *fiber.Ctx) error {
	orders, err := h.UseCase.GetAllOrders()
	if err != nil {
		logs.LogError("Failed to get all orders", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	logs.LogInfo("Get all orders successfully", nil)
	var ResOrders []*payload.ResponseOrder
	for _, order := range orders {
		ResOrders = append(ResOrders, payload.OrderToOrderRespond(order))
	}
	return c.Status(fiber.StatusOK).JSON(ResOrders)
}
