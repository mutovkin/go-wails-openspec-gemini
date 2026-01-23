## Context
We are building a GUI application for lottery ticket generation. The requirement is to use Hexagonal Architecture to ensure the business logic (ticket rules) is independent of the UI framework (Wails).

## Goals / Non-Goals
- **Goals**:
    - Strict dependency rule: Domain depends on nothing.
    - Testability: Core logic testable without GUI.
    - Modern UI: MD3 look and feel.
- **Non-Goals**:
    - Persistent storage (database) is not required for this iteration.
    - Complex animation is secondary to architecture correctness.

## Decisions
- **Wails as Primary Adapter**: The Wails `App` struct acts as the entry point from the UI. It transforms JSON-serializable requests from the frontend into Go method calls on the Service layer.
- **Crypto RNG**: We use `crypto/rand` for high-quality randomness, wrapped in an adapter to satisfy the `RandomProvider` interface. This allows us to swap it for a deterministic generator during tests.
- **Vue Composition API**: Chosen for better TypeScript support (if needed later) and cleaner component logic organization.

## Risks / Trade-offs
- **Complexity**: Hexagonal architecture adds boilerplate (ports, adapters) for a simple app.
    - *Mitigation*: The project structure is clearly defined in `project.md` to prevent confusion.
- **Performance**: Reflection/Serialization overhead in Wails.
    - *Mitigation*: Negligible for simple data structs like a ticket.

## Open Questions
- None.
