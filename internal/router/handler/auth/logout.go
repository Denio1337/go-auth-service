package auth

import (
	cerror "app/internal/router/types/error"
	"app/internal/router/types/response"
	"app/internal/service/auth"

	"github.com/gofiber/fiber/v2"
)

// DELETE /auth/logout: clear tokens
func Logout(c *fiber.Ctx) error {
	// Get identity
	identity, ok := c.Locals("identity").(string)
	if !ok {
		cerr := *cerror.ErrUnauthorized
		cerr.Message = "can not identify user"
		return &cerr
	}

	// Get refresh token from cookie
	refreshToken := c.Cookies(CookieNameRefreshToken)
	if refreshToken == "" {
		cerr := *cerror.ErrUnauthorized
		cerr.Message = "invalid refresh token"
		return &cerr
	}

	// Route to service
	err := auth.Logout(&auth.LogoutParams{
		Refresh:  c.Cookies(CookieNameRefreshToken),
		Identity: identity,
	})

	// Clear cookies
	setAuthCookies(c, nil)

	// Handle error from service
	if err != nil {
		cerr := *cerror.ErrForbidden
		cerr.Message = err.Error()
		return &cerr
	}

	return c.JSON(response.SuccessResponse(struct{}{}))
}
