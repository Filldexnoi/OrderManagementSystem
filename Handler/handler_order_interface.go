package Handler

import "github.com/gofiber/fiber/v2"

type OrderHandlerI interface {
	CreateOrder(c *fiber.Ctx) error
}
