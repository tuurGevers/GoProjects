package repository

import (
	"context"
	"fmt"
	todoModel "todo-app/models"
	"todo-app/repository"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func InsertTask(db *pgxpool.Pool, task todoModel.Task) error {
	return repository.WithTransaction(db, func(tx pgx.Tx) error {
		sql := `INSERT INTO tasks (description, completed, created_on, finished_on)
                VALUES ($1, $2, $3, $4)
                RETURNING id`
		row := tx.QueryRow(context.Background(), sql, task.Description, task.Completed, task.CreatedOn, task.FinishedOn)
		var id int
		if err := row.Scan(&id); err != nil {
			return err
		}
		fmt.Println("Inserted new task with ID:", id)
		return nil
	})
}
