# Content Management System API

## Overview

This API provides a complete content management system with folder organization, label-based categorization, rich text documents, and full-text search capabilities.

## API Version

- Package: `cms.v1`
- Protocol: gRPC/Connect

## Data Models

### Folder
Hierarchical container for organizing documents.

**Fields:**
- `id` (int64): Unique identifier
- `name` (string): Display name
- `parent_id` (optional int64): Parent folder ID (null for root folders)
- `created_at` (timestamp): Creation time
- `updated_at` (timestamp): Last update time

**Constraints:**
- Cannot delete folder if it contains subfolders or documents
- Folder names must be unique within same parent

### Label
Tag for cross-cutting categorization of documents.

**Fields:**
- `id` (int64): Unique identifier
- `name` (string): Display name
- `color` (string): Hex color code (e.g., "#FF5733")
- `created_at` (timestamp): Creation time

**Constraints:**
- Label names must be unique globally

### Document
Content item with rich text formatting.

**Fields:**
- `id` (int64): Unique identifier
- `title` (string): Document title
- `content` (string): Rich text content (HTML/JSON from Tiptap)
- `folder_id` (int64): Parent folder ID
- `labels` (repeated Label): Applied labels
- `created_at` (timestamp): Creation time
- `updated_at` (timestamp): Last update time

### SearchResult
Document search result with relevance ranking.

**Fields:**
- `document` (Document): The matching document
- `rank` (float): Relevance score (higher = more relevant)
- `snippet` (string): Highlighted content excerpt

### FilterMode
Enum for label filtering logic.

**Values:**
- `FILTER_MODE_UNSPECIFIED` (0): Default
- `FILTER_MODE_OR` (1): Match ANY label (OR logic)
- `FILTER_MODE_AND` (2): Match ALL labels (AND logic)

## Service: ContentService

### Folder Management (5 RPCs)

#### CreateFolder
Creates a new folder.

**Request:**
- `name` (string): Folder name
- `parent_id` (optional int64): Parent folder ID

**Response:**
- `folder` (Folder): Created folder

#### GetFolder
Retrieves a specific folder.

**Request:**
- `id` (int64): Folder ID

**Response:**
- `folder` (Folder): Requested folder

#### ListFolders
Lists folders, optionally filtered by parent.

**Request:**
- `parent_id` (optional int64): Parent folder ID (omit for root folders)
- `page_limit` (int32): Max results per page
- `page_offset` (int32): Pagination offset

**Response:**
- `folders` (repeated Folder): List of folders
- `next_page_offset` (int32): Next page offset
- `has_more` (bool): More results available

#### UpdateFolder
Updates a folder's name.

**Request:**
- `id` (int64): Folder ID
- `name` (string): New name

**Response:**
- `folder` (Folder): Updated folder

#### DeleteFolder
Deletes an empty folder.

**Request:**
- `id` (int64): Folder ID

**Response:**
- `success` (bool): Operation success

**Note:** Fails if folder contains subfolders or documents (RESTRICT constraint).

### Label Management (4 RPCs)

#### CreateLabel
Creates a new label.

**Request:**
- `name` (string): Label name
- `color` (string): Hex color code

**Response:**
- `label` (Label): Created label

#### ListLabels
Lists all labels.

**Request:**
- `page_limit` (int32): Max results per page
- `page_offset` (int32): Pagination offset

**Response:**
- `labels` (repeated Label): List of labels
- `next_page_offset` (int32): Next page offset
- `has_more` (bool): More results available

#### UpdateLabel
Updates label properties.

**Request:**
- `id` (int64): Label ID
- `name` (optional string): New name
- `color` (optional string): New color

**Response:**
- `label` (Label): Updated label

#### DeleteLabel
Deletes a label.

**Request:**
- `id` (int64): Label ID

**Response:**
- `success` (bool): Operation success

### Document Management (5 RPCs)

#### CreateDocument
Creates a new document.

**Request:**
- `title` (string): Document title
- `content` (string): Rich text content
- `folder_id` (int64): Parent folder ID

**Response:**
- `document` (Document): Created document

#### GetDocument
Retrieves a specific document.

**Request:**
- `id` (int64): Document ID

**Response:**
- `document` (Document): Requested document

#### ListDocuments
Lists documents with optional filtering.

**Request:**
- `folder_id` (optional int64): Filter by folder
- `label_ids` (repeated int64): Filter by labels
- `filter_mode` (FilterMode): Label filter logic (AND/OR)
- `page_limit` (int32): Max results per page
- `page_offset` (int32): Pagination offset

**Response:**
- `documents` (repeated Document): List of documents
- `next_page_offset` (int32): Next page offset
- `has_more` (bool): More results available

#### UpdateDocument
Updates document properties.

**Request:**
- `id` (int64): Document ID
- `title` (optional string): New title
- `content` (optional string): New content
- `folder_id` (optional int64): New folder ID

**Response:**
- `document` (Document): Updated document

#### DeleteDocument
Deletes a document.

**Request:**
- `id` (int64): Document ID

**Response:**
- `success` (bool): Operation success

### Document-Label Association (2 RPCs)

#### AddLabelToDocument
Adds a label to a document.

**Request:**
- `document_id` (int64): Document ID
- `label_id` (int64): Label ID

**Response:**
- `success` (bool): Operation success

#### RemoveLabelFromDocument
Removes a label from a document.

**Request:**
- `document_id` (int64): Document ID
- `label_id` (int64): Label ID

**Response:**
- `success` (bool): Operation success

### Search (1 RPC)

#### SearchDocuments
Full-text search across documents.

**Request:**
- `query` (string): Search query text
- `folder_id` (optional int64): Limit search to folder
- `label_ids` (repeated int64): Filter by labels
- `filter_mode` (FilterMode): Label filter logic (AND/OR)
- `page_limit` (int32): Max results per page
- `page_offset` (int32): Pagination offset

**Response:**
- `results` (repeated SearchResult): Search results with ranking
- `next_page_offset` (int32): Next page offset
- `has_more` (bool): More results available
- `total_count` (int32): Total matching documents

## Generated SDKs

### Go SDK
- Location: `api/go-sdk/cms/v1/`
- Package: `cmsV1`
- Connect service: `cmsV1connect.ContentServiceClient`

### TypeScript SDK
- Location: `api/web-sdk/src/cms/v1/`
- Import: `import { ContentService } from '@labset/web-sdk/cms/v1/cms_service_pb'`

## Implementation Notes

1. **Folder Deletion**: Protected by database RESTRICT constraints. Must be empty.
2. **Label Filtering**: Supports both AND (all labels) and OR (any label) logic.
3. **Search**: Uses PostgreSQL full-text search with tsvector and ts_rank.
4. **Rich Text**: Content stored as HTML/JSON from Tiptap editor.
5. **Pagination**: All list operations support offset-based pagination.

## Next Steps

- **Backend**: Implement RPC handlers in `platform/backend/internal/api/`
- **Database**: Create migrations in `platform/backend/data/migrations/`
- **Frontend**: Integrate TypeScript SDK in React components
