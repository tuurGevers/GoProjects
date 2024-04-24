package repository

import (
	"context"
	"fmt"
	"todo-app/repository"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func DeleteTask(db *pgxpool.Pool, id int) error {
	return repository.WithTransaction(db, func(tx pgx.Tx) error {
		sql := `DELETE FROM tasks WHERE id = $1`
		_, err := tx.Exec(context.Background(), sql, id)
		if err != nil {
			return err
		}
		fmt.Println("Deleted task with ID:", id)
		return nil
	})
}
