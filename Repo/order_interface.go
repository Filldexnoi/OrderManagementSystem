package Repo

import "awesomeProject/entities"

type OrderRepoI interface {
	SaveCreateOrder(order entities.Order) error
}
