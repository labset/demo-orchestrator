-- name: CreateTodo :one
INSERT INTO todo (id, description, status)
VALUES (@id, @description, @status)
RETURNING id, description, status, created_at, updated_at;

-- name: GetTodo :one
SELECT id, description, status, created_at, updated_at
FROM todo
WHERE id = @id;

-- name: UpdateTodo :one
UPDATE todo
SET description = @description, status = @status, updated_at = CURRENT_TIMESTAMP
WHERE id = @id
RETURNING id, description, status, created_at, updated_at;

-- name: DeleteTodo :exec
DELETE FROM todo
WHERE id = @id;

-- name: ListTodos :many
SELECT id, description, status, created_at, updated_at
FROM todo
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: ListTodosByStatus :many
SELECT id, description, status, created_at, updated_at
FROM todo
WHERE status = @status
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;