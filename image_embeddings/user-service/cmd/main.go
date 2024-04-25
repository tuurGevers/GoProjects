// user-service/cmd/main.go

package main

import (
	"context"
	"log"

	weaverutil "shared/pkg" // Import the shared weaver utility

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	// Assuming user-service has its own database management or other specific packages
	db "user-service/pkg/model"
)

func setupFiberApp() *fiber.App {
	// Initialize Fiber app
	fiberApp := fiber.New()

	// Add middleware for logging
	fiberApp.Use(logger.New())

	// Connect to the database (optional based on user-service needs)
	dbConnection, err := db.ConnectDB()
	if err != nil {
		log.Println("Failed to connect to the database:", err)
		return nil
	}

	// Middleware to attach database connection to Fiber context
	fiberApp.Use(func(c *fiber.Ctx) error {
		c.Locals("db", dbConnection)
		return c.Next()
	})

	// Setup user-specific routes
	setupUserRoutes(fiberApp)

	return fiberApp
}

func main() {
	// Setup the Fiber app
	fiberApp := setupFiberApp()
	if fiberApp == nil {
		log.Fatal("Failed to initialize Fiber app")
	}

	// Run the service using the shared Service Weaver utility
	weaverutil.RunService(context.Background(), fiberApp, true)
}
