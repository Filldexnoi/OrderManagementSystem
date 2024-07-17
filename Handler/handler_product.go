package Handler

import (
	"awesomeProject/Usecase"
	"awesomeProject/observability/logs"
	"awesomeProject/payload"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
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
		logs.LogError("Failed to bodyParser product request", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	productEntity, err := h.UseCase.CreateProduct(product.ToProduct())
	if err != nil {
		logs.LogError("Failed to create product", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	logs.LogInfo("Product created successfully", logrus.Fields{"product_id": productEntity.ProductId})
	ResProduct := payload.ProductToRespondProduct(productEntity)
	return c.Status(fiber.StatusCreated).JSON(ResProduct)
}

func (h *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	products, err := h.UseCase.GetAllProducts()
	if err != nil {
		logs.LogError("Failed to get all products", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	logs.LogInfo("Get all product successfully", nil)
	var ResProduct []*payload.RespondProduct
	for _, product := range products {
		ResProduct = append(ResProduct, payload.ProductToRespondProduct(product))
	}
	return c.Status(fiber.StatusOK).JSON(ResProduct)
}

func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		logs.LogError("Failed to params int product id", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}
	product, err := h.UseCase.GetByIDProduct(uint(id))
	if err != nil {
		logs.LogError("Failed to get product by id", logrus.Fields{"error": err.Error(), "product_id": id})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	logs.LogInfo("Found product by ID", logrus.Fields{"product_id": product.ProductId})
	ResProduct := payload.ProductToRespondProduct(product)
	return c.Status(fiber.StatusOK).JSON(ResProduct)
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		logs.LogError("Failed to params int product id", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}
	product := new(payload.RequestProduct)
	if err := c.BodyParser(product); err != nil {
		logs.LogError("Failed to bodyParser product request", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	productEntity, err := h.UseCase.UpdateProduct(product.ToProduct(), uint(id))
	if err != nil {
		logs.LogError("Failed to update product", logrus.Fields{"error": err.Error(), "product_id": id})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	logs.LogInfo("update product successfully", logrus.Fields{"product_id": productEntity.ProductId})
	ResProduct := payload.ProductToRespondProduct(productEntity)
	return c.Status(fiber.StatusOK).JSON(ResProduct)
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		logs.LogError("Failed to params int product id", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}
	productEntity, err := h.UseCase.DeleteProduct(uint(id))
	if err != nil {
		logs.LogError("Failed to delete product", logrus.Fields{"error": err.Error(), "product_id": id})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	logs.LogInfo("Delete product successfully", logrus.Fields{"product_id": productEntity.ProductId})
	ResProduct := payload.ProductToRespondProduct(productEntity)
	return c.Status(fiber.StatusOK).JSON(ResProduct)
}
