package api_todo_v1

import (
	todoV1 "api/go-sdk/todo/v1"
	"context"
	"database/sql"
	"errors"
	"fmt"
	todogendb "platform/backend/internal/domain/todo/gendb"

	"github.com/gofrs/uuid/v5"

	"connectrpc.com/connect"
)

var (
	errTodoIDRequired = errors.New("todo id is required")
	errInvalidTodoID  = errors.New("invalid todo id format")
	errTodoNotFound   = errors.New("todo not found")
)

func (s *todoService) UpdateTodo(
	ctx context.Context,
	req *connect.Request[todoV1.UpdateTodoRequest],
) (*connect.Response[todoV1.UpdateTodoResponse], error) {
	msg := req.Msg

	// Validate request
	if msg.GetId() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errTodoIDRequired)
	}

	// Parse UUID
	todoID, err := uuid.FromString(msg.GetId())
	if err != nil {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			fmt.Errorf("%w: %w", errInvalidTodoID, err),
		)
	}

	// Get existing todo
	existingTodo, err := s.store.GetTodo(ctx, todoID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, connect.NewError(connect.CodeNotFound, errTodoNotFound)
		}

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Apply updates
	s.applyUpdates(existingTodo, msg)

	// Save updated todo
	updated, err := s.store.UpdateTodo(ctx, &todogendb.UpdateTodoParams{
		ID:          existingTodo.ID,
		Description: existingTodo.Description,
		Status:      existingTodo.Status,
	})
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Convert to proto response
	response := &todoV1.UpdateTodoResponse{
		Todo: dbTodoToAPI(updated),
	}

	return connect.NewResponse(response), nil
}

func (s *todoService) applyUpdates(todo *todogendb.Todo, msg *todoV1.UpdateTodoRequest) {
	updateMask := msg.GetUpdateMask()
	todoUpdate := msg.GetTodo()

	if todoUpdate == nil {
		return
	}

	// Apply updates based on field mask
	if updateMask != nil && len(updateMask.GetPaths()) > 0 {
		// Update only specified fields
		for _, path := range updateMask.GetPaths() {
			switch path {
			case "description":
				todo.Description = todoUpdate.GetDescription()
			case "status":
				todo.Status = int32(todoUpdate.GetStatus())
			}
		}
	} else {
		// If no field mask provided, update all provided fields
		todo.Description = todoUpdate.GetDescription()
		todo.Status = int32(todoUpdate.GetStatus())
	}
}
