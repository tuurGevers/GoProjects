package util

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// FiberContext is a wrapper around fiber.Ctx that provides additional error handling functionality.
type FiberContext struct {
	Ctx *fiber.Ctx
}

// HandleError sends a JSON response with the error message and logs the error.
func (fc FiberContext) HandleError(statusCode int, message string) error {
	log.Printf("Error %d: %s", statusCode, message) // Optional: Add logging here.
	return fc.Ctx.Status(statusCode).JSON(fiber.Map{"error": message})
}
