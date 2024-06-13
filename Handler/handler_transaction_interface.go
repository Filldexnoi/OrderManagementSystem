package Handler

import "github.com/gofiber/fiber/v2"

type TransactionHandlerI interface {
	CreateTransaction(c *fiber.Ctx) error
}
