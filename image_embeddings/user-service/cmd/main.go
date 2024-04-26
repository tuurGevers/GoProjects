// user-service/cmd/main.go

package user

import (

	// Import the shared weaver utility

	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	// Assuming user-service has its own database management or other specific packages
	db "user-service/pkg/model"
)

// NewService creates a new Fiber app and sets up the user-service routes
func NewService() (*fiber.App, error) {
	// Initialize Fiber app
	fiberApp := fiber.New()

	// Add middleware for logging
	fiberApp.Use(logger.New())

	// Connect to the database (optional based on user-service needs)
	dbConnection, err := db.ConnectDB()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	// Middleware to attach database connection to Fiber context
	fiberApp.Use(func(c *fiber.Ctx) error {
		c.Locals("db", dbConnection)
		return c.Next()
	})

	// Setup user-specific routes
	setupUserRoutes(fiberApp)

	// Everything set up successfully
	return fiberApp, nil
}
