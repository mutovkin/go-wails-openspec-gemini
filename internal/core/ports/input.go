package ports

import "lottery-picker/internal/core/domain"

// TicketGenerator defines the use case for generating lottery tickets.
type TicketGenerator interface {
	// GenerateTicket creates a new valid lottery ticket.
	GenerateTicket() (*domain.Ticket, error)
}
