package Handler

import (
	"awesomeProject/Usecase"
	"awesomeProject/payload"
	"errors"
	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	UseCase Usecase.TransactionUseCaseI
}

func NewTransactionHandler(u Usecase.TransactionUseCaseI) TransactionHandlerI {
	return &TransactionHandler{UseCase: u}
}

func (h *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	transactionPayload := new(payload.IncomingTransaction)
	if err := c.BodyParser(transactionPayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return errors.New("5555")
}
