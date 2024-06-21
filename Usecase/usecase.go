package Usecase

import "awesomeProject/Repo"

type UseCase struct {
	Product     ProductUseCaseI
	Stock       StockUseCaseI
	Transaction TransactionUseCaseI
	Order       OrderUseCaseI
}

func NewUseCase(Repo *Repo.GormRepo) *UseCase {
	return &UseCase{
		Product:     NewProductUseCase(Repo.ProductRepo),
		Stock:       NewStockUseCase(Repo.StockRepo),
		Transaction: NewTransactionUseCase(Repo.TransactionRepo, Repo.ProductRepo),
		Order:       NewOrderUseCase(Repo.OrderRepo, Repo.StockRepo, Repo.TransactionRepo),
	}
}
