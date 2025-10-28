-- +goose Up
-- ============================================================================
-- Content Management System Schema
-- ============================================================================
-- This migration creates the complete CMS database schema with:
-- - folders: Hierarchical organization with parent-child relationships
-- - labels: Tags for cross-cutting categorization
-- - documents: Rich text content with full-text search
-- - document_labels: Many-to-many relationship between documents and labels
--
-- Key features:
-- - RESTRICT constraints prevent deletion of folders with content
-- - CASCADE constraints cleanup associations when parent is deleted
-- - Full-text search using tsvector with GIN index
-- - Auto-updating search vector via trigger
-- - UNIQUE constraints prevent duplicate names
-- ============================================================================

-- ============================================================================
-- Table: folders
-- Hierarchical container for organizing documents
-- ============================================================================
CREATE TABLE folders (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    parent_id BIGINT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    -- Self-referential foreign key with RESTRICT to prevent deleting folders with subfolders
    CONSTRAINT fk_folders_parent
        FOREIGN KEY (parent_id)
        REFERENCES folders(id)
        ON DELETE RESTRICT,

    -- Prevent duplicate folder names at the same level
    CONSTRAINT uq_folders_parent_name
        UNIQUE (parent_id, name)
);

-- Indexes for folder queries
CREATE INDEX idx_folders_parent_id ON folders(parent_id);
CREATE INDEX idx_folders_created_at ON folders(created_at);

-- ============================================================================
-- Table: labels
-- Tags for cross-cutting categorization of documents
-- ============================================================================
CREATE TABLE labels (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    color VARCHAR(7) NOT NULL, -- Hex color code (e.g., #FF5733)
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    -- Prevent duplicate label names globally
    CONSTRAINT uq_labels_name
        UNIQUE (name)
);

-- Index for label queries
CREATE INDEX idx_labels_name ON labels(name);

-- ============================================================================
-- Table: documents
-- Rich text content with full-text search capability
-- ============================================================================
CREATE TABLE documents (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(500) NOT NULL,
    content TEXT NOT NULL, -- Stores HTML/JSON from Tiptap editor
    folder_id BIGINT NOT NULL,
    search_vector TSVECTOR, -- Full-text search index
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    -- Foreign key with RESTRICT to prevent deleting folders with documents
    CONSTRAINT fk_documents_folder
        FOREIGN KEY (folder_id)
        REFERENCES folders(id)
        ON DELETE RESTRICT
);

-- Indexes for document queries
CREATE INDEX idx_documents_folder_id ON documents(folder_id);
CREATE INDEX idx_documents_created_at ON documents(created_at);
CREATE INDEX idx_documents_updated_at ON documents(updated_at);

-- GIN index for full-text search (highly efficient for tsvector)
CREATE INDEX idx_documents_search_vector ON documents USING GIN(search_vector);

-- ============================================================================
-- Table: document_labels
-- Many-to-many relationship between documents and labels
-- ============================================================================
CREATE TABLE document_labels (
    document_id BIGINT NOT NULL,
    label_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    -- Composite primary key
    CONSTRAINT pk_document_labels
        PRIMARY KEY (document_id, label_id),

    -- Foreign keys with CASCADE to cleanup when parent is deleted
    CONSTRAINT fk_document_labels_document
        FOREIGN KEY (document_id)
        REFERENCES documents(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_document_labels_label
        FOREIGN KEY (label_id)
        REFERENCES labels(id)
        ON DELETE CASCADE
);

-- Indexes for association queries
CREATE INDEX idx_document_labels_document_id ON document_labels(document_id);
CREATE INDEX idx_document_labels_label_id ON document_labels(label_id);

-- ============================================================================
-- Function: update_document_search_vector
-- Extracts plain text from HTML/JSON content and builds tsvector
-- ============================================================================
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION update_document_search_vector()
RETURNS TRIGGER AS $func$
BEGIN
    NEW.search_vector := setweight(to_tsvector('english', COALESCE(NEW.title, '')), 'A') || setweight(to_tsvector('english', COALESCE(regexp_replace(NEW.content, '<[^>]+>', '', 'g'), '')), 'B');
    RETURN NEW;
END;
$func$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- ============================================================================
-- Trigger: trg_documents_search_vector
-- Auto-update search_vector when document is created or content changes
-- ============================================================================
CREATE TRIGGER trg_documents_search_vector
    BEFORE INSERT OR UPDATE OF title, content
    ON documents
    FOR EACH ROW
    EXECUTE FUNCTION update_document_search_vector();

-- ============================================================================
-- Function: update_updated_at
-- Updates the updated_at timestamp on row modification
-- ============================================================================
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $func$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$func$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- ============================================================================
-- Triggers: Auto-update updated_at timestamps
-- ============================================================================
CREATE TRIGGER trg_folders_updated_at
    BEFORE UPDATE ON folders
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at();

CREATE TRIGGER trg_documents_updated_at
    BEFORE UPDATE ON documents
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at();

-- +goose Down
-- ============================================================================
-- Rollback: Drop all CMS tables, triggers, and functions
-- ============================================================================

-- Drop triggers first
DROP TRIGGER IF EXISTS trg_documents_updated_at ON documents;
DROP TRIGGER IF EXISTS trg_folders_updated_at ON folders;
DROP TRIGGER IF EXISTS trg_documents_search_vector ON documents;

-- Drop functions
DROP FUNCTION IF EXISTS update_updated_at();
DROP FUNCTION IF EXISTS update_document_search_vector();

-- Drop tables in reverse order (respecting foreign key dependencies)
DROP TABLE IF EXISTS document_labels;
DROP TABLE IF EXISTS documents;
DROP TABLE IF EXISTS labels;
DROP TABLE IF EXISTS folders;
