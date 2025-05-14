package auth

import (
	"app/internal/router/utils"
	"app/internal/service/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Me(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)

	// Route to service
	result, err := auth.Me(&auth.MeParams{
		ID: uint(id),
	})

	// Some error in service
	if err != nil {
		return c.JSON(utils.ErrorResponse(err.Error()))
	}

	// Success
	return c.JSON(utils.SuccessResponse(LoginResponse{
		ID:       result.ID,
		Username: result.Username,
	}))
}
