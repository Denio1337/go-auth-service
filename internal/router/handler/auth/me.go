package auth

import (
	cerror "app/internal/router/types/error"
	"app/internal/router/types/response"
	"app/internal/service/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// GET /auth/me: get user profile
func Me(c *fiber.Ctx) error {
	// Get user info from cookie
	user, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		cerr := *cerror.ErrUnauthorized
		cerr.Message = "invalid access token"
		return &cerr
	}

	// Get user id
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)

	// Route to service
	result, err := auth.Me(&auth.MeParams{
		ID: uint(id),
	})

	// Handle error from service
	if err != nil {
		cerr := *cerror.ErrUnauthorized
		cerr.Message = err.Error()
		return &cerr
	}

	return c.JSON(response.SuccessResponse(MeResponse{
		ID:       result.ID,
		Username: result.Username,
	}))
}
