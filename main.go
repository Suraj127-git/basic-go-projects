package main

import (
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

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New(fiber.Config{})

	app.Use(logger.New())

	setupRoutes(app)

	address := ":" + os.Getenv("APP_PORT")

	log.Fatal(app.Listen(address))
}
