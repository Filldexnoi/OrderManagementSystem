package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"time"
)

func LoggerMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	fields := logrus.Fields{
		"method":   c.Method(),
		"path":     c.Path(),
		"query":    c.OriginalURL(),
		"remoteIP": c.IP(),
	}

	logrus.WithFields(fields).Info("HTTP request received")

	err := c.Next()
	if err != nil {
		return err
	}

	fields["status"] = c.Response().StatusCode()
	fields["latency"] = time.Since(start).Seconds()

	logrus.WithFields(fields).Info("HTTP response sent")
	return nil
}
