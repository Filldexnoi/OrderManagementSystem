package Handler

import (
	"awesomeProject/Usecase"
	"awesomeProject/adpter"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	UseCase Usecase.ProductUseCaseI
}

func NewProductHandler(UseCase Usecase.ProductUseCaseI) ProductHandlerI {
	return &ProductHandler{UseCase: UseCase}
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	product := new(adpter.ProductBrowserInput)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err := h.UseCase.CreateProduct(product.ToProduct())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(product)
}

func (h *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	products, err := h.UseCase.GetAllProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(products)
}

func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}
	product, err := h.UseCase.GetByIDProduct(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(adpter.ToProduct(product))
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}
	product := new(adpter.ProductBrowserInput)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err = h.UseCase.UpdateProduct(product.ToProduct(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(product)
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}
	err = h.UseCase.DeleteProduct(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
