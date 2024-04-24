package handlers

import (
	"fmt"
	"strconv"
	repository "todo-app/repository/todo"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DeleteTask(ctx *fiber.Ctx) error {

	//get the id from the url
	param, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	fmt.Println("Deleting task with ID:", param)

	err = repository.DeleteTask(
		ctx.Locals("db").(*gorm.DB),
		param,
	)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusOK)
}
