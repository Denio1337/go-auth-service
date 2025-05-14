package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Saves "identity" (user agent + IP) in locals
func Identified() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get user agent and ip
		userAgent := c.Get("User-Agent")
		ip := c.IP()

		// Raw user agent + ip string
		rawIdentity := fmt.Sprintf("%s|%s", userAgent, ip)

		// Save identity to locals
		c.Locals("identity", rawIdentity)

		return c.Next()
	}
}
