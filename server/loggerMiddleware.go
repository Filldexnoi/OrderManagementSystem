package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"time"
)

func LoggerMiddleware(next fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		fields := logrus.Fields{
			"method":   c.Method(),
			"path":     c.Path(),
			"query":    c.OriginalURL(),
			"remoteIP": c.IP(),
		}
		err := next(c)
		if err != nil {
			return err
		}

		fields["status"] = c.Response().StatusCode()
		fields["latency"] = time.Since(start).Seconds()

		logrus.WithFields(fields).Info("HTTP response")
		return nil
	}
}
