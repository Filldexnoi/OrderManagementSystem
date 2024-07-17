package server

import (
	"awesomeProject/Usecase"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	app *fiber.App
}

func NewFiberServer() FiberServer {
	app := fiber.New()
	app.Use(LoggerMiddleware)
	app.Use(TracingMiddleware)
	return FiberServer{app: app}
}

func (s *FiberServer) Start(port string, u *Usecase.UseCase) error {
	s.SetupFiberRoute(u)
	return s.app.Listen(fmt.Sprintf(":%s", port))
}
