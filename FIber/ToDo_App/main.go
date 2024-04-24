package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"todo-app/db"

	"github.com/ServiceWeaver/weaver"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

// app is the main component of the application. weaver.Run creates
// it and passes it to serve.
type app struct {
	weaver.Implements[weaver.Main]
	todo weaver.Listener
}

func serve(ctx context.Context, a *app) error {
	// Initialize Fiber app
	fiberApp := fiber.New()

	// Add middleware for logging
	fiberApp.Use(logger.New())

	// Connect to the database
	dbConnection, err := db.ConnectDB()
	if err != nil {
		// Handle database connection error gracefully
		log.Println("Failed to connect to the database:", err)
		return err
	}

	// Middleware function to attach database connection to Fiber context
	fiberApp.Use(func(c *fiber.Ctx) error {
		c.Locals("db", dbConnection)
		return c.Next()
	})

	// Setup routes
	setupRoutes(fiberApp)

	// Convert fiber app to http.Handler
	handler := adaptor.FiberApp(fiberApp)

	// The todo listener will be used by the http.Serve method
	fmt.Printf("todo listener available on %v\n", a.todo)

	// Serve the application
	if err := http.Serve(a.todo, handler); err != nil {
		log.Println("Failed to start Fiber app:", err)
		return err
	}

	return nil
}

func main() {
	if err := weaver.Run(context.Background(), serve); err != nil {
		log.Fatal(err)
	}
}
