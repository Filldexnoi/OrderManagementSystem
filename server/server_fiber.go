package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	app *fiber.App
}

func NewFiberServer() *FiberServer {
	app := fiber.New()
	return &FiberServer{app: app}
}

func (s *FiberServer) Start(port string) error {
	return s.app.Listen(fmt.Sprintf(":%s", port))

}
