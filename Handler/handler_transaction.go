package Handler

import (
	"awesomeProject/Usecase"
	"awesomeProject/payload"
	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	UseCase Usecase.TransactionUseCaseI
}

func NewTransactionHandler(u Usecase.TransactionUseCaseI) TransactionHandlerI {
	return &TransactionHandler{UseCase: u}
}

func (h *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	transactionPayload := new(payload.RequestTransaction)
	if err := c.BodyParser(transactionPayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	transaction := transactionPayload.ToTransaction()
	if !transaction.IsValidCountry(transaction.OrderAddress) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Dont have this country"})
	}
	err := h.UseCase.CreateTransaction(transaction)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(transactionPayload)
}
