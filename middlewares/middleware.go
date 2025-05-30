package middlewares

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func PathLogger(c *fiber.Ctx) error {
	log.WithFields(log.Fields{
		"method":     c.Method(),
		"path":       c.Path(),
		"IP Address": c.IP(),
	}).Info("Request received!")
	return c.Next()
}
