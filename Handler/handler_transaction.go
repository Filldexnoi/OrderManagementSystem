package Handler

import (
	"awesomeProject/Usecase"
	"awesomeProject/adpter"
	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	UseCase Usecase.TransactionUseCaseI
}

func NewTransactionHandler(u Usecase.TransactionUseCaseI) TransactionHandlerI {
	return &TransactionHandler{UseCase: u}
}

func (h *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	transaction := new(adpter.TransactionJsonInput)
	if err := c.BodyParser(transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	orderId := transaction.OrderId
	order, err := h.UseCase.GetOrder(orderId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.UseCase.CreateStock()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(stock)
}
