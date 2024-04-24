package handlers

import (
	repository "todo-app/repository/todo"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func FetchAllTasks(ctx *fiber.Ctx) error {
	tasks, err := repository.FetchAllTasks(
		ctx.Locals("db").(*gorm.DB),
	)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"tasks": tasks,
	})
}
