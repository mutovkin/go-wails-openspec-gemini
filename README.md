# 6/49 Lottery Picker

> **Note**: This project is a demo/experiment for **OpenSpec** + **Gemini-cli** + **Go** + **Wails**. It showcases spec-driven development using a Hexagonal Architecture (Ports and Adapters) in a modern desktop application environment.

## Overview

A modern GUI application for generating random 6/49 lottery tickets. The application provides a simple, aesthetically pleasing user experience using Material Design 3 principles while serving as a robust example of maintainable software architecture.

## Tech Stack

- **Backend**: [Go](https://go.dev/) (latest stable)
- **GUI Framework**: [Wails.io](https://wails.io/) (v2)
- **Frontend**: [Vue.js](https://vuejs.org/) (Composition API)
- **UI Styling**: [Vuetify 3](https://vuetifyjs.com/) (Material Design 3)
- **Architecture**: Hexagonal (Ports & Adapters)
- **Development Process**: [OpenSpec](https://openspec.dev/) (Spec-driven development)

## Architecture

The application strictly follows **Hexagonal Architecture** to decouple business logic from the GUI and infrastructure:

- **Domain (`internal/core/domain`)**: Pure business entities (`Ticket`) and validation rules.
- **Ports (`internal/core/ports`)**: Interfaces for driving (input) and driven (output) adapters.
- **Services (`internal/core/services`)**: Implementation of use cases and orchestration.
- **Adapters (`internal/adapters`)**:
  - **Primary**: Wails App binding connecting Vue.js to Go services.
  - **Secondary**: Cryptographically secure random number generation (`crypto/rand`).

For more details, see [design/architecture.md](design/architecture.md).

## Project Structure

```text
├── internal/
│   ├── core/
│   │   ├── domain/         # Business entities (Ticket)
│   │   ├── ports/          # Interfaces (Input/Output)
│   │   └── services/       # Application logic (LotteryService)
│   └── adapters/
│       ├── primary/        # Driving adapters (Wails App)
│       └── secondary/      # Driven adapters (RNG)
├── design/                 # Architecture & UML documentation
├── frontend/               # Vue.js + Vuetify frontend
├── openspec/               # Specifications and Change Proposals
├── main.go                 # Entry point
└── Makefile                # Developer workflow commands
```

## Domain Rules

- **6/49 Lottery**: 6 distinct numbers drawn from a pool of 1–49.
- Numbers within a ticket must be **unique** and strictly within range.
- Randomness must be **cryptographically secure** (`crypto/rand`).

## Code Style

- **Go**: Standard idioms enforced by `gofmt`. Explicit error handling; CamelCase naming.
- **Frontend**: Vue Composition API with `<script setup>`.
- **Linting**: `golangci-lint` (v2.9+) — strict configuration.
- **Commits**: [Conventional Commits](https://www.conventionalcommits.org/) (`feat:`, `fix:`, `docs:`, `chore:`).

## Testing Strategy

- **Methodology**: Test-Driven Development (TDD).
- **Coverage**: Core logic (`services/`, `domain/`) must be 100% covered.
- **Mocking**: Output Ports (e.g., `RandomProvider`) are mocked to ensure deterministic unit tests.
- **Integration**: Adapter integration tests where feasible.
- **Verification**: Run `make fmt` and `make lint` after changes.

## Getting Started

### Prerequisites

- [Go](https://go.dev/doc/install)
- [Node.js & npm](https://nodejs.org/)
- [Wails CLI](https://wails.io/docs/gettingstarted/installation)

### Installation & Development

1. **Install tools**:

   ```bash
   make install-tools
   ```

2. **Run in development mode**:

   ```bash
   make dev
   ```

3. **Build production binary**:

   ```bash
   make build
   ```

4. **Run tests**:

   ```bash
   make test
   ```

## Workflow & Releases

### Git Flow

We use a modified Git Flow:

- **`develop`**: Main integration branch. All features merge here via PR.
- **`master`**: Stable release branch. Updated only via release automation.
- **`feature/*`**: Feature branches (rebase on `develop` before merge).
- **`hotfix/*`**: Hotfix branches merge to `master` (squash) -> auto-rebase `develop`.

### Release Process

Releases are automated via GitHub Actions:

1. Go to **Actions** -> **Release**.
2. Run workflow on:
   - `develop` for Standard Release (Minor/Major).
   - `master` for Hotfix Release (Patch).
3. Select increment (Patch, Minor, Major).
4. The workflow will tag, merge/rebase, build artifacts, and publish a GitHub Release.

See [design/release-flow.md](design/release-flow.md) for details.

### Repository Configuration

- **Branch Protection** (`master`, `develop`):
  - Require Pull Request reviews.
  - Require status checks to pass (Lint, Test, Build).
  - Require linear history.

## Development with OpenSpec

This project uses **OpenSpec** for all changes. Before implementing a feature:

1. Create a proposal in `openspec/changes/`.
2. Define requirements and scenarios.
3. Once approved, implement following the tasks list.
4. Archive the change to update the living specifications in `openspec/specs/`.

---
Created as a showcase for modern Go desktop application development.
