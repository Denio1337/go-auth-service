package main

import (
	"app/internal/config"
	"app/internal/router"
	"log"
)

func main() {
	// Create application instance
	app := router.New()

	// Run application
	log.Fatal(app.Listen(config.Get("APP_ADDRESS")))
}
