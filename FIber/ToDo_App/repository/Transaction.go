package repository

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// DBTxFunc type for functions that run inside a transaction
type DBTxFunc func(tx pgx.Tx) error

// withTransaction is a helper function to manage transactions
func WithTransaction(db *pgxpool.Pool, fn DBTxFunc) error {
	// Begin a transaction
	tx, err := db.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background()) // Ensure rollback is called if needed

	// Execute the function within the transaction
	err = fn(tx)
	if err != nil {
		return err
	}

	// Commit the transaction
	if err = tx.Commit(context.Background()); err != nil {
		return err
	}
	return nil
}
