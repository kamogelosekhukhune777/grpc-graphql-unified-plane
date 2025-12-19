package migrations

import (
	"context"
	"database/sql"
)

// InitSchema attempts to bring the database up to date with the migrations
func InitSchema(ctx context.Context, db *sql.DB) error {
	if err := db.PingContext(ctx); err != nil {
		return err
	}

	table := `
	CREATE TABLE IF NOT EXISTS brands (
		id text PRIMARY KEY UNIQUE,
		name text NOT NULL,
		description text NOT NULL,
		created_at timestamp NOT NULL,
		updated_at timestamp NOT NULL
    );`

	_, err := db.ExecContext(ctx, table)
	return err
}
