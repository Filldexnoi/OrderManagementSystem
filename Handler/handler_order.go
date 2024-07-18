package Handler

import (
	"awesomeProject/Usecase"
	"awesomeProject/observability/logs"
	"awesomeProject/payload"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"reflect"
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
	ctx, sp := otel.Tracer("order").Start(c.Context(), "orderCreateHandler")
	defer sp.End()
	orderPayload := new(payload.RequestCreateOrder)
	if err := c.BodyParser(orderPayload); err != nil {
		logs.LogError("Failed to bodyParser order request", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	order, err := h.UseCase.CreateOrder(ctx, orderPayload.ToOrder())
	if err != nil {
		logs.LogError("Failed to create order", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	logs.LogInfo("Create order successfully", logrus.Fields{"order_id": order.OrderId})
	ResOrder := payload.OrderToOrderRespond(order)
	h.SetSubAttributesWithJson(ResOrder, sp)
	return c.Status(fiber.StatusCreated).JSON(ResOrder)
}
func (h *OrderHandler) UpdateOrderStatus(c *fiber.Ctx) error {
	ctx, sp := otel.Tracer("order").Start(c.Context(), "orderUpdateHandler")
	defer sp.End()
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
	order, err := h.UseCase.UpdateStatusOrder(ctx, ReqOrder.ToOrder(), id)
	if err != nil {
		logs.LogError("Failed to update order status", logrus.Fields{"error": err.Error(), "order_id": id})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	logs.LogInfo("Update order successfully", logrus.Fields{"order_id": order.OrderId})
	ResOrder := payload.OrderToOrderRespond(order)
	h.SetSubAttributesWithJson(ResOrder, sp)
	return c.Status(fiber.StatusOK).JSON(ResOrder)
}

func (h *OrderHandler) GetAllOrders(c *fiber.Ctx) error {
	ctx, sp := otel.Tracer("order").Start(c.Context(), "orderGetAllHandler")
	defer sp.End()
	orders, err := h.UseCase.GetAllOrders(ctx)
	if err != nil {
		logs.LogError("Failed to get all orders", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	logs.LogInfo("Get all orders successfully", nil)
	var ResOrders []*payload.ResponseOrder
	for _, order := range orders {
		ResOrders = append(ResOrders, payload.OrderToOrderRespond(order))
	}
	h.SetSubAttributesWithJson(ResOrders, sp)
	return c.Status(fiber.StatusOK).JSON(ResOrders)
}

func (h *OrderHandler) SetSubAttributesWithJson(obj any, sp trace.Span) {
	if orders, ok := obj.([]*payload.ResponseOrder); ok {
		var orderID []string
		var transactionID []string
		var status []string

		for _, order := range orders {
			orderID = append(orderID, order.OrderId.String())
			transactionID = append(transactionID, order.TransactionId.String())
			status = append(status, order.Status)
		}
		sp.SetAttributes(
			attribute.StringSlice("orderID", orderID),
			attribute.StringSlice("transactionID", transactionID),
			attribute.StringSlice("status", status),
		)
	} else if order, ok := obj.(*payload.ResponseOrder); ok {
		sp.SetAttributes(
			attribute.String("orderID", order.OrderId.String()),
			attribute.String("transactionID", order.TransactionId.String()),
			attribute.String("status", order.Status),
		)
	} else {
		sp.RecordError(errors.New("invalid type" + reflect.TypeOf(obj).String()))
	}
}
