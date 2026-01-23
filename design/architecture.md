# Architecture Design: 6/49 Lottery Picker

## Overview

This project follows the **Hexagonal Architecture** (also known as Ports and Adapters) to create a maintainable, testable, and loosely coupled application.

## Core Concepts

### 1. Domain (The Hexagon)

The core business logic resides here. It knows nothing about the outside world (database, UI, HTTP, etc.).

- **Entities**: `Ticket` (a set of 6 numbers), `LotteryRules` (validation constraints).
- **Responsibility**: Validation, Business Rules (e.g., "Numbers must be unique", "Range 1-49").

### 2. Ports

Interfaces that define how the application interacts with the outside world.

- **Input Ports (Driving)**: API for the UI to request actions (e.g., `GenerateTicket()`).
- **Output Ports (Driven)**: Interfaces for infrastructure needs (e.g., `RandomNumberGenerator`).

### 3. Adapters

Concrete implementations of the ports.

- **Primary Adapter**: The Wails Application struct. It receives events from the GUI (JS) and calls the Input Port (Service).
- **Secondary Adapter**: The implementation of `RandomNumberGenerator`. We can have a `CryptoRandomAdapter` for production and a `MockRandomAdapter` for testing.

## Directory Structure Strategy

We align the file structure with the architecture:

```
internal/
  core/
    domain/      # Structs: Ticket
    ports/       # Interfaces: LotteryService (In), RandomProvider (Out)
    services/    # Logic: LotteryServiceImpl (Implements LotteryService)
  adapters/
    primary/
      wailsapp/  # Wails App struct (calls LotteryService)
    secondary/
      random/    # Implements RandomProvider
```

## Data Flow

1. **User** clicks "Generate" in Vue GUI.
2. **Vue** calls Wails Runtime (`Go` method).
3. **Wails Adapter** (`App.GenerateTicket`) is triggered.
4. **Wails Adapter** calls **Service** (`LotteryService.CreateTicket`).
5. **Service** calls **Random Port** (`RandomProvider.GetNumbers`).
6. **Random Adapter** generates numbers and returns them.
7. **Service** validates/sorts logic via **Domain** entities.
8. **Service** returns `Ticket` struct.
9. **Wails Adapter** returns JSON to **Vue**.
10. **Vue** updates the display.
