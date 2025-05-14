package validator

// Validation error structure
type ValidationError struct {
	FailedField string
	Tag         string
	Value       any
}
