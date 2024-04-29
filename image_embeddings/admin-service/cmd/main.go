// admin-service/cmd/main.go

package admin

import (
	db "admin-service/pkg/model/db"
	"fmt"

	// Import the shared weaver utility

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewService() (*fiber.App, error) {
	// Initialize Fiber app
	fiberApp := fiber.New()

	// Add middleware for logging
	fiberApp.Use(logger.New())

	// Connect to the database
	dbConnection, err := db.ConnectDB()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	// Middleware function to attach database connection to Fiber context
	fiberApp.Use(func(c *fiber.Ctx) error {
		c.Locals("db", dbConnection)
		return c.Next()
	})

	// Setup routes
	setupRoutes(fiberApp)

	// Everything set up successfully
	return fiberApp, nil
}
