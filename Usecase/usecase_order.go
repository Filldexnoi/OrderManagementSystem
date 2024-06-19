package Usecase

import (
	"awesomeProject/Repo"
	"awesomeProject/entities"
)

type OrderUseCase struct {
	Repo Repo.OrderRepoI
}

func NewOrderUseCase(repo Repo.OrderRepoI) OrderUseCaseI {
	return &OrderUseCase{
		Repo: repo,
	}
}

func (u *OrderUseCase) CreateOrder(order entities.Order) error {
	return u.Repo.SaveCreateOrder(order)
}
