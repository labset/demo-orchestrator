package api_todo_v1

import (
	"api/go-sdk/todo/v1/todoV1connect"
	todogendb "platform/backend/internal/domain/todo/gendb"
)

type todoService struct {
	store todogendb.Querier
}

func newTodoService(store todogendb.Querier) todoV1connect.TodoServiceHandler {
	return &todoService{
		store: store,
	}
}
