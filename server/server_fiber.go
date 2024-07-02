package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
)

type FiberServer struct {
	app *fiber.App
}

func NewFiberServer() *FiberServer {
	app := fiber.New()
	return &FiberServer{app: app}
}

func (s *FiberServer) Start(port string, logfile *os.File) error {
	s.app.Use(logger.New(logger.Config{
		Format:     "${time} ${method} ${path}\n",
		TimeFormat: "15:04:05",
		TimeZone:   "Asia/Bangkok",
		Output:     logfile,
	}))
	s.app.Static("/", "../static", fiber.Static{
		Compress: true,
	})

	s.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	return s.app.Listen(fmt.Sprintf(":%s", port))
}

func SetUpLogger() *os.File {
	file, err := os.OpenFile("./temp/log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	return file
}
