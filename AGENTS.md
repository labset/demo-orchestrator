# AI Product Engineering Team

This repository includes an AI-powered product engineering team implemented as Claude Code slash commands. These agents work together to take a customer feature request from concept to working prototype.

## Team Structure

```
Customer Request
      ↓
┌─────────────────┐
│  Orchestrator   │  ← /product-request
└────────┬────────┘
         ↓
┌─────────────────┐
│ Product Manager │  ← /product-manager
└────────┬────────┘
         ↓ (PRD)
┌─────────────────┐
│   Architect     │  ← /architect
└────────┬────────┘
         ↓ (ADR)
┌─────────────────┐
│  Feature Lead   │  ← /feature-lead
└────────┬────────┘
         ↓ (Tasks)
┌─────────────────────────────────────┐
│        Specialized Developers       │
├──────────┬──────────┬──────────┬────┤
│   API    │   Infra  │  Backend │ FE │
│ /dev-api │ /dev-inf │ /dev-bac │/dev│
└──────────┴──────────┴──────────┴────┘
         ↓
   Working Prototype
```

## Available Commands

### 🎯 Orchestrator

**Command**: `/product-request`

The main entry point that coordinates the entire team. Takes a customer feature request and manages the full development lifecycle:

1. Creates PRD via Product Manager
2. Generates ADR via Architect
3. Breaks down tasks via Feature Lead
4. Coordinates specialized developers
5. Delivers working prototype

**Usage**:
```
/product-request I need a user authentication system with email/password login
```

### 📋 Product Manager

**Command**: `/product-manager`

Transforms customer requests into comprehensive Product Requirement Documents (PRDs).

**Outputs**:
- Executive summary
- Problem statement
- User stories with acceptance criteria
- Functional and non-functional requirements
- Success metrics
- Dependencies and constraints

**Usage** (typically invoked by orchestrator):
```
/product-manager [customer request]
```

### 🏗️ System Architect

**Command**: `/architect`

Creates Architecture Decision Records (ADRs) based on PRDs. Designs system architecture aligned with the project's tech stack.

**Outputs**:
- Architecture diagrams (Mermaid)
- Component design (Frontend, Backend, API, Infrastructure)
- Database schema
- API contracts (Protobuf)
- Implementation phases
- Trade-offs and alternatives

**Usage** (typically invoked by orchestrator):
```
/architect [PRD content]
```

### 📊 Feature Lead

**Command**: `/feature-lead`

Breaks down features into deliverable milestones and actionable tasks for developers.

**Outputs**:
- Delivery milestones (Foundation → Backend → Frontend → Integration)
- User story breakdown
- Task database in Notion
- Task assignments by specialist
- Dependency graph
- Risk assessment

**Usage** (typically invoked by orchestrator):
```
/feature-lead [PRD and ADR content]
```

### 👨‍💻 Specialized Developers

#### API Designer
**Command**: `/dev-api`

Designs and implements Protobuf schemas and API contracts.

**Responsibilities**:
- Create `.proto` files
- Define messages, services, and RPCs
- Generate Go and TypeScript SDKs
- Document API contracts

**Tech**: Protobuf, Buf, Connect RPC

#### Infrastructure Engineer
**Command**: `/dev-infrastructure`

Sets up database schemas, migrations, and Docker services.

**Responsibilities**:
- Create database migrations (Goose)
- Write SQL queries
- Configure Docker services
- Manage environment variables

**Tech**: PostgreSQL, Docker Compose, Goose

#### Backend Developer
**Command**: `/dev-backend`

Implements backend services and business logic.

**Responsibilities**:
- Implement Connect RPC handlers
- Write domain logic
- Create database access layer
- Write unit and integration tests

**Tech**: Go, Gin, Connect RPC, PostgreSQL

#### Frontend Developer
**Command**: `/dev-frontend`

Builds user interfaces and integrates with backend APIs.

**Responsibilities**:
- Create React components
- Implement pages and routing
- Integrate with APIs via Connect Web
- Write component tests

**Tech**: React, TypeScript, Mantine UI, React Query

## Notion Integration

All documentation and tasks are stored in Notion:

- **PRDs**: Formatted pages with all requirement details
- **ADRs**: Technical design documents with diagrams
- **Task Database**: Structured database with:
  - Task ID, Title, Description
  - Phase (Foundation / Backend / Frontend / Integration)
  - Specialist (API / Infrastructure / Backend / Frontend)
  - Status (Not Started / In Progress / Review / Done)
  - Effort (S / M / L / XL)
  - Dependencies

## Quick Start

### 1. Full Feature Implementation

Use the orchestrator to implement a complete feature:

```bash
/product-request Add a dashboard that shows system metrics in real-time
```

The orchestrator will:
- Create PRD, ADR, and task breakdown in Notion
- Coordinate all developers to implement the feature
- Test the implementation end-to-end
- Provide you with links to all artifacts

### 2. Individual Agent Usage

You can also invoke agents directly:

```bash
# Create just a PRD
/product-manager I need analytics tracking for user behavior

# Create just an ADR from existing PRD
/architect [paste PRD here]

# Break down existing PRD and ADR into tasks
/feature-lead [paste PRD and ADR here]

# Implement specific development tasks
/dev-api Create a metrics API service
/dev-infrastructure Set up Redis for caching
/dev-backend Implement authentication logic
/dev-frontend Build the login page
```

## Workflow Example

Let's say you want to add a user profile feature:

### Step 1: Request the Feature
```
/product-request Users should be able to view and edit their profile information
```

### Step 2: Orchestrator Coordinates

The orchestrator automatically:

1. **Creates Notion workspace** for the feature
2. **Invokes Product Manager** who creates a PRD with:
   - User stories (view profile, edit name/email, upload avatar)
   - Acceptance criteria
   - Success metrics
3. **Invokes Architect** who creates an ADR with:
   - Database schema (users table with profile fields)
   - API design (GetProfile, UpdateProfile RPCs)
   - Frontend components (ProfilePage, ProfileForm)
   - Implementation phases
4. **Invokes Feature Lead** who creates:
   - Phase 1 tasks: Database migration, Protobuf schemas
   - Phase 2 tasks: Backend API implementation
   - Phase 3 tasks: Frontend UI components
   - Phase 4 tasks: Integration testing

### Step 3: Development

The orchestrator coordinates developers in phases:

**Phase 1 - Foundation** (parallel):
- API Designer: Creates user profile Protobuf schemas
- Infrastructure: Creates user_profiles table migration

**Phase 2 - Backend** (sequential):
- Backend Developer: Implements GetProfile and UpdateProfile handlers

**Phase 3 - Frontend** (sequential):
- Frontend Developer: Creates ProfilePage and ProfileForm components

**Phase 4 - Integration**:
- All developers: Fix integration issues, add tests

### Step 4: Delivery

You receive:
- Notion page with complete documentation
- Working implementation in the codebase
- Instructions to test the feature
- Summary of what was built

## Best Practices

### When to Use the Orchestrator

Use `/product-request` when:
- You have a complete feature request
- You want end-to-end implementation
- You want comprehensive documentation

### When to Use Individual Agents

Use individual commands when:
- You only need documentation (PRD or ADR)
- You want to work on specific parts manually
- You're iterating on an existing feature
- You need specialized expertise (e.g., just database work)

### Review Opportunities

While the orchestrator runs automatically, you can review:
- **After PRD**: Check if requirements are correct
- **After ADR**: Verify architecture decisions
- **After Task Breakdown**: Confirm implementation approach
- **During Development**: Review code as it's created
- **Final Output**: Test the working prototype

## Configuration

### Notion Setup

The agents use the MCP Notion integration. Ensure:
1. Notion integration is configured in Claude Code
2. You have a workspace where pages can be created
3. The agents will create a "Product Engineering Workspace" parent page

### Repository Structure

The agents are familiar with:
- Protobuf schemas: `_schema/protos/`
- Backend code: `platform/backend/`
- Frontend code: `platform/frontend/`
- Infrastructure: `_harness/docker-compose.yml`
- Generated APIs: `api/go-sdk/` and `api/web-sdk/`

## Troubleshooting

### Orchestrator Can't Find Agents

Make sure all command files exist in `.claude/commands/`:
- `product-request.md`
- `product-manager.md`
- `architect.md`
- `feature-lead.md`
- `dev-api.md`
- `dev-infrastructure.md`
- `dev-backend.md`
- `dev-frontend.md`

### Notion Integration Issues

Verify:
- MCP Notion integration is running
- You have permissions to create pages
- The Notion workspace is accessible

### Implementation Issues

If implementations fail:
- Check the task breakdown for correct dependencies
- Ensure infrastructure is running (`npm run harness start`)
- Verify generated code is up to date (`npm run schema codegen`)

## Extending the Team

You can add new agents by creating `.claude/commands/[agent-name].md` files:

1. Define the agent's role and expertise
2. Specify input and output formats
3. Provide implementation guidelines
4. Update the orchestrator to invoke the new agent

Example agents you might add:
- **QA Engineer**: Automated testing and quality assurance
- **DevOps Engineer**: CI/CD and deployment
- **Security Engineer**: Security reviews and audits
- **Technical Writer**: User documentation

## Credits

This AI engineering team system was designed to demonstrate:
- Multi-agent orchestration with Claude Code
- Integration with Notion for documentation
- Automated end-to-end feature development
- Specialized AI personas for different roles

Feel free to customize the agents to match your workflow and preferences!
