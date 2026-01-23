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

## Development with OpenSpec

This project uses **OpenSpec** for all changes. Before implementing a feature:

1. Create a proposal in `openspec/changes/`.
2. Define requirements and scenarios.
3. Once approved, implement following the tasks list.
4. Archive the change to update the living specifications in `openspec/specs/`.

---
Created as a showcase for modern Go desktop application development.
