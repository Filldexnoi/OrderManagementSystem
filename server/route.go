package server

import (
	"awesomeProject/Handler"
	"awesomeProject/Usecase"
)

func (s *FiberServer) SetupFiberRoute(productUseCase Usecase.UseCase) {
	productHandler := Handler.NewProductHandler(productUseCase)

	s.app.Post("/products", productHandler.CreateProduct)
	s.app.Get("/products", productHandler.GetAllProducts)
	s.app.Get("/products/:id", productHandler.GetProductByID)
	s.app.Put("/products/:id", productHandler.UpdateProduct)
	s.app.Delete("/products/:id", productHandler.DeleteProduct)
}
