package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/suraj/url-shortener/api/routes"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {

	env_data := godotenv.Load()

	if env_data != nil {
		fmt.Println(env_data)
	}

	app := fiber.New()

	app.Use(logger.New())

	setupRoutes(app)

	address := fmt.Sprintf("0.0.0.0:%s", os.Getenv("APP_PORT"))
	log.Fatal(app.Listen(address))

}
