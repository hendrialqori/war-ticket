package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/hendrialqori/war-ticket/backend/internal/config"
	"github.com/hendrialqori/war-ticket/backend/internal/middleware"
	"github.com/hendrialqori/war-ticket/backend/internal/model"
)

func main() {
	// load initialy config
	config.LoadConfig()

	// app config
	appConfig := config.GetAppConfig()
	// database config
	dbConfig := config.GetDatabaseConfig()
	// make connectin to mysql
	db := config.NewMysqlConnection(dbConfig)
	// make auto migration
	model.AutoMigrationModels(db)

	// instace fiber app
	app := fiber.New(fiber.Config{
		AppName:      appConfig.Name,
		ErrorHandler: middleware.NewErrorMiddleware(),
	})

	// write debug at terminal
	app.Use(logger.New())
	// recover so that the server doesn't not die if panic
	app.Use(recover.New())

	log.Fatal(app.Listen(":" + appConfig.Port))
}
