package domain_test

import (
	"lottery-picker/internal/core/domain"
	"testing"
)

func TestNewTicket_Valid(t *testing.T) {
	t.Parallel()

	nums := []int{1, 2, 3, 4, 5, 6}

	ticket, err := domain.NewTicket(nums)
	if err != nil {
		t.Fatalf("expected valid ticket, got error: %v", err)
	}

	if len(ticket.Numbers) != 6 {
		t.Errorf("expected 6 numbers, got %d", len(ticket.Numbers))
	}
}

func TestNewTicket_Sorting(t *testing.T) {
	t.Parallel()

	nums := []int{6, 5, 4, 3, 2, 1}

	ticket, err := domain.NewTicket(nums)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []int{1, 2, 3, 4, 5, 6}
	for i, n := range ticket.Numbers {
		if n != expected[i] {
			t.Errorf("expected %d at index %d, got %d", expected[i], i, n)
		}
	}
}

func TestNewTicket_InvalidLength(t *testing.T) {
	t.Parallel()

	nums := []int{1, 2, 3, 4, 5}

	_, err := domain.NewTicket(nums)
	if err == nil {
		t.Error("expected error for 5 numbers, got nil")
	}
}

func TestNewTicket_OutOfRange(t *testing.T) {
	t.Parallel()

	nums := []int{0, 2, 3, 4, 5, 6}

	_, err := domain.NewTicket(nums)
	if err == nil {
		t.Error("expected error for number < 1, got nil")
	}

	nums = []int{1, 2, 3, 4, 5, 50}

	_, err = domain.NewTicket(nums)
	if err == nil {
		t.Error("expected error for number > 49, got nil")
	}
}

func TestNewTicket_Duplicates(t *testing.T) {
	t.Parallel()

	nums := []int{1, 2, 3, 4, 5, 5}

	_, err := domain.NewTicket(nums)
	if err == nil {
		t.Error("expected error for duplicate numbers, got nil")
	}
}
