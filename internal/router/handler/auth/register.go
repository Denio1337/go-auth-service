package auth

import (
	"app/internal/router/utils"
	"app/internal/router/validator"
	"app/internal/service/auth"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	//Parsing body
	dto := new(RegisterDTO)
	if err := c.BodyParser(dto); err != nil {
		return err
	}

	// Query parameters validation
	if errs := validator.Validate(dto); len(errs) > 0 {
		return utils.ValidationError(errs)
	}

	// Route to service
	result, err := auth.Register(&auth.RegisterParams{
		Username: dto.Username,
		Password: dto.Password,
	})

	// Some error in service
	if err != nil {
		return c.JSON(utils.ErrorResponse(err.Error()))
	}

	// Success
	return c.JSON(utils.SuccessResponse(RegisterResponse{
		ID:       result.ID,
		Username: result.Username,
	}))
}
