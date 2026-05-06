package utils

import (
	"errors"
	"reflect"
	"strings"

	"github.com/AgufSamudra/subscription/src/internal/apperror"
	"github.com/go-playground/validator/v10"
)

var validate = newValidator()

func newValidator() *validator.Validate {
	v := validator.New()
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return v
}

func ValidateStruct(payload any) error {
	if err := validate.Struct(payload); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) && len(validationErrors) > 0 {
			return apperror.BadRequestError(validationMessage(validationErrors[0]), err)
		}

		return apperror.BadRequestError("request validation failed", err)
	}

	return nil
}

func validationMessage(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return fieldError.Field() + " is required"
	case "email":
		return fieldError.Field() + " must be a valid email"
	default:
		return fieldError.Field() + " is invalid"
	}
}
