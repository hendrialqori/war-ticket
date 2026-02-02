package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/hendrialqori/war-ticket/backend/internal/config"
)

func main() {
	config.LoadConfig()
	appConfig := config.GetAppConfig()

	app := 	 fiber.New(fiber.Config{
		AppName: appConfig.Name,
	})

	stringify, _ := json.Marshal(appConfig)
	fmt.Println(fmt.Sprintln(string(stringify)))

	log.Fatal(app.Listen(":" + appConfig.Port))
}