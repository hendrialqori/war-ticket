package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/hendrialqori/war-ticket/backend/internal/config"
	"github.com/hendrialqori/war-ticket/backend/internal/exception"
	"github.com/hendrialqori/war-ticket/backend/internal/middleware"
)

func main() {
	config.LoadConfig()
	appConfig := config.GetAppConfig()

	app := fiber.New(fiber.Config{
		AppName:      appConfig.Name,
		ErrorHandler: middleware.NewErrorMiddleware(),
	})

	app.Use(recover.New())

	app.Get("/error", func(c *fiber.Ctx) error {
		return exception.ErrDataNotFound
	})

	stringify, _ := json.Marshal(appConfig)
	fmt.Println(fmt.Sprintln(string(stringify)))

	log.Fatal(app.Listen(":" + appConfig.Port))
}
