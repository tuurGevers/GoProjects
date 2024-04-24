package main

import (
	"fmt"
	"todo-app/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Add middleware for logging
	app.Use(logger.New())

	// Connect to the database
	dbConnection, err := db.ConnectDB()
	if err != nil {
		// Handle database connection error gracefully
		fmt.Println("Failed to connect to the database:", err)
		return
	}

	// Middleware function to attach database connection to Fiber context
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", dbConnection)
		return c.Next()
	})

	// Setup routes
	setupRoutes(app)

	// Start the Fiber app on port 3000
	err = app.Listen(":3000")
	if err != nil {
		// Handle Fiber app startup error gracefully
		fmt.Println("Failed to start Fiber app:", err)
		return
	}
}
