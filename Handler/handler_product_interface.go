package Handler

import "github.com/gofiber/fiber/v2"

type ProductHandlerI interface {
	CreateProduct(c *fiber.Ctx) error
	GetProductByID(c *fiber.Ctx) error
	GetAllProducts(c *fiber.Ctx) error
	UpdateProduct(c *fiber.Ctx) error
	DeleteProduct(c *fiber.Ctx) error
}
