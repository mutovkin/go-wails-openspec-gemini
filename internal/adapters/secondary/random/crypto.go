// Package random provides random number generation adapters.
package random

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

// ErrRangeTooSmall is returned when the range is insufficient for the requested unique count.
var ErrRangeTooSmall = errors.New("range is too small to generate unique numbers")

// CryptoRandomAdapter implements ports.RandomProvider using crypto/rand.
type CryptoRandomAdapter struct{}

// NewCryptoRandomAdapter creates a new instance.
func NewCryptoRandomAdapter() *CryptoRandomAdapter {
	return &CryptoRandomAdapter{}
}

// GetSixUniqueNumbers generates 6 unique integers between min and max.
func (c *CryptoRandomAdapter) GetSixUniqueNumbers(minVal, maxVal int) ([]int, error) {
	const count = 6
	if maxVal-minVal+1 < count {
		return nil, fmt.Errorf(
			"%w: [%d, %d] for %d numbers",
			ErrRangeTooSmall,
			minVal,
			maxVal,
			count,
		)
	}

	seen := make(map[int]bool)

	var numbers []int

	for len(numbers) < count {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(maxVal-minVal+1)))
		if err != nil {
			return nil, fmt.Errorf("crypto/rand error: %w", err)
		}

		val := int(n.Int64()) + minVal

		if !seen[val] {
			seen[val] = true
			numbers = append(numbers, val)
		}
	}

	return numbers, nil
}
