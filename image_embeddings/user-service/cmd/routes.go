package user

import (
	"user-service/pkg/handlers"

	"github.com/gofiber/fiber/v2"
)

func setupUserRoutes(app *fiber.App) {
	app.Get("/search", handlers.Search)
	app.Get("/searchstring", handlers.SearchEmbedding)
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/folder-test", handlers.FetchFolder)

}
