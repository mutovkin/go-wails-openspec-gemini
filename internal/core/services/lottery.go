package services

import (
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
	numbers, err := s.rng.GetSixUniqueNumbers(1, 49)
	if err != nil {
		return nil, err
	}
	return domain.NewTicket(numbers)
}
