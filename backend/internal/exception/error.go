package exception

import "github.com/gofiber/fiber/v2"

var (
	// error user domain
	ErrUserUnauthorized = fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	//error
	ErrDataAlreadyExists   = fiber.NewError(fiber.StatusBadRequest, "Data already exists")
	ErrDataNotFound        = fiber.NewError(fiber.StatusNotFound, "Data not found")
	ErrInternalServerError = fiber.NewError(fiber.StatusInternalServerError, "Internal server error")
)
