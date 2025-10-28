package main

import (
	"log"
	"platform/backend/config"
	"platform/backend/data"
	"strconv"

	"github.com/spf13/cobra"
)

func DbSchemaMigrateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "db-schema-migrate",
		Args:  cobra.MinimumNArgs(1),
		Short: "Apply the latest database schema migrations",
		Run: func(cmd *cobra.Command, args []string) {
			versionStr := args[0]

			version, err := strconv.Atoi(versionStr)
			if err != nil {
				log.Fatalf("Invalid version number: %v", err)
			}

			log.Printf("Migrate database schema to version: %v\n", version)

			cfg := config.Load()

			db, err := data.InitDb(cfg.Database)
			if err != nil {
				log.Fatalf("Failed to initialize database: %v", err)
			}

			defer func() { _ = db.Close() }()

			migrateErr := data.MigrateTo(db, int64(version))
			if migrateErr != nil {
				log.Fatalf("Failed to migrate database schema: %v", migrateErr)
			}

			log.Println("Database schema migrated successfully.")
		},
	}
}
