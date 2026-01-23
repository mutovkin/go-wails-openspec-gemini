// Package services implements the core business logic input ports.
package services

import (
	"fmt"
	"lottery-picker/internal/core/domain"
	"lottery-picker/internal/core/ports"
)

// LotteryService implements the TicketGenerator interface.
type LotteryService struct {
	rng ports.RandomProvider
}

// NewLotteryService creates a new instance of LotteryService.
func NewLotteryService(rng ports.RandomProvider) *LotteryService {
	return &LotteryService{rng: rng}
}

// GenerateTicket creates a new valid lottery ticket using the random provider.
func (s *LotteryService) GenerateTicket() (*domain.Ticket, error) {
	numbers, err := s.rng.GetSixUniqueNumbers(domain.MinNumber, domain.MaxNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to generate random numbers: %w", err)
	}

	ticket, err := domain.NewTicket(numbers)
	if err != nil {
		return nil, fmt.Errorf("failed to create ticket: %w", err)
	}

	return ticket, nil
}
