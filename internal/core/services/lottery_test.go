package services_test

import (
	"errors"
	"lottery-picker/internal/core/services"
	"testing"
)

var errRNGFail = errors.New("rng fail")

// MockRandomProvider is a mock implementation of ports.RandomProvider.
type MockRandomProvider struct {
	Numbers []int
	Err     error
}

func (m *MockRandomProvider) GetSixUniqueNumbers(_, _ int) ([]int, error) {
	return m.Numbers, m.Err
}

func TestGenerateTicket_Success(t *testing.T) {
	t.Parallel()

	expectedNums := []int{10, 20, 30, 40, 45, 49}
	mockRNG := &MockRandomProvider{Numbers: expectedNums, Err: nil}
	service := services.NewLotteryService(mockRNG)

	ticket, err := service.GenerateTicket()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for i, n := range ticket.Numbers {
		if n != expectedNums[i] {
			t.Errorf("expected %d, got %d", expectedNums[i], n)
		}
	}
}

func TestGenerateTicket_RNGError(t *testing.T) {
	t.Parallel()

	mockRNG := &MockRandomProvider{Numbers: nil, Err: errRNGFail}
	service := services.NewLotteryService(mockRNG)

	_, err := service.GenerateTicket()
	if err == nil {
		t.Error("expected error from RNG, got nil")
	}
}

func TestGenerateTicket_DomainValidationFail(t *testing.T) {
	t.Parallel()

	// Mock returns invalid numbers (duplicates)
	invalidNums := []int{1, 1, 2, 3, 4, 5}
	mockRNG := &MockRandomProvider{Numbers: invalidNums, Err: nil}
	service := services.NewLotteryService(mockRNG)

	_, err := service.GenerateTicket()
	if err == nil {
		t.Error("expected domain validation error, got nil")
	}
}
