// Package domain contains the core business entities and logic for the lottery application.
package domain

import (
	"errors"
	"fmt"
	"sort"
)

const (
	// TicketSize is the required number of integers in a ticket.
	TicketSize = 6
	// MinNumber is the minimum allowed value for a lottery number.
	MinNumber = 1
	// MaxNumber is the maximum allowed value for a lottery number.
	MaxNumber = 49
)

var (
	// ErrInvalidCount is returned when the ticket has the wrong number of values.
	ErrInvalidCount = errors.New("ticket must have exactly 6 numbers")
	// ErrOutOfRange is returned when a number is outside the 1-49 range.
	ErrOutOfRange = errors.New("number is out of range")
	// ErrDuplicate is returned when a duplicate number is found.
	ErrDuplicate = errors.New("duplicate number found")
)

// Ticket represents a set of 6/49 lottery numbers.
type Ticket struct {
	Numbers []int
}

// NewTicket creates a new Ticket and validates it.
func NewTicket(numbers []int) (*Ticket, error) {
	ticket := &Ticket{Numbers: numbers}
	if err := ticket.Validate(); err != nil {
		return nil, err
	}

	// Sort numbers for consistent display/storage
	sort.Ints(ticket.Numbers)

	return ticket, nil
}

// Validate ensures the ticket follows 6/49 rules.
func (ticket *Ticket) Validate() error {
	if len(ticket.Numbers) != TicketSize {
		return fmt.Errorf("%w: got %d", ErrInvalidCount, len(ticket.Numbers))
	}

	seen := make(map[int]bool)

	for _, number := range ticket.Numbers {
		if number < MinNumber || number > MaxNumber {
			return fmt.Errorf("%w: %d (1-49)", ErrOutOfRange, number)
		}

		if seen[number] {
			return fmt.Errorf("%w: %d", ErrDuplicate, number)
		}

		seen[number] = true
	}

	return nil
}
