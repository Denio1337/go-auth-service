package main

import (
	"app/internal/config"
	"app/internal/router"
	"app/internal/storage"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create app instance
	app := fiber.New(fiber.Config{
		Prefork:      true,              // Spawn multiple Go processes listening on the same port
		ServerHeader: "Go Auth Service", // "Server" HTTP-header
		AppName:      "Go Auth Service",
	})

	// DB Connection
	storage.MustConnect()

	// Setup routes
	router.SetupRoutes(app)

	// Run application
	log.Fatal(app.Listen(config.Get("APP_ADDRESS")))
}
