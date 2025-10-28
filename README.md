# Template

A full-stack application template featuring a React frontend, Go backend, and gRPC/Connect API with PostgreSQL database.

## Quick Start

### Prerequisites

- [Node.js](https://nodejs.org/) v22 or higher
- [Go](https://go.dev/) v1.25 or higher
- [Docker](https://www.docker.com/) and Docker Compose

### Getting Started

- Clone and install dependencies
  ```bash
  npm install
  ```
- Code generation

  ```bash
  npm run schema codegen
  npm run backend codegen
  ```

- Build dependencies

  ```bash
  npm run build
  ```

- Start the development services:

  ```bash
  # Terminal 1: Start Docker services (PostgreSQL)
  npm run harness start

  # Terminal 2: Start backend with live reload
  npm run backend start

  # Terminal 3: Start frontend dev server
  npm run frontend start
  ```

- Access the application:
  - Frontend: http://localhost:5173
  - Backend API: http://localhost:8080

## Project Structure

This is a monorepo with the following workspaces:

- **`_harness/`** - Docker Compose configuration for local development
- **`_schema/`** - Protobuf API schema definitions
- **`api/go-sdk/`** - Generated Go API client (from Protobuf)
- **`api/web-sdk/`** - Generated TypeScript API client (from Protobuf)
- **`platform/backend/`** - Go backend service (Gin + PostgreSQL + Connect RPC)
- **`platform/frontend/`** - React frontend application (Vite + Mantine UI)

## Development

### Code Generation

After modifying Protobuf schemas in `_schema/protos/`:

```bash
npm run schema codegen      # Regenerate API clients
npm run backend codegen     # Regenerate backend code (if applicable)
```

### Running Tests

```bash
npm test                              # Run all tests
npm run test:coverage                 # Run tests with coverage
npm run backend test:integration      # Run integration tests
npm run backend test:e2e              # Run e2e tests
```

### Linting and Formatting

```bash
npm run lint              # Lint all workspaces
npm run backend lint:fix  # Auto-fix backend linting issues
npm run format            # Format code with Prettier
```

### Building for Production

```bash
npm run build    # Build all workspaces
```

### Database Management

```bash
npm run backend cli db-schema-migrate    # Apply migrations
npm run backend cli db-schema-rollback   # Rollback last migration
```

## Technology Stack

### Backend

- **Language:** Go 1.25
- **Framework:** Gin
- **Database:** PostgreSQL 17
- **Migrations:** Goose
- **API:** Connect RPC (gRPC-compatible)

### Frontend

- **Framework:** React 19
- **Build Tool:** Vite 7
- **UI Library:** Mantine 8
- **Routing:** React Router 7
- **State Management:** TanStack Query
- **API Client:** Connect Web

### API & Code Generation

- **Schema:** Protocol Buffers (Protobuf)
- **Tooling:** Buf
- **Protocol:** Connect RPC

## Environment Configuration

Create a `.env` file in `platform/backend/` based on your environment needs. The backend expects configuration for:

- Environment name
- Server settings (host, port)
- Database connection
- Session configuration
- Frontend URL (for CORS)

## Available Scripts

### Root Level

- `npm run harness` - Manage Docker infrastructure
- `npm run schema` - Manage Protobuf schemas
- `npm run backend` - Manage backend service
- `npm run frontend` - Manage frontend application
- `npm run build` - Build all workspaces
- `npm run lint` - Lint all workspaces
- `npm test` - Run tests across all workspaces
- `npm run format` - Format code with Prettier

### Backend Specific

- `npm run backend start` - Start with hot-reload
- `npm run backend debug` - Start in debug mode
- `npm run backend cli` - Run CLI commands
- `npm run backend test:verbose` - Run tests with verbose output
- `npm run backend test:all` - Run all test types

### Frontend Specific

- `npm run frontend start` - Start dev server
- `npm run frontend build` - Build for production
- `npm run frontend preview` - Preview production build
