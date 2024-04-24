package main

import (
	"todo-app/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	dbConnection := db.ConnectDB()

	defer dbConnection.Close()
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", dbConnection)
		return c.Next()
	})
	setupRoutes(app)

	app.Listen(":3000")
}
