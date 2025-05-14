package utils

import (
	"app/internal/router/types"
	"app/internal/router/validator"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Create fiber error about invalid validation
func ValidationError(errs []validator.ValidationError) *fiber.Error {
	errMsgs := make([]string, 0)

	// Formatting list of validation errors
	for _, err := range errs {
		errMsgs = append(errMsgs, fmt.Sprintf(
			"[%s]: '%v' | Needs to implement '%s'",
			err.FailedField,
			err.Value,
			err.Tag,
		))
	}

	return &fiber.Error{
		Code:    fiber.ErrBadRequest.Code,
		Message: strings.Join(errMsgs, " and "),
	}
}

func SuccessResponse(data any) *types.Response {
	return &types.Response{
		Success: true,
		Message: "",
		Data:    data,
	}
}

func ErrorResponse(message string) *types.Response {
	return &types.Response{
		Success: false,
		Message: message,
		Data:    nil,
	}
}
