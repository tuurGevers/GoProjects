package handlers

import (
	"admin-service/pkg/service"
	"fmt"
	"io"
	"shared/pkg/db"
	sharedModels "shared/pkg/models"
	"shared/pkg/util"
	models "user-service/pkg/model"

	"github.com/gofiber/fiber/v2"
)

func FetchAllVectors(ctx *fiber.Ctx) error {
	fc := util.FiberContext{Ctx: ctx}

	fmt.Println("Fetching all vectors")
	// res := db.FetchTest()
	fmt.Println("result found")

	res, err := db.FetchCollection()

	if err != nil {
		return fc.HandleError(fiber.StatusInternalServerError, "Failed to fetch data")
	}
	fmt.Println("result found")

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	ctx.SendString(string(body))
	return nil
}

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

	// get vectors from body
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

	body, _ := io.ReadAll(res.Body)

	ctx.SendString(string(body))
	return nil
}
