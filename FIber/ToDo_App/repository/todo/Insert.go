package repository

import (
	"fmt"
	tasks "todo-app/models"

	"gorm.io/gorm"
)

func InsertTask(db *gorm.DB, task tasks.Task) error {
	// Create the task in the database
	result := db.Create(&task)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("Inserted new task with ID:", task.ID)
	return nil
}
