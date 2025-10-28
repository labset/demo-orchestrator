package main

import (
	"log"
	"platform/backend/config"
	"platform/backend/data"
	"strconv"

	"github.com/spf13/cobra"
)

func DbSchemaRollbackCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "db-schema-rollback",
		Args:  cobra.MinimumNArgs(1),
		Short: "Rollback the database schema to the previous version",
		Run: func(cmd *cobra.Command, args []string) {
			versionStr := args[0]

			version, err := strconv.Atoi(versionStr)
			if err != nil {
				log.Fatalf("Invalid version number: %v", err)
			}

			log.Printf("Rolling back database schema to version: %v\n", version)

			cfg := config.Load()

			db, err := data.InitDb(cfg.Database)
			if err != nil {
				log.Fatalf("Failed to initialize database: %v", err)
			}

			defer func() { _ = db.Close() }()

			rollbackErr := data.RollbackTo(db, int64(version))
			if rollbackErr != nil {
				log.Fatalf("Failed to rollback database schema: %v", rollbackErr)
			}

			log.Println("Database schema rolled back successfully.")
		},
	}
}
