package auth

import (
	"app/internal/router/utils"
	"app/internal/service/auth"

	"github.com/gofiber/fiber/v2"
)

// Clear tokens
func Logout(c *fiber.Ctx) error {
	// Route to service
	identity := c.Locals("identity").(string)
	err := auth.Logout(&auth.LogoutParams{
		Refresh:  c.Cookies(CookieNameRefreshToken),
		Identity: identity,
	})

	// Clear cookies
	c.ClearCookie(CookieNameAccessToken, CookieNameRefreshToken)

	// Some error in service
	if err != nil {
		return c.JSON(utils.ErrorResponse(err.Error()))
	}

	// Success
	return c.JSON(utils.SuccessResponse(LoginResponse{
		ID:       0,
		Username: "sex",
	}))
}
