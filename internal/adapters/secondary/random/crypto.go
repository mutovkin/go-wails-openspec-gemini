package random

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// CryptoRandomAdapter implements ports.RandomProvider using crypto/rand.
type CryptoRandomAdapter struct{}

// NewCryptoRandomAdapter creates a new instance.
func NewCryptoRandomAdapter() *CryptoRandomAdapter {
	return &CryptoRandomAdapter{}
}

// GetSixUniqueNumbers generates 6 unique integers between min and max.
func (c *CryptoRandomAdapter) GetSixUniqueNumbers(min, max int) ([]int, error) {
	if max-min+1 < 6 {
		return nil, fmt.Errorf("range [%d, %d] is too small to generate 6 unique numbers", min, max)
	}

	seen := make(map[int]bool)
	var numbers []int

	for len(numbers) < 6 {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
		if err != nil {
			return nil, fmt.Errorf("crypto/rand error: %w", err)
		}
		val := int(n.Int64()) + min

		if !seen[val] {
			seen[val] = true
			numbers = append(numbers, val)
		}
	}

	return numbers, nil
}
