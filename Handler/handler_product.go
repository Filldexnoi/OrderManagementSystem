package Handler

import (
	"awesomeProject/Usecase"
	"awesomeProject/payload"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	UseCase Usecase.ProductUseCaseI
}

func NewProductHandler(UseCase Usecase.ProductUseCaseI) ProductHandlerI {
	return &ProductHandler{UseCase: UseCase}
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	product := new(payload.RequestProduct)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	productEntity, err := h.UseCase.CreateProduct(product.ToProduct())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	ResProduct := payload.ProductToRespondProduct(productEntity)
	return c.Status(fiber.StatusCreated).JSON(ResProduct)
}

func (h *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	products, err := h.UseCase.GetAllProducts()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	var ResProduct []*payload.RespondProduct
	for _, product := range products {
		ResProduct = append(ResProduct, payload.ProductToRespondProduct(product))
	}
	return c.Status(fiber.StatusOK).JSON(ResProduct)
}

func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}
	product, err := h.UseCase.GetByIDProduct(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	ResProduct := payload.ProductToRespondProduct(product)
	return c.Status(fiber.StatusOK).JSON(ResProduct)
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}
	product := new(payload.RequestProduct)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	productEntity, err := h.UseCase.UpdateProduct(product.ToProduct(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	ResProduct := payload.ProductToRespondProduct(productEntity)
	return c.Status(fiber.StatusOK).JSON(ResProduct)
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}
	productEntity, err := h.UseCase.DeleteProduct(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	ResProduct := payload.ProductToRespondProduct(productEntity)
	return c.Status(fiber.StatusOK).JSON(ResProduct)
}
