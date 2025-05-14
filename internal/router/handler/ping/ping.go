package ping

import (
	"github.com/gofiber/fiber/v2"
)

func Ping(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "I'm alive",
		"data":    nil,
	})
}
