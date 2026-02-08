package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hendrialqori/war-ticket/backend/internal/handler"
)

func SetupUserRoute(app *fiber.App, handler handler.UserHandler, jwtMiddleware fiber.Handler) {
	publicApp := app.Group("/public")

	publicApp.Post("/register", handler.Register)
	publicApp.Post("/login", handler.Login)
	publicApp.Post("/activation", handler.SetActive)

	protectApp := app.Group("/credential", jwtMiddleware)
	protectApp.Get("/me", handler.GetProfile)

}
