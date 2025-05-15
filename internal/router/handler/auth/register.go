package auth

import (
	cerror "app/internal/router/types/error"
	"app/internal/router/types/response"
	"app/internal/router/validator"
	"app/internal/service/auth"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	// Parse body
	dto := new(RegisterDTO)
	if err := c.BodyParser(dto); err != nil {
		cerr := *cerror.ErrInvalidInput
		cerr.Message = "can not parse query body"
		return &cerr
	}

	// Validate body
	if errs := validator.Validate(dto); len(errs) > 0 {
		return cerror.ValidationError(errs)
	}

	// Route to service
	result, err := auth.Register(&auth.RegisterParams{
		Username: dto.Username,
		Password: dto.Password,
	})

	// Handle error from service
	if err != nil {
		var cerr fiber.Error

		if errors.Is(err, auth.ErrUserExists) {
			cerr = *cerror.ErrConflict
		} else {
			cerr = *cerror.ErrInvalidInput
		}

		cerr.Message = err.Error()
		return &cerr
	}

	return c.JSON(response.SuccessResponse(RegisterResponse{
		ID:       result.ID,
		Username: result.Username,
	}))
}
