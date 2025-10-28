package api_todo_v1

import (
	todoV1 "api/go-sdk/todo/v1"
	"context"

	"github.com/gofrs/uuid/v5"

	"connectrpc.com/connect"
	"github.com/rs/zerolog/log"
)

func (s *todoService) GetTodo(
	ctx context.Context,
	request *connect.Request[todoV1.GetTodoRequest],
) (*connect.Response[todoV1.GetTodoResponse], error) {
	identifier := uuid.FromStringOrNil(request.Msg.GetId())
	if identifier.IsNil() {
		log.Warn().
			Str("provided_id", request.Msg.GetId()).
			Msg("invalid todo ID provided")

		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			ErrInvalidTodoID,
		)
	}

	log.Debug().
		Str("todo_id", identifier.String()).
		Msg("getting todo")

	found, err := s.store.GetTodo(ctx, identifier)
	if err != nil {
		log.Error().
			Err(err).
			Str("todo_id", identifier.String()).
			Msg("failed to get todo")

		return nil, dbErrorToAPI(err, "failed to get todo")
	}

	log.Debug().
		Str("todo_id", identifier.String()).
		Msg("todo retrieved successfully")

	return connect.NewResponse(&todoV1.GetTodoResponse{
		Todo: dbTodoToAPI(found),
	}), nil
}
