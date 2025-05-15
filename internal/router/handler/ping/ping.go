package ping

import (
	"app/internal/router/types/response"

	"github.com/gofiber/fiber/v2"
)

// GET: check API's availability
func Ping(c *fiber.Ctx) error {
	return c.JSON(response.SuccessResponse(nil))
}
