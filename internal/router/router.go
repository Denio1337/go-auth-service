package router

import (
	"app/internal/config"
	"app/internal/router/handler/auth"
	"app/internal/router/handler/ping"
	"app/internal/router/middleware"
	"app/internal/router/types/response"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Create and configure Fiber application
func New() *fiber.App {
	router := fiber.New(fiber.Config{
		Prefork:      true,              // Spawn multiple Go processes listening on the same port
		ServerHeader: "Go Auth Service", // Set "Server" HTTP-header
		AppName:      "Go Auth Service",
		ErrorHandler: handleError,
	})

	// Configure endpoints
	setupRoutes(router)

	return router
}

// Set router api
func setupRoutes(app *fiber.App) {
	// Middleware
	apiGroup := app.Group("/api", logger.New())
	apiGroup.Get("/", ping.Ping)

	// Auth
	authGroup := apiGroup.Group("/auth")

	// Login
	authGroup.Get("/login", middleware.Identified(), auth.Login)

	// Logout
	authGroup.Delete(
		"/logout",
		middleware.Identified(),
		middleware.Protected(auth.CookieNameAccessToken),
		auth.Logout,
	)

	// User GUID
	authGroup.Get(
		"/me",
		middleware.Identified(),
		middleware.Protected(auth.CookieNameAccessToken),
		auth.Me,
	)

	// Register
	authGroup.Post(
		"/register",
		middleware.Identified(),
		middleware.Protected(auth.CookieNameAccessToken),
		auth.Register,
	)

	// Refresh
	authGroup.Patch(
		"/refresh",
		middleware.Identified(),
		middleware.Protected(auth.CookieNameRefreshToken),
		middleware.Webhooked(config.Get(config.EnvWebhookURL)),
		auth.Refresh,
	)
}

// Handle error response
func handleError(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	return c.Status(code).JSON(response.ErrorResponse(err.Error()))
}
