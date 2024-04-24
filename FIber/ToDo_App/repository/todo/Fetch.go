package repository

import (
	"context"
	"fmt"
	tasks "todo-app/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

// fetches all tasks from the database
func FetchAllTasks(db *pgxpool.Pool) ([]tasks.Task, error) {

	tx, err := db.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(context.Background())

	sql := `SELECT * FROM tasks`

	// check if there are any tasks
	rows, err := tx.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// convert the fetch to tasks
	var allTasks []tasks.Task

	fmt.Println("Fetching all tasks")
	// iterate over the rows
	for rows.Next() {
		var task tasks.Task
		if err := rows.Scan(&task.ID, &task.Description, &task.Completed, &task.CreatedOn, &task.FinishedOn); err != nil {
			return nil, err
		}
		//print task
		fmt.Println("Task:", task.ID, task.Description, task.Completed)
		// append the task to the allTasks slice
		allTasks = append(allTasks, task)
	}

	return allTasks, nil

}
