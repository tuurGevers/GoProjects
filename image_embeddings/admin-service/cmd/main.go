// admin-service/cmd/main.go

package main

import (
	"context"
	"log"

	db "admin-service/pkg/model"

	weaverutil "shared/pkg" // Import the shared weaver utility

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupFiberApp() *fiber.App {
	// Initialize Fiber app
	fiberApp := fiber.New()

	// Add middleware for logging
	fiberApp.Use(logger.New())

	// Connect to the database
	dbConnection, err := db.ConnectDB()
	if err != nil {
		log.Println("Failed to connect to the database:", err)
		return nil
	}

	// Middleware function to attach database connection to Fiber context
	fiberApp.Use(func(c *fiber.Ctx) error {
		c.Locals("db", dbConnection)
		return c.Next()
	})

	// Setup routes
	setupRoutes(fiberApp)

	return fiberApp
}

func main() {
	// Setup the Fiber app
	fiberApp := setupFiberApp()
	if fiberApp == nil {
		log.Fatal("Failed to initialize Fiber app")
	}

	// Run the service with Service Weaver
	weaverutil.RunService(context.Background(), fiberApp, false)
}
