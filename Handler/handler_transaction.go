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
	err := h.UseCase.CreateTransaction(transactionPayload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(transactionPayload)
}

func (h *TransactionHandler) GetAllTransactions(c *fiber.Ctx) error {
	Transactions, err := h.UseCase.GetAllTransaction()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(Transactions)
}
