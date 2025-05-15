package main

import (
	"app/internal/config"
	"app/internal/router"
	"log"
)

// @title           Go auth service API
// @version         2.0
// @description     Go auth service API

// @license.name  MIT
// @license.url   https://mit-license.org/

// @host      localhost:8080
// @BasePath  /api

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// Create application instance
	app := router.New()

	// Run application
	log.Fatal(app.Listen(config.Get(config.EnvAppAddress)))
}
