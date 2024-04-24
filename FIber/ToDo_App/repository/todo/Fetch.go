package repository

import (
	tasks "todo-app/models"

	"gorm.io/gorm"
)

// fetches all tasks from the database
func FetchAllTasks(db *gorm.DB) ([]tasks.Task, error) {
	var allTasks []tasks.Task

	// Retrieve all tasks from the database
	result := db.Find(&allTasks)
	if result.Error != nil {
		return nil, result.Error
	}

	return allTasks, nil
}
