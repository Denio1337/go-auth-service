package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Save "identity" (user agent + IP) in context's locals
func Identified() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get user agent and ip
		userAgent := c.Get("User-Agent")
		ip := c.IP()

		// Build identity string
		identity := fmt.Sprintf("%s|%s", userAgent, ip)

		// Save identity string to locals
		c.Locals("identity", identity)
		c.Locals("userAgent", userAgent)
		c.Locals("ip", ip)

		return c.Next()
	}
}
