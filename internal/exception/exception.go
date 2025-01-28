package exception

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Status *fiber.Error `json:"status"`
	Errors any          `json:"errors"`
}

func (err *ErrorResponse) Error() string {
	return fmt.Sprintf("code %d status %s errors %+v", err.Status.Code, err.Status.Message, err.Errors)
}

func Validate(status *fiber.Error, err error) *ErrorResponse {
	var errors []map[string]string
	var message string

	for _, err := range err.(validator.ValidationErrors) {
		field := err.Field()

		switch err.Tag() {
		case "required":
			message = fmt.Sprintf("%s is %s", err.Field(), err.Tag())
		case "email":
			message = fmt.Sprintf("%s is invalid", err.Field())
		case "min":
			message = fmt.Sprintf("%s requires a minimum of %s characters", err.Field(), err.Param())
		case "max":
			message = fmt.Sprintf("%s requires a maximum of %s characters", err.Field(), err.Param())
		}

		errors = append(errors, map[string]string{field: message})
	}

	return &ErrorResponse{
		Status: status,
		Errors: errors,
	}
}

func Error(status *fiber.Error, err any) *ErrorResponse {
	return &ErrorResponse{
		Status: status,
		Errors: err,
	}
}
