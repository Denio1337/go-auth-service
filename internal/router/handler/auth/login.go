package auth

import (
	"app/internal/router/utils"
	"app/internal/router/validator"
	"app/internal/service/auth"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	// Parsing query parameters
	dto := &LoginDTO{
		Username: c.Query("username"),
		Password: c.Query("password"),
	}

	// Query parameters validation
	if errs := validator.Validate(dto); len(errs) > 0 {
		return utils.ValidationError(errs)
	}

	// Route to service
	identity := c.Locals("identity").(string)
	result, err := auth.Login(&auth.LoginParams{
		Username: dto.Username,
		Password: dto.Password,
		Identity: identity,
	})

	// Some error in service
	if err != nil {
		return c.JSON(utils.ErrorResponse(err.Error()))
	}

	// Setting auth cookies
	setAuthCookies(c, result)

	// Success
	return c.JSON(utils.SuccessResponse(LoginResponse{
		ID:       result.ID,
		Username: result.Username,
	}))
}

// Set auth cookies with access and refresh tokens
func setAuthCookies(c *fiber.Ctx, result *auth.LoginResult) {
	// Setting access token cookie
	setCookie(c, &CookieParams{
		Name:    CookieNameAccessToken,
		Value:   result.Access,
		Expires: time.Now().Add(auth.TokenLiveTimeAccess),
	})

	// Setting refresh token cookie
	setCookie(c, &CookieParams{
		Name:    CookieNameRefreshToken,
		Value:   result.Refresh,
		Expires: time.Now().Add(auth.TokenLiveTimeRefresh),
	})
}

// General wrapper to set cookie
func setCookie(c *fiber.Ctx, params *CookieParams) {
	c.Cookie(&fiber.Cookie{
		Name:     params.Name,
		Value:    params.Value,
		Path:     "/api",
		Domain:   "localhost",
		Expires:  params.Expires,
		HTTPOnly: true,
		Secure:   false,
		SameSite: fiber.CookieSameSiteStrictMode,
	})
}
