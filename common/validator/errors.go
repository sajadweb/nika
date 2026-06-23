package validator

import (
	"github.com/go-playground/validator/v10"
)

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// FormatErrors converts validator.ValidationErrors into a slice of FieldError.
func FormatErrors(err error) []FieldError {
	var result []FieldError

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return result
	}

	for _, e := range validationErrors {
		result = append(result, FieldError{
			Field:   e.Field(),
			Message: messageForTag(e),
		})
	}

	return result
}

func messageForTag(e validator.FieldError) string {
	switch e.Tag() {
	case "ir_mobile":
		return "Mobile number is not valid"
	case "objectid":
		return "ObjectId not valid"
	default:
		return e.Error()
	}
}
