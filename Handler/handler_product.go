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

type ProductHandler struct {
	UseCase Usecase.ProductUseCaseI
}

func NewProductHandler(UseCase Usecase.ProductUseCaseI) ProductHandlerI {
	return &ProductHandler{UseCase: UseCase}
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	ctx, sp := otel.Tracer("product").Start(c.Context(), "productCreateHandler")
	defer sp.End()

	product := new(payload.RequestProduct)
	if err := c.BodyParser(product); err != nil {
		logs.LogError("Failed to bodyParser product request", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	productEntity, err := h.UseCase.CreateProduct(ctx, product.ToProduct())
	if err != nil {
		logs.LogError("Failed to create product", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	logs.LogInfo("Product created successfully", logrus.Fields{"product_id": productEntity.ProductId})
	ResProduct := payload.ProductToRespondProduct(productEntity)
	h.SetSubAttributesWithJson(ResProduct, sp)
	return c.Status(fiber.StatusCreated).JSON(ResProduct)
}

func (h *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	ctx, sp := otel.Tracer("product").Start(c.Context(), "productGetAllHandler")
	defer sp.End()

	products, err := h.UseCase.GetAllProducts(ctx)
	if err != nil {
		logs.LogError("Failed to get all products", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	logs.LogInfo("Get all product successfully", nil)
	var ResProduct []*payload.RespondProduct
	for _, product := range products {
		ResProduct = append(ResProduct, payload.ProductToRespondProduct(product))
	}
	h.SetSubAttributesWithJson(ResProduct, sp)
	return c.Status(fiber.StatusOK).JSON(ResProduct)
}

func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	ctx, sp := otel.Tracer("product").Start(c.Context(), "productGetByIDHandler")
	defer sp.End()
	id, err := c.ParamsInt("id")
	if err != nil {
		logs.LogError("Failed to params int product id", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}
	product, err := h.UseCase.GetByIDProduct(ctx, uint(id))
	if err != nil {
		logs.LogError("Failed to get product by id", logrus.Fields{"error": err.Error(), "product_id": id})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	logs.LogInfo("Found product by ID", logrus.Fields{"product_id": product.ProductId})
	ResProduct := payload.ProductToRespondProduct(product)
	h.SetSubAttributesWithJson(ResProduct, sp)
	return c.Status(fiber.StatusOK).JSON(ResProduct)
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	ctx, sp := otel.Tracer("product").Start(c.Context(), "productUpdateHandler")
	defer sp.End()
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
	productEntity, err := h.UseCase.UpdateProduct(ctx, product.ToProduct(), uint(id))
	if err != nil {
		logs.LogError("Failed to update product", logrus.Fields{"error": err.Error(), "product_id": id})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	logs.LogInfo("update product successfully", logrus.Fields{"product_id": productEntity.ProductId})
	ResProduct := payload.ProductToRespondProduct(productEntity)
	h.SetSubAttributesWithJson(ResProduct, sp)
	return c.Status(fiber.StatusOK).JSON(ResProduct)
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	ctx, sp := otel.Tracer("product").Start(c.Context(), "productDeleteHandler")
	defer sp.End()
	id, err := c.ParamsInt("id")
	if err != nil {
		logs.LogError("Failed to params int product id", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}
	productEntity, err := h.UseCase.DeleteProduct(ctx, uint(id))
	if err != nil {
		logs.LogError("Failed to delete product", logrus.Fields{"error": err.Error(), "product_id": id})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	logs.LogInfo("Delete product successfully", logrus.Fields{"product_id": productEntity.ProductId})
	ResProduct := payload.ProductToRespondProduct(productEntity)
	h.SetSubAttributesWithJson(ResProduct, sp)
	return c.Status(fiber.StatusOK).JSON(ResProduct)
}

func (h *ProductHandler) SetSubAttributesWithJson(obj any, sp trace.Span) {
	if products, ok := obj.([]*payload.RespondProduct); ok {
		var productIDs []int
		var productNames []string
		var productTypes []string
		var productPrices []float64

		for _, product := range products {
			productIDs = append(productIDs, int(product.ProductId))
			productNames = append(productNames, product.ProductName)
			productTypes = append(productTypes, product.ProductTypes)
			productPrices = append(productPrices, product.ProductPrice)
		}

		sp.SetAttributes(
			attribute.IntSlice("ProductID", productIDs),
			attribute.StringSlice("ProductName", productNames),
			attribute.StringSlice("ProductType", productTypes),
			attribute.Float64Slice("ProductPrice", productPrices),
		)
	} else if product, ok := obj.(*payload.RespondProduct); ok {
		sp.SetAttributes(
			attribute.Int("ProductID", int(product.ProductId)),
			attribute.String("ProductName", product.ProductName),
			attribute.String("ProductType", product.ProductTypes),
			attribute.Float64("ProductPrice", product.ProductPrice),
		)
	} else {
		sp.RecordError(errors.New("invalid type" + reflect.TypeOf(obj).String()))
	}
}
