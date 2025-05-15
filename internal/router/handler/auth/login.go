package auth

import (
	cerror "app/internal/router/types/error"
	"app/internal/router/types/response"
	"app/internal/router/validator"
	"app/internal/service/auth"
	"time"

	"github.com/gofiber/fiber/v2"
)

// GET /auth/login: get access and refresh tokens
func Login(c *fiber.Ctx) error {
	// Parse query parameters
	dto := &LoginDTO{
		Username: c.Query("username"),
		Password: c.Query("password"),
	}

	// Validate query parameters
	if errs := validator.Validate(dto); len(errs) > 0 {
		return cerror.ValidationError(errs)
	}

	// Get identity
	identity, ok := c.Locals("identity").(string)
	if !ok {
		cerr := *cerror.ErrUnauthorized
		cerr.Message = "can not identify user"
		return &cerr
	}
	userAgent := c.Locals("userAgent").(string)
	ip := c.Locals("ip").(string)

	// Route to service
	result, err := auth.Login(&auth.LoginParams{
		Username:  dto.Username,
		Password:  dto.Password,
		Identity:  identity,
		IP:        ip,
		UserAgent: userAgent,
	})

	// Handle error from service
	if err != nil {
		cerr := *cerror.ErrUnauthorized
		cerr.Message = err.Error()
		return &cerr
	}

	// Set auth cookies
	setAuthCookies(c, result.Tokens)

	return c.JSON(response.SuccessResponse(LoginResponse{
		ID:       result.ID,
		Username: result.Username,
	}))
}

// Set or clear auth cookies based on the 'clear' flag
func setAuthCookies(c *fiber.Ctx, tokens *auth.Tokens) {
	if tokens == nil {
		// Clear cookies by setting past time
		clearTime := time.Unix(0, 0)

		setCookie(c, &CookieParams{
			Name:    CookieNameAccessToken,
			Value:   "",
			Expires: clearTime,
		})

		setCookie(c, &CookieParams{
			Name:    CookieNameRefreshToken,
			Value:   "",
			Expires: clearTime,
		})
	} else {
		// Set cookies normally
		setCookie(c, &CookieParams{
			Name:    CookieNameAccessToken,
			Value:   tokens.Access,
			Expires: time.Now().Add(auth.TokenLiveTimeAccess),
		})

		setCookie(c, &CookieParams{
			Name:    CookieNameRefreshToken,
			Value:   tokens.Refresh,
			Expires: time.Now().Add(auth.TokenLiveTimeRefresh),
		})
	}
}

// Set cookie
func setCookie(c *fiber.Ctx, params *CookieParams) {
	c.Cookie(&fiber.Cookie{
		Name:     params.Name,
		Value:    params.Value,
		Path:     "/api",
		Domain:   "localhost",
		Expires:  params.Expires,
		MaxAge:   int(time.Until(params.Expires).Seconds()),
		HTTPOnly: true,
		Secure:   false,
		SameSite: fiber.CookieSameSiteStrictMode,
	})
}
