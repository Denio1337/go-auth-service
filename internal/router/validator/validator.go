package validator

import "github.com/go-playground/validator/v10"

// Package instance of validator
var instance *validator.Validate

// Initialize instance
func init() {
	instance = validator.New()
}

func Validate(data any) []ValidationError {
	// Initialize an array of validation errors
	validationErrors := []ValidationError{}

	// Process and fill validation errors
	errs := instance.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			elem := ValidationError{
				FailedField: err.Field(),
				Tag:         err.Tag(),
				Value:       err.Value(),
			}

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}
