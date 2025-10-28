package main

import (
	"log"
	"net/http"
	"platform/backend/config"
	"platform/backend/data"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	cfg := config.Load()

	conn, connErr := setupConnections(cfg)
	if connErr != nil {
		log.Fatalf("failed to initialise connections: %v", connErr)
	}

	migrateErr := data.Migrate(conn.db)
	if migrateErr != nil {
		log.Fatalf("failed to migrate db: %v", migrateErr)
	}

	deps, depsErr := setupDependencies(cfg, conn)
	if depsErr != nil {
		log.Fatalf("failed to initialise clients: %v", depsErr)
	}

	router := setupRouter(cfg, deps)

	log.Printf("starting server on :8080")

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      h2c.NewHandler(router, &http2.Server{}),
		ReadTimeout:  5 * time.Second,   //nolint: mnd
		WriteTimeout: 10 * time.Second,  //nolint: mnd
		IdleTimeout:  120 * time.Second, //nolint: mnd
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
