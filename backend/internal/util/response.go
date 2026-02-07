package util

import "github.com/gofiber/fiber/v2"

type APIResponse struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
	Message any  `json:"message,omitempty"`
	Error   any  `json:"error,omitempty"`
}

func MapToResponse(c *fiber.Ctx, status int, data any, message any) error {
	return c.Status(status).JSON(APIResponse{
		Success: true,
		Data:    data,
		Message: message,
	})
}
