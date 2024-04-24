package handlers

import (
	repository "todo-app/repository/todo"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

func FetchAllTasks(ctx *fiber.Ctx) error {
	tasks, err := repository.FetchAllTasks(
		ctx.Locals("db").(*pgxpool.Pool),
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
