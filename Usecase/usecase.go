package Usecase

type UseCase struct {
	Product     ProductUseCaseI
	Stock       StockUseCaseI
	Transaction TransactionUseCaseI
	Order       OrderUseCaseI
}
