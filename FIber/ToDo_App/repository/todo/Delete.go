package repository

import (
	"fmt"

	tasks "todo-app/models"

	"gorm.io/gorm"
)

func DeleteTask(db *gorm.DB, id int) error {
	// Find the task by ID
	var task tasks.Task
	result := db.First(&task, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("task with ID %d not found", id)
	}

	// Delete the task
	result = db.Delete(&task)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("Deleted task with ID:", id)

	return nil
}
