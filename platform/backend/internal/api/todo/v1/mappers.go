package api_todo_v1

import (
	todoV1 "api/go-sdk/todo/v1"
	"database/sql"
	"errors"
	"fmt"
	todogendb "platform/backend/internal/domain/todo/gendb"

	"connectrpc.com/connect"
)

func dbErrorToAPI(err error, msg string) *connect.Error {
	if errors.Is(err, sql.ErrNoRows) {
		return connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewError(connect.CodeInternal, fmt.Errorf("%s: %w", msg, err))
}

func dbTodoToAPI(todo *todogendb.Todo) *todoV1.Todo {
	return &todoV1.Todo{
		Id:          todo.ID.String(),
		Description: todo.Description,
		Status:      todoV1.TodoStatus(todo.Status),
	}
}

func nonZeroOrDefaultInt32(i int32, def int32) int32 {
	if i == 0 {
		return def
	}

	return i
}
