package middleware

import (
	"github.com/MatTwix/RE-minder-Bots/config"
	"github.com/gofiber/fiber/v3"
)

var cfg = config.LoadConfig()

func APIKeyMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		apiKey := c.Get("X-API-Key")
		if apiKey == "" || apiKey != cfg.InternalAPIKey {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or missing API key"})
		}

		return c.Next()
	}
}
