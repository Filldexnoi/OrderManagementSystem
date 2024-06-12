package server

import (
	"awesomeProject/Handler"
	"awesomeProject/Usecase"
)

func (s *FiberServer) SetupFiberRoute(productUseCase Usecase.ProductUseCaseI, stockUseCase Usecase.StockUseCaseI) {
	productHandler := Handler.NewProductHandler(productUseCase)

	s.app.Post("/product", productHandler.CreateProduct)
	s.app.Get("/products", productHandler.GetAllProducts)
	s.app.Get("/product/:id", productHandler.GetProductByID)
	s.app.Put("/product/:id", productHandler.UpdateProduct)
	s.app.Delete("/product/:id", productHandler.DeleteProduct)

	stockHandler := Handler.NewStockHandler(stockUseCase)

	s.app.Post("/stock", stockHandler.CreateStock)
	s.app.Get("/stocks", stockHandler.GetAllQtyProducts)
	s.app.Get("/stock/:id", stockHandler.GetQtyProductByID)
	s.app.Put("/stock/:id", stockHandler.UpdateStock)
	s.app.Delete("/stock/:id", stockHandler.DeleteStock)

}
