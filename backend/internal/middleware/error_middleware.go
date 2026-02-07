package middleware

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/hendrialqori/war-ticket/backend/internal/exception"
)

func NewErrorMiddleware() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		var e any
		var appError *exception.AppError
		var valError validator.ValidationErrors

		code := 500
		message := "Error occured!"

		// mapping app error
		if errors.As(err, &appError) {
			code = appError.Code
			message = appError.Message
			e = appError.Err
		}

		// mapping validation error
		if errors.As(err, &valError) {
			code = 400
			message = "Validation error"

			fields := make([]map[string]any, 0)
			for _, fe := range valError {

				fields = append(fields,
					map[string]any{
						"field":      strings.ToLower(fe.Field()),
						"constraint": fe.Tag(),
					},
				)
			}
			e = fields
		}

		errMap := &exception.AppError{
			Success: false,
			Code:    code,
			Message: message,
			Err:     e,
		}

		return c.Status(code).JSON(errMap)
	}

}
