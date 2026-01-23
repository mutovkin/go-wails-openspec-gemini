# Change: Initialize Wails Application with Hexagonal Architecture

## Why
To establish the foundational codebase for the 6/49 Lottery Picker application, ensuring a clear separation of concerns from the start using Hexagonal Architecture. This setup is required to support the functional goal of generating lottery tickets while adhering to strict architectural and tooling constraints.

## What Changes
- **Initialize Wails Project**: Scaffold a standard Wails v2 application with Vue.js.
- **Restructure for Hexagonal Architecture**:
    - Create `internal/core/domain`, `internal/core/ports`, `internal/core/services`.
    - Create `internal/adapters/primary` and `internal/adapters/secondary`.
- **Implement Core Logic**: Define `Ticket` domain entity and `LotteryService`.
- **Implement Adapters**:
    - **Primary**: Connect Wails App to `LotteryService`.
    - **Secondary**: Implement `CryptoRandom` adapter for RNG.
- **Frontend Setup**: Initialize Vue.js with basic Material Design 3 styling.
- **Tooling**: Configure `Makefile` and `golangci-lint`.

## Impact
- **New Capabilities**: `project-structure`, `gui`, `randomness`.
- **New Files**: All initial source code in `cmd/`, `internal/`, `frontend/`.
- **Build System**: `Makefile` and `wails.json` will be the source of truth for building.
