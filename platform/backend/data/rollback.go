package data

import (
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
)

func RollbackTo(db *sql.DB, version int64) error {
	err := setupGoose()
	if err != nil {
		return fmt.Errorf("failed to set up goose: %w", err)
	}

	downErr := goose.DownTo(db, "migrations", version)
	if downErr != nil {
		return fmt.Errorf("failed to rollback to version %d: %w", version, downErr)
	}

	return nil
}
