package data

import (
	"database/sql"
	"fmt"
	"platform/backend/config"

	_ "github.com/lib/pq"
)

func InitDb(cfg config.DatabaseConfig) (*sql.DB, error) {
	connectionString := cfg.ConnectionString()

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	return db, nil
}
