package error

import (
	"app/internal/router/validator"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// General router errors
var (
	// 400 Bad Request
	ErrInvalidInput = fiber.NewError(fiber.StatusBadRequest, "invalid input data")

	// 401 Unauthorized
	ErrUnauthorized = fiber.NewError(fiber.StatusUnauthorized, "unauthorized access")

	// 403 Forbidden
	ErrForbidden = fiber.NewError(fiber.StatusForbidden, "access forbidden")

	// 404 Not Found
	ErrNotFound = fiber.NewError(fiber.StatusNotFound, "resource not found")

	// 409 Conflict
	ErrConflict = fiber.NewError(fiber.StatusConflict, "resource conflict")

	// 500 Internal Server Error
	ErrInternalServer = fiber.NewError(fiber.StatusInternalServerError, "internal server error")
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
