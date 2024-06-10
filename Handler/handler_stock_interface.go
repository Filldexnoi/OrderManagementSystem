package Handler

import "github.com/gofiber/fiber/v2"

type StockHandlerI interface {
	CreateStock(c *fiber.Ctx) error
	GetQtyProductByID(c *fiber.Ctx) error
	GetAllQtyProducts(c *fiber.Ctx) error
	UpdateStock(c *fiber.Ctx) error
	DeleteStock(c *fiber.Ctx) error
}
