package Handler

import (
	"awesomeProject/Usecase"
	"awesomeProject/observability/logs"
	"awesomeProject/payload"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

type TransactionHandler struct {
	UseCase Usecase.TransactionUseCaseI
}

func NewTransactionHandler(u Usecase.TransactionUseCaseI) TransactionHandlerI {
	return &TransactionHandler{UseCase: u}
}

func (h *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	ctx, sp := otel.Tracer("transaction").Start(c.UserContext(), "TransactionCreateHandler")
	defer sp.End()

	transactionPayload := new(payload.RequestTransaction)
	if err := c.BodyParser(transactionPayload); err != nil {
		logs.LogError("Failed to bodyParser transaction request", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	transaction, err := h.UseCase.CreateTransaction(ctx, transactionPayload.ToTransaction())
	if err != nil {
		logs.LogError("Failed to create transaction", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	logs.LogInfo("Create transaction successfully", logrus.Fields{"transaction_id": transaction.TransactionId})
	ResTransaction := payload.TransactionToResTransaction(transaction)
	h.SetTransactionSubAttributes(ResTransaction, sp)
	return c.Status(fiber.StatusCreated).JSON(ResTransaction)
}

func (h *TransactionHandler) GetAllTransactions(c *fiber.Ctx) error {
	ctx, sp := otel.Tracer("transaction").Start(c.UserContext(), "TransactionGetAllHandler")
	defer sp.End()
	Transactions, err := h.UseCase.GetAllTransaction(ctx)
	if err != nil {
		logs.LogError("Failed to get all transactions", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	logs.LogInfo("Get all transactions successfully", nil)
	var ResTransaction []*payload.RespondTransaction
	for _, transaction := range Transactions {
		ResTransaction = append(ResTransaction, payload.TransactionToResTransaction(transaction))
	}
	h.SetTransactionSubAttributes(ResTransaction, sp)
	return c.Status(fiber.StatusOK).JSON(ResTransaction)
}

func (h *TransactionHandler) SetTransactionSubAttributes(obj any, sp trace.Span) {
	if transactions, ok := obj.([]*payload.RespondTransaction); ok {
		transactionID := make([]string, len(transactions))
		transactionOrderAddress := make([]string, len(transactions))
		transactionTotalPrice := make([]float64, len(transactions))

		for _, transaction := range transactions {
			transactionID = append(transactionID, transaction.TransactionID.String())
			transactionOrderAddress = append(transactionOrderAddress, transaction.Address)
			transactionTotalPrice = append(transactionTotalPrice, transaction.TotalPrice)
		}
		sp.SetAttributes(
			attribute.StringSlice("TransactionID", transactionID),
			attribute.StringSlice("TransactionOrderAddress", transactionOrderAddress),
			attribute.Float64Slice("TransactionTotalPrice", transactionTotalPrice),
		)
	} else if transaction, ok := obj.(*payload.RespondTransaction); ok {
		sp.SetAttributes(
			attribute.String("TransactionID", transaction.TransactionID.String()),
			attribute.String("TransactionOrderAddress", transaction.Address),
			attribute.Float64("TransactionTotalPrice", transaction.TotalPrice),
		)
	} else {
		sp.RecordError(errors.New("invalid type" + reflect.TypeOf(obj).String()))
	}
}
