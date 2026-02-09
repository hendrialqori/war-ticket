package middleware

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hendrialqori/war-ticket/backend/internal/config"
	"github.com/hendrialqori/war-ticket/backend/internal/entity"
	"github.com/hendrialqori/war-ticket/backend/internal/exception"
	"github.com/hendrialqori/war-ticket/backend/internal/util"
)

func NewTokenMiddleware(config *config.AppConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtToken := util.NewJwtToken(config.Secret)

		auth := c.Get("Authorization")

		if auth == "" {
			return exception.New(401, "Missing authorization header")
		}

		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			return exception.New(401, "Invalid authorization format")
		}

		tokenStr := parts[1]
		claims, err := jwtToken.Verify(tokenStr)
		if err != nil {
			// handle expired
			if errors.Is(err, jwt.ErrTokenExpired) {
				return exception.New(401, "Token expired")
			}
			return exception.New(401, "Invalid token")
		}

		user := &entity.User{
			ID:       claims.ID,
			Email:    claims.Email,
			Username: claims.Username,
			Role:     claims.Role,
		}
		c.Locals("credential", user)

		return c.Next()
	}
}
