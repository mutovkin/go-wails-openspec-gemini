# Project Context

## Purpose
A modern GUI application for generating random 6/49 lottery tickets. The application aims to provide a simple, aesthetically pleasing user experience using Material Design 3 principles, while serving as a robust example of Hexagonal Architecture in Go.

## Tech Stack
- **Language:** Go (Latest Stable)
- **GUI Framework:** Wails.io (v2)
- **Frontend:** Vue.js (latest, with Composition API)
- **UI System:** Material Design 3 (e.g., via Vuetify 3 or similar MD3 library, or custom CSS)
- **Linting:** golangci-lint (version 2.8 requested, or latest stable compatible version)
- **Build Tool:** Make

## Project Conventions

### Architecture: Hexagonal (Ports & Adapters)
The application strictly follows Hexagonal Architecture to decouple business logic from the GUI and infrastructure.

- **Domain (`internal/core/domain`)**: Contains pure business entities and logic (e.g., `Ticket`, `NumberSet`, `Rules`). No dependencies on external libraries.
- **Ports (`internal/core/ports`)**: Defines interfaces for driving adapters (incoming) and driven adapters (outgoing).
    - *Input Port*: Service interface used by the GUI (e.g., `TicketGenerator`).
    - *Output Port*: Interfaces for randomness (e.g., `RandomProvider`).
- **Services (`internal/core/services`)**: Implements the input ports/use cases (e.g., `LotteryService`).
- **Adapters (`internal/adapters`)**:
    - *Primary (Driving)*: The Wails `App` struct which binds to the frontend and calls the Service.
    - *Secondary (Driven)*: Implementations for `RandomProvider` (e.g., `CryptoRandomAdapter`, `MathRandomAdapter`).

### Project Structure
```text
/
├── cmd/
│   └── app/                # Main entry point
├── frontend/               # Vue.js frontend (Wails standard)
├── internal/
│   ├── core/
│   │   ├── domain/         # Business entities
│   │   ├── ports/          # Interfaces (Input/Output)
│   │   └── services/       # Application logic
│   └── adapters/
│       ├── primary/        # Driving adapters
│       │   └── wailsapp/   # Wails App binding
│       └── secondary/      # Driven adapters
│           └── random/     # RNG implementation
├── design/                 # Design documentation
│   └── uml/                # PlantUML diagrams
├── build/                  # Wails build artifacts
├── wails.json              # Wails configuration
├── Makefile                # Developer convenience scripts
└── go.mod
```

### Code Style
- **Go**: Follows standard Go idioms (`gofmt`).
    - Error handling: Explicit error checking.
    - Naming: CamelCase, short but descriptive.
- **Frontend**: Vue Composition API with `<script setup>`.
- **Linting**: Strict linting using `golangci-lint` with configuration compatible with v2.8 requirements.

### Testing Strategy
- **Methodology**: Test-Driven Development (TDD).
- **Unit Tests**:
    - Core logic (`services`, `domain`) must be 100% covered.
    - Mocks used for Output Ports (e.g., mocking the Random Provider to ensure deterministic tests for ticket generation).
- **Integration Tests**:
    - Adapter integration tests where feasible.
- **Frontend Tests**: Basic component testing for UI interaction.
- **Verification**: Run `make fmt` and `make lint` after completing each OpenSpec task to ensure code quality and formatting compliance.

### Git Workflow & Release Process
- **Git Flow**:
    - `develop`: Default branch for PRs.
    - `master`: Stable release branch (read-only, updated via automation).
    - **Feature Branches**: `feature/name` -> Merge to `develop` (Squash).
    - **Hotfix Branches**: `hotfix/name` -> Merge to `master` (Squash) -> Auto-rebase `develop`.
- **Commits**: Follow Conventional Commits (e.g., `feat:`, `fix:`, `docs:`, `chore:`).
- **Rebasing**: Feature branches must be rebased on `develop` before merging.
- **Releases**:
    - Triggered manually via GitHub Actions (`workflow_dispatch`).
    - Automated tagging and fast-forward merge from `develop` to `master`.
    - Automated artifact build and GitHub Release creation.

### Repository Configuration
- **Branch Protection**:
    - `master` and `develop`: Require Pull Request reviews.
    - **Status Checks**: CI (Lint, Test, Build) must pass before merging.
    - **Linear History**: Required.

## Domain Context
- **6/49 Lottery**: A standard lottery game where 6 distinct numbers are drawn from a range of 1 to 49.
- **Ticket**: A set of 6 unique numbers.
- **Randomness**: Critical for lottery applications. The solution must support cryptographically secure random number generation.

## Tasks & Roadmap
1.  **Design**: Create PlantUML diagrams for Hexagonal structure.
2.  **Setup**: Initialize Wails project and restructure for Hexagonal.
3.  **Implementation**:
    - Define Domain & Ports (TDD).
    - Implement Services (TDD).
    - Implement Adapters (RNG, Wails App).
    - Build Frontend (Vue + MD3).
4.  **Tooling**: Create `Makefile` for `build`, `test`, `lint`, `run`.

## Important Constraints
- **Validation**: Generated numbers must be unique within a ticket and strictly within 1-49 range.
- **UI**: Must be simple, modern, and follow Material Design 3.
