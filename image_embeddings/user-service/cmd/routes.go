package user

import (
	"github.com/gofiber/fiber/v2"
)

func setupUserRoutes(app *fiber.App) {
	// app.Get("/tasks", handlers.FetchAllTasks)
	// app.Post("/tasks", handlers.CreateTask)
	// app.Delete("/tasks/:id", handlers.DeleteTask)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, User Service!")
	})
}
