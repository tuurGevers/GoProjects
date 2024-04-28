package handlers

import (
	"admin-service/pkg/service"
	"encoding/json"
	"io"
	"log"
	"shared/pkg/db"
	sharedModels "shared/pkg/models"
	"shared/pkg/util"
	cloudinaryservice "shared/pkg/util/cloudinary"
	models "user-service/pkg/model"

	"github.com/gofiber/fiber/v2"
)

func Search(ctx *fiber.Ctx) error {
	fc := util.FiberContext{Ctx: ctx}

	// get vectors from body
	var vectors models.Search

	err := ctx.BodyParser(&vectors)
	if err != nil {
		return fc.HandleError(fiber.StatusBadRequest, "Invalid request body")
	}

	res, err := db.SearchData(vectors.Vectors)
	if err != nil {
		return fc.HandleError(fiber.StatusInternalServerError, "Failed to fetch data")
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	ctx.SendString(string(body))
	return nil
}

func SearchEmbedding(ctx *fiber.Ctx) error {
	fc := util.FiberContext{Ctx: ctx}

	// Get vectors from body
	var description sharedModels.Embedding

	err := ctx.BodyParser(&description)
	if err != nil {
		return fc.HandleError(fiber.StatusInternalServerError, "Failed to generate embeddings")
	}

	vectors, err := service.Embed(description.Description)
	if err != nil {
		return fc.HandleError(fiber.StatusInternalServerError, "Failed to search data")
	}

	res, err := db.SearchData(vectors)
	if err != nil {
		return fc.HandleError(fiber.StatusInternalServerError, "Failed to read response from database")
	}

	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fc.HandleError(fiber.StatusInternalServerError, "Failed to read response body")
	}

	// Convert body to EmbeddingResponse
	var embeddingResponse sharedModels.SearchResponse
	err = json.Unmarshal(body, &embeddingResponse) // Use json.Unmarshal to parse the response body
	if err != nil {
		return fc.HandleError(fiber.StatusInternalServerError, "Failed to unmarshal response body")
	}

	log.Printf("Response status: %s, Body: %s", res.Status, string(body))

	// Check if the Data slice is empty before attempting to access its elements
	if len(embeddingResponse.Data) == 0 {
		return fc.HandleError(fiber.StatusInternalServerError, "No data found in database response")
	}

	log.Printf("Embedding data found: %v", embeddingResponse.Data[0].AutoID)
	// Now you can safely access the first element of the Data slice
	FinalRes, err := db.QueryData(int(embeddingResponse.Data[0].AutoID)) // Make sure the function accepts int or cast to int64 as necessary
	if err != nil {
		return fc.HandleError(fiber.StatusInternalServerError, "Failed to query data from database")
	}

	defer FinalRes.Body.Close()

	finalBody, err := io.ReadAll(FinalRes.Body) // Use a new variable name to avoid confusion
	if err != nil {
		return fc.HandleError(fiber.StatusInternalServerError, "Failed to read final response body")
	}

	// Send the final response body as the context response
	return ctx.SendString(string(finalBody))
}

// fetxhFOlder handler
func FetchFolder(ctx *fiber.Ctx) error {
	folders, err := cloudinaryservice.FetchFolder()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch folders",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(folders)
}
