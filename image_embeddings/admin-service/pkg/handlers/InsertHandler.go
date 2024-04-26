package handlers

import (
	"admin-service/pkg/service"
	"io"
	"shared/pkg/db"
	"shared/pkg/models"

	"github.com/gofiber/fiber/v2"
)

// InsertTask inserts a task into the database
func InsertEmbedding(ctx *fiber.Ctx) error {
	// get vectors from body
	var description models.Embedding

	err := ctx.BodyParser(&description)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	vectors, err := service.Embed(description.Description)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate embeddings",
		})
	}

	res, err := db.InsertData(vectors)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert data",
		})
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	ctx.SendString(string(body))
	return nil
}
