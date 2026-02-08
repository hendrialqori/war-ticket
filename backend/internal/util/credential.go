package util

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hendrialqori/war-ticket/backend/internal/entity"
)

func GetCredential(c *fiber.Ctx) *entity.User {
	return c.Locals("credential").(*entity.User)
}
