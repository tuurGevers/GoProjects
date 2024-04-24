package main

import (
	"todo-app/handlers"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/tasks", handlers.FetchAllTasks)
	app.Post("/tasks", handlers.CreateTask)
	app.Delete("/tasks/:id", handlers.DeleteTask)
}
