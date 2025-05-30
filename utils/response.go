package utils

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type ResponseWithToken struct {
	Status      string      `json:"status"`
	AccessToken string      `json:"access_token"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data,omitempty"`
	Success     bool        `json:"success"`
}

func JsonResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	log.WithFields(log.Fields{
		"status":  statusCode,
		"message": message,
		"data":    data,
		"success": statusCode >= 200 && statusCode < 300,
	}).Info("Sending Response")

	if c.Locals("access_token") != nil && c.Locals("access_token") != "" {
		return c.Status(statusCode).JSON(ResponseWithToken{
			Status:      "success",
			AccessToken: c.Locals("access_token").(string),
			Message:     message,
			Data:        data,
			Success:     statusCode >= 200 && statusCode < 300,
		})
	}
	return c.Status(statusCode).JSON(fiber.Map{
		"status":  "success",
		"message": message,
		"data":    data,
		"success": statusCode >= 200 && statusCode < 300,
	})
}

func JsonErrorResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	log.WithFields(log.Fields{
		"status":  statusCode,
		"message": message,
		"data":    data,
		"success": statusCode >= 200 && statusCode < 300,
	}).Info("Sending Response")
	if c.Locals("access_token") != nil && c.Locals("access_token") != "" {
		return c.Status(statusCode).JSON(ResponseWithToken{
			Status:      "error",
			AccessToken: c.Locals("access_token").(string),
			Message:     message,
			Success:     false,
			Data:        data,
		})
	}

	return c.Status(statusCode).JSON(ResponseWithToken{
		Status:  "error",
		Message: message,
		Success: false,
		Data:    data,
	})
}
