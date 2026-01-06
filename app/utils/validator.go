package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// CustomValidator wraps the validator package
type CustomValidator struct {
	validator *validator.Validate
}

// Validate method for Echo
func (cv *CustomValidator) Validate(i any) error {
	return cv.validator.Struct(i)
}

// NewValidator initializes a new validator
func NewValidator() *CustomValidator {
	validate := validator.New()

	validate.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
		phone := fl.Field().String()

		if phone == "" {
			return true
		}

		pattern := regexp.MustCompile(`^(\+60|0)1[0-9]{8,9}$`)
		return pattern.MatchString(phone)
	})

	validate.RegisterValidation("email", func(fl validator.FieldLevel) bool {
		email := fl.Field().String()

		if email == "" {
			return true
		}

		pattern := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
		return pattern.MatchString(email)
	})

	return &CustomValidator{validator: validate}
}
