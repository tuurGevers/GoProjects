package admin

import (
	"admin-service/pkg/handlers"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Post("/insert", handlers.InsertEmbedding)
	app.Post("/insert-image", handlers.InsertBasicEmbedding)
	app.Post("/insert-folder", handlers.InsertFolder)
	app.Post("/delete", handlers.DeleteMultiple)
	app.Post("/delete-folder", handlers.ClearFolder)
}
