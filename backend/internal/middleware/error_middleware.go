package middleware

import "github.com/gofiber/fiber/v2"

type ErrorResponse struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func NewErrorMiddleware() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		var (
			success    = false
			statusCode = fiber.StatusInternalServerError
			message    = "Internal server error"
		)

		// api error
		if e, ok := err.(*fiber.Error); ok {
			success = false
			statusCode = e.Code
			message = e.Message
		}

		return c.Status(statusCode).JSON(ErrorResponse{
			Success:    success,
			StatusCode: statusCode,
			Message:    message,
		})
	}

}
