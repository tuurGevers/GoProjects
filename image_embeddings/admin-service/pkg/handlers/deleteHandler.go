package handlers

import (
	"log"
	"shared/pkg/db"
	cloudinaryservice "shared/pkg/util/cloudinary"

	"github.com/gofiber/fiber/v2"
)

func DeleteMultiple(ctx *fiber.Ctx) error {
	// Get the IDs from the request body
	var ids []int
	err := ctx.BodyParser(&ids)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	log.Printf("Deleting data with IDs: %v", ids)

	// Delete the data from the database
	for _, id := range ids {
		res, err := db.DeleteData(id)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete data",
			})
		}
		defer res.Body.Close()
	}

	ctx.SendStatus(fiber.StatusOK)
	return nil
}

func ClearFolder(ctx *fiber.Ctx) error {
	folders, err := cloudinaryservice.FetchFolder()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch folders",
		})
	}

	urls := make([]string, 0)
	for _, folder := range folders {
		urls = append(urls, folder.SecureURL)
	}

	res, err := db.DeleteMultipleByUrl(urls)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete data",
		})
	}

	defer res.Body.Close()
	return ctx.SendStatus(fiber.StatusOK)

}
