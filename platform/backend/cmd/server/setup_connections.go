package main

import (
	"database/sql"
	"fmt"
	"platform/backend/config"
	"platform/backend/data"
)

type connections struct {
	db *sql.DB
}

func setupConnections(cfg config.Config) (*connections, error) {
	db, dbErr := data.InitDb(cfg.Database)
	if dbErr != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", dbErr)
	}

	return &connections{
		db: db,
	}, nil
}
