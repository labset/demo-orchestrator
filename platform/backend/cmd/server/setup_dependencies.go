package main

import (
	"platform/backend/config"
	todogendb "platform/backend/internal/domain/todo/gendb"
)

type dependencies struct {
	todoStore todogendb.Querier
}

func setupDependencies(_ config.Config, conn *connections) (*dependencies, error) {
	todoStore := todogendb.New(conn.db)

	return &dependencies{
		todoStore: todoStore,
	}, nil
}
