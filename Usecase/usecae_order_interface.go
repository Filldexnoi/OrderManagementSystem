package Usecase

import "awesomeProject/entities"

type OrderUseCaseI interface {
	CreateOrder(order entities.Order) error
}
