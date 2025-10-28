# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a monorepo template for building full-stack applications with:

- **Frontend**: React + TypeScript + Vite + Mantine UI
- **Backend**: Go + Gin + PostgreSQL
- **API**: gRPC/Connect protocol with Protobuf schemas
- **Infrastructure**: Docker Compose for local development

The project uses npm workspaces for JavaScript/TypeScript projects and Go workspaces for Go modules.

## Development Commands

### Installation & Setup

```bash
npm install                    # Install all workspace dependencies
npm run harness start          # Start Docker infrastructure (PostgreSQL)
npm run backend cli db-schema-migrate  # Run database migrations
```

### Code Generation

```bash
npm run schema codegen         # Generate API clients from Protobuf schemas
npm run backend codegen        # Generate Go code (if applicable)
```

### Running Applications

```bash
npm run backend start          # Start backend server with hot-reload (uses modd)
npm run backend debug          # Start backend in debug mode
npm run frontend start         # Start frontend dev server (Vite)
```

### Testing

```bash
npm test                       # Run all tests across workspaces
npm run test:coverage          # Run tests with coverage
npm run backend test:verbose   # Run backend tests with verbose output
npm run backend test:integration  # Run backend integration tests
npm run backend test:e2e       # Run backend e2e tests
npm run backend test:all       # Run all backend tests (unit + integration + e2e)
```

### Linting & Formatting

```bash
npm run lint                   # Lint all workspaces
npm run backend lint:fix       # Auto-fix backend linting issues
npm run format                 # Format all files with Prettier
npm run format:check           # Check formatting without making changes
```

### Building

```bash
npm run build                  # Build all workspaces
```

### Database Operations

```bash
npm run backend cli db-schema-migrate   # Apply database migrations
npm run backend cli db-schema-rollback  # Rollback last migration
```

## Architecture

### Monorepo Structure

- **`_schema/`**: Protobuf schema definitions (`.proto` files in `protos/`)
  - Uses Buf for linting and code generation
  - Generates both Go SDK (`api/go-sdk/`) and TypeScript SDK (`api/web-sdk/`)

- **`_harness/`**: Docker Compose configuration for local infrastructure
  - PostgreSQL database on port 5432
  - Network: `template_network`

- **`api/`**: Generated API client libraries
  - `go-sdk/`: Generated Go code (Connect RPC + Protobuf)
  - `web-sdk/`: Generated TypeScript code (Bufbuild ES + Protobuf)

- **`platform/backend/`**: Go backend service
  - `cmd/server/`: HTTP server entry point
  - `cmd/cli/`: CLI tool for database operations
  - `internal/api/`: API handlers
  - `internal/domain/`: Domain logic and business rules
  - `data/`: Database migrations (using Goose) and SQL queries
  - Uses Gin framework with sessions and CORS
  - Hot-reload with `modd` during development

- **`platform/frontend/`**: React frontend application
  - `src/pages/`: Page components
  - `src/components/`: Reusable UI components
  - `src/context-providers/`: React context providers (including Connect API client)
  - Uses Vite for build tooling
  - UI framework: Mantine
  - API client: Connect Web (@connectrpc/connect-web)

### Key Technologies

**Backend:**

- Framework: Gin (HTTP router)
- Database: PostgreSQL with lib/pq driver
- Migrations: Goose (pressly/goose/v3)
- API Protocol: Connect RPC (gRPC-compatible)
- Development: modd for hot-reload
- Linting: golangci-lint

**Frontend:**

- Framework: React 19
- Build Tool: Vite 7
- UI Library: Mantine 8
- Routing: React Router 7
- State Management: TanStack Query (React Query)
- API Client: Connect Web
- Linting: ESLint 9

**Schema & Code Generation:**

- Protobuf with Buf
- Generates both Go (Connect RPC) and TypeScript (Bufbuild ES) clients

### API Communication

The frontend and backend communicate via Connect RPC, which is compatible with gRPC but works over standard HTTP/1.1 and HTTP/2. The API schema is defined in `_schema/protos/` and code is generated using Buf.

### Database

- PostgreSQL database managed via Docker Compose
- Migrations stored in `platform/backend/data/migrations/`
- SQL queries in `platform/backend/data/queries/`
- Migration tool: Goose (forward and rollback migrations supported)
- Database is initialized via `data/init_db.go`

### Development Workflow

1. Start infrastructure: `npm run harness start`
2. Run migrations: `npm run backend cli db-schema-migrate`
3. Start backend: `npm run backend start` (auto-reloads on file changes)
4. Start frontend: `npm run frontend start` (Vite dev server with HMR)

When modifying Protobuf schemas:

1. Edit files in `_schema/protos/`
2. Run `npm run schema codegen` to regenerate API clients
3. Restart backend and frontend to pick up new types

### Configuration

Backend configuration is managed through environment variables (loaded from `.env` file):

- Environment name
- Server settings (host, port)
- Database connection (PostgreSQL)
- Session configuration
- Frontend URL (for CORS)

Configuration types are defined in `platform/backend/config/`.

### Go Workspace

The project uses Go workspaces (`go.work`) with the backend as the primary module. When adding new Go modules, update `go.work` accordingly.
