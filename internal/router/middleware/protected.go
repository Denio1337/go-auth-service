package middleware

import (
	"app/internal/config"
	cerror "app/internal/router/types/error"
	"fmt"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

// Check if token valid
func Protected(cookieName string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     jwtware.SigningKey{Key: []byte(config.Get(config.EnvSecret))},
		ErrorHandler:   jwtError,
		TokenLookup:    fmt.Sprintf("cookie:%s", cookieName),
		SuccessHandler: checkIdentity,
	})
}

// Compare token and query identity
func checkIdentity(c *fiber.Ctx) error {
	// // Get raw identity string
	// rawIdentity, ok := c.Locals("identity").(string)
	// if !ok {
	// 	return &fiber.Error{
	// 		Code:    fiber.ErrUnauthorized.Code,
	// 		Message: "invalid local identity",
	// 	}
	// }

	// // Get user local after authorization
	// user, ok := c.Locals("user").(*jwt.Token)
	// if !ok {
	// 	return &fiber.Error{
	// 		Code:    fiber.ErrUnauthorized.Code,
	// 		Message: "invalid or expired auth token",
	// 	}
	// }

	// // Get token identity
	// claims := user.Claims.(jwt.MapClaims)
	// tokenIdentity, ok := claims["identity"].(string)
	// if !ok {
	// 	return &fiber.Error{
	// 		Code:    fiber.ErrUnauthorized.Code,
	// 		Message: "can't get identity from token",
	// 	}
	// }

	// // Check if token identity equals local identity
	// isValid := utils.CompareWithHash(rawIdentity, tokenIdentity)
	// if !isValid {
	// 	return &fiber.Error{
	// 		Code:    fiber.ErrUnauthorized.Code,
	// 		Message: "invalid identity",
	// 	}
	// }

	return c.Next()
}

// JWT error handler
func jwtError(c *fiber.Ctx, err error) error {
	cerr := *cerror.ErrUnauthorized
	cerr.Message = "invalid or expired auth token"
	return &cerr
}
