## ADDED Requirements

### Requirement: Hexagonal Directory Structure
The application source code MUST be organized according to Hexagonal Architecture principles to separate concerns.

#### Scenario: Verify Directory Layout
- **WHEN** the project is initialized
- **THEN** `internal/core/domain` exists for business logic
- **AND** `internal/core/ports` exists for interfaces
- **AND** `internal/adapters` exists for implementation details
- **AND** `cmd/app` exists as the entry point
