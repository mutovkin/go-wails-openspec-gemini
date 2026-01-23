// Package wailsapp provides the Primary Adapter for the Wails application.
package wailsapp

import (
	"context"
	"fmt"
	"lottery-picker/internal/core/domain"
	"lottery-picker/internal/core/ports"
)

// App struct acts as the Primary Adapter for Wails.
type App struct {
	// ctx is required by Wails to perform runtime operations.
	//nolint:containedctx // Wails stores context in the struct for runtime access.
	ctx     context.Context
	service ports.TicketGenerator
}

// NewApp creates a new App application struct with dependencies.
func NewApp(service ports.TicketGenerator) *App {
	return &App{
		service: service,
	}
}

// Startup is called when the app starts. The context is saved.
// Exported so main can register it.
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// GenerateTicket calls the core service to generate a new ticket.
// This method is exposed to the Frontend.
func (a *App) GenerateTicket() (*domain.Ticket, error) {
	ticket, err := a.service.GenerateTicket()
	if err != nil {
		return nil, fmt.Errorf("app: %w", err)
	}

	return ticket, nil
}
