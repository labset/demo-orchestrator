package data

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func Migrate(db *sql.DB) error {
	err := setupGoose()
	if err != nil {
		return fmt.Errorf("failed to set up goose: %w", err)
	}

	// platform
	upErr := goose.Up(db, "migrations")
	if upErr != nil {
		return fmt.Errorf("failed to run migrations: %w", upErr)
	}

	return nil
}

func MigrateTo(db *sql.DB, version int64) error {
	err := setupGoose()
	if err != nil {
		return fmt.Errorf("failed to set up goose: %w", err)
	}

	upErr := goose.UpTo(db, "migrations", version)
	if upErr != nil {
		return fmt.Errorf("failed to migrate to version %d: %w", version, upErr)
	}

	return nil
}
