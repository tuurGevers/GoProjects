package handlers

import (
	"admin-service/pkg/model"
	"admin-service/pkg/service"
	"fmt"
	"log"
	"os"
	db "shared/pkg/db"
	"shared/pkg/models"
	cloudinaryservice "shared/pkg/util/cloudinary"
	"shared/pkg/util/gpt"
	"sync"

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

// fetxhFOlder handler
func InsertFolder(ctx *fiber.Ctx) error {
	folders, err := cloudinaryservice.FetchFolder()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch folders",
		})
	}

	resultChan := make(chan model.FolderResult, 250) // Buffer of 10
	semaphoreChan := make(chan struct{}, 3)          // Semaphore to limit concurrency to 3
	var wg sync.WaitGroup

	for i, folder := range folders {
		wg.Add(1) // Increment the WaitGroup counter
		log.Printf("Attempting to acquire semaphore for folder %d", i)
		semaphoreChan <- struct{}{}
		log.Printf("Semaphore acquired for folder %d", i)

		go func(i int, folder cloudinaryservice.BriefAssetResult) {
			defer wg.Done()
			defer func() {
				<-semaphoreChan
				log.Printf("Semaphore released for folder %d", i)
			}()

			log.Printf("Starting InsertBasicEmbeddingURL for folder %d", i)
			err := service.InsertBasicEmbeddingURL(folder.SecureURL)
			if err != nil {
				log.Printf("Error in InsertBasicEmbeddingURL for folder %d: %v", i, err)
			} else {
				log.Printf("Successful InsertBasicEmbeddingURL for folder %d", i)
			}
			log.Printf("Sending result for folder %d", i)
			resultChan <- model.FolderResult{
				Index: i,
				Error: err,
			}
			log.Printf("Result sent for folder %d", i)
		}(i, folder)

	}

	go func() {
		wg.Wait()         // Wait for all goroutines to finish
		close(resultChan) // Close the result channel to break the loop below
	}()

	// Process results as they come in
	for result := range resultChan {
		if result.Error != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("Failed to insert data at index %d", result.Index),
			})
		}
		log.Printf("Inserted data at index %d", result.Index)
	}

	return ctx.Status(fiber.StatusOK).JSON(folders)
}
