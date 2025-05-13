package hello

import (
	"app/internal/service/hello"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := hello.Hello()
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": fmt.Sprintf("%v", err),
				"data":    nil,
			})
		}

		return c.JSON(fiber.Map{
			"status":  "success",
			"message": result,
			"data":    nil,
		})
	}
}
