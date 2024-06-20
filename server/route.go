package server

import (
	"awesomeProject/Handler"
	"awesomeProject/Usecase"
)

func (s *FiberServer) SetupFiberRoute(UseCase *Usecase.UseCase) {

	productHandler := Handler.NewProductHandler(UseCase.Product)
	s.app.Post("/product", productHandler.CreateProduct)
	s.app.Get("/products", productHandler.GetAllProducts)
	s.app.Get("/product/:id", productHandler.GetProductByID)
	s.app.Put("/product/:id", productHandler.UpdateProduct)
	s.app.Delete("/product/:id", productHandler.DeleteProduct)

	stockHandler := Handler.NewStockHandler(UseCase.Stock)
	s.app.Post("/stock", stockHandler.CreateStock)
	s.app.Get("/stocks", stockHandler.GetAllQtyProducts)
	s.app.Get("/stock/:id", stockHandler.GetQtyProductByID)
	s.app.Put("/stock/:id", stockHandler.UpdateStock)
	s.app.Delete("/stock/:id", stockHandler.DeleteStock)

	TransactionHandler := Handler.NewTransactionHandler(UseCase.Transaction)
	s.app.Post("/order/calculate", TransactionHandler.CreateTransaction)
	s.app.Get("/order/calculate", TransactionHandler.GetAllTransactions)

	OrderHandler := Handler.NewOrderHandler(UseCase.Order)
	s.app.Post("/order", OrderHandler.CreateOrder)
	s.app.Patch("/order/status/:id", OrderHandler.UpdateOrderStatus)
	s.app.Get("/orders", OrderHandler.GetAllOrders)
}
