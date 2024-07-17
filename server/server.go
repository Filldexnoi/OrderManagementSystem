package server

import "awesomeProject/Usecase"

type Server interface {
	Start(address string, u *Usecase.UseCase) error
}
