-- +goose Up
CREATE TABLE todo (
    id UUID PRIMARY KEY,
    description TEXT NOT NULL,
    status INTEGER NOT NULL DEFAULT 1,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Add indexes for common queries
CREATE INDEX idx_todo_status ON todo(status);
CREATE INDEX idx_todo_created_at ON todo(created_at);
CREATE INDEX idx_todo_updated_at ON todo(updated_at);


-- +goose Down
DROP TABLE IF EXISTS todo;
DROP INDEX IF EXISTS idx_todo_status;
DROP INDEX IF EXISTS idx_todo_created_at;
DROP INDEX IF EXISTS idx_todo_updated_at;