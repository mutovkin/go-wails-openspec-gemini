package domain

import (
	"fmt"
	"sort"
)

// Ticket represents a set of 6/49 lottery numbers.
type Ticket struct {
	Numbers []int
}

// NewTicket creates a new Ticket and validates it.
func NewTicket(numbers []int) (*Ticket, error) {
	t := &Ticket{Numbers: numbers}
	if err := t.Validate(); err != nil {
		return nil, err
	}
	// Sort numbers for consistent display/storage
	sort.Ints(t.Numbers)
	return t, nil
}

// Validate ensures the ticket follows 6/49 rules.
func (t *Ticket) Validate() error {
	if len(t.Numbers) != 6 {
		return fmt.Errorf("ticket must have exactly 6 numbers, got %d", len(t.Numbers))
	}

	seen := make(map[int]bool)
	for _, n := range t.Numbers {
		if n < 1 || n > 49 {
			return fmt.Errorf("number %d is out of range (1-49)", n)
		}
		if seen[n] {
			return fmt.Errorf("duplicate number %d found", n)
		}
		seen[n] = true
	}

	return nil
}
