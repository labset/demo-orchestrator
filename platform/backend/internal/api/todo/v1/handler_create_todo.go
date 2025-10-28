package api_todo_v1

import (
	todoV1 "api/go-sdk/todo/v1"
	"context"
	todogendb "platform/backend/internal/domain/todo/gendb"

	"github.com/gofrs/uuid/v5"

	"connectrpc.com/connect"
	"github.com/rs/zerolog/log"
)

func (s *todoService) CreateTodo(
	ctx context.Context,
	request *connect.Request[todoV1.CreateTodoRequest],
) (*connect.Response[todoV1.CreateTodoResponse], error) {
	id := uuid.Must(uuid.NewV4())

	log.Debug().
		Str("todo_id", id.String()).
		Str("description", request.Msg.GetDescription()).
		Msg("creating todo")

	todo, err := s.store.CreateTodo(ctx, &todogendb.CreateTodoParams{
		ID:          id,
		Description: request.Msg.GetDescription(),
		Status:      int32(todoV1.TodoStatus_TODO_STATUS_PENDING), // Set initial status to PENDING
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("todo_id", id.String()).
			Msg("failed to create todo")

		return nil, dbErrorToAPI(err, "failed to create todo")
	}

	log.Info().
		Str("todo_id", id.String()).
		Msg("todo created successfully")

	return connect.NewResponse(&todoV1.CreateTodoResponse{
		Todo: dbTodoToAPI(todo),
	}), nil
}
