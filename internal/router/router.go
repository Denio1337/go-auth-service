package router

import (
	"app/internal/router/handler/auth"
	"app/internal/router/handler/ping"
	"app/internal/router/middleware"
	"app/internal/router/types"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Create and configure Fiber application
func New() *fiber.App {
	// Create app instance
	router := fiber.New(fiber.Config{
		Prefork:      true,              // Spawn multiple Go processes listening on the same port
		ServerHeader: "Go Auth Service", // "Server" HTTP-header
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
	authGroup.Get("/login", middleware.Identified(), auth.Login)
	authGroup.Get("/logout", middleware.Identified(), middleware.Protected(auth.CookieNameAccessToken), auth.Logout)
	authGroup.Get("/me", middleware.Identified(), middleware.Protected(auth.CookieNameAccessToken), auth.Me)
	authGroup.Post("/register", auth.Register)
	//auth.Get

	// User
	// user := api.Group("/user")
	// user.Get("/:id", handler.GetUser)
	// user.Post("/", handler.CreateUser)
	// user.Patch("/:id", middleware.Protected(), handler.UpdateUser)
	// user.Delete("/:id", middleware.Protected(), handler.DeleteUser)

	// Product
	// product := api.Group("/product")
	// product.Get("/", handler.GetAllProducts)
	// product.Get("/:id", handler.GetProduct)
	// product.Post("/", middleware.Protected(), handler.CreateProduct)
	// product.Delete("/:id", middleware.Protected(), handler.DeleteProduct)
}

// Handle error response
func handleError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusBadRequest).JSON(types.Response{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	})
}
