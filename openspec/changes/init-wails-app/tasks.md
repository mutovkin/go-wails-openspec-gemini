## 1. Project Initialization
- [x] 1.1 Install Wails (if not present) and scaffold project using `wails init -n lottery-picker -t vue`.
- [x] 1.2 Move scaffolded files to match the directory structure defined in `project.md`.
- [x] 1.3 Configure `wails.json` to point to the correct frontend and main entry points.

## 2. Core Domain & Ports (TDD)
- [x] 2.1 Define `Ticket` entity in `internal/core/domain/ticket.go` with validation logic (unique numbers, range 1-49).
- [x] 2.2 Write unit tests for `Ticket` validation.
- [x] 2.3 Define `RandomProvider` interface in `internal/core/ports/output.go`.
- [x] 2.4 Define `TicketGenerator` interface in `internal/core/ports/input.go`.

## 3. Services (TDD)
- [x] 3.1 Implement `LotteryService` in `internal/core/services/lottery.go`.
- [x] 3.2 Write unit tests for `LotteryService` using a mock `RandomProvider`.

## 4. Adapters
- [x] 4.1 Implement `CryptoRandomAdapter` in `internal/adapters/secondary/random/crypto.go`.
- [x] 4.2 Refactor default Wails `App` struct in `internal/adapters/primary/wailsapp/app.go` to use `TicketGenerator`.
- [x] 4.3 Wire up dependencies in `cmd/app/main.go`.

## 5. Frontend Implementation
- [x] 5.1 Clean up default Wails frontend code.
- [x] 5.2 Create a simple Vue component to trigger `GenerateTicket` and display results.
- [x] 5.3 Apply basic Material Design 3 styling (CSS or library).

## 6. Verification
- [x] 6.1 Run `golangci-lint` and fix violations.
- [x] 6.2 Run all tests via `make test`.
- [x] 6.3 Verify manual build via `make build`.