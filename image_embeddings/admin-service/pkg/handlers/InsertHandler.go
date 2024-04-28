package handlers

import (
	"admin-service/pkg/service"
	"log"
	"os"
	"shared/pkg/db"
	"shared/pkg/models"
	"shared/pkg/util/gpt"

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

	log.Printf("VectorsLength: %v url:%s", len(vectors), description.Url)
	res, err := db.InsertData(vectors, description.Url)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert data",
		})
	}
	defer res.Body.Close()

	ctx.SendStatus(fiber.StatusOK)
	return nil
}

func InsertBasicEmbedding(ctx *fiber.Ctx) error {
	// Get URL from request body
	var description models.Embedding
	err := ctx.BodyParser(&description)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Call GptVisionRequest function to generate a description with low detail
	request := gpt.GptVisionRequest(description.Url, false)
	apiKey := os.Getenv("OPENAI_API_KEY")
	endpoint := "https://api.openai.com/v1/chat/completions"

	response, err := gpt.MakeGPTRequest("gpt-4-turbo", apiKey, endpoint, request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate description",
		})
	}

	log.Printf("Response: %v", response.Choices[0].Message.Content)

	// Use the description to generate vectors
	vectors, err := service.Embed(response.Choices[0].Message.Content)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate embeddings",
		})
	}

	log.Printf("VectorsLength: %v, URL: %s", len(vectors), description.Url)
	res, err := db.InsertData(vectors, description.Url)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert data",
		})
	}
	defer res.Body.Close()

	return ctx.SendString(response.Choices[0].Message.Content)
}
