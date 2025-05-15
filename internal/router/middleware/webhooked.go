package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// Save webhook URL in locals
func Webhooked(webhookURL string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Save webhook URL in locals
		c.Locals("webhook", webhookURL)

		return c.Next()
	}
}
