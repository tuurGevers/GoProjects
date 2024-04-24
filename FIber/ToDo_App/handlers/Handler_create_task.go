package handlers

import (
	tasks "todo-app/models"
	repository "todo-app/repository/todo"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

func CreateTask(ctx *fiber.Ctx) error {

	var task tasks.Task
	if err := ctx.BodyParser(&task); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	println("Creating task:", task.ID, task.Description, task.Completed)

	err := repository.InsertTask(
		ctx.Locals("db").(*pgxpool.Pool),
		task,
	)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusCreated)
}
