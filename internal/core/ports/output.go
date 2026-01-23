package ports

// RandomProvider defines the interface for generating random numbers.
type RandomProvider interface {
	// GetSixUniqueNumbers returns 6 unique integers between min and max (inclusive).
	GetSixUniqueNumbers(min, max int) ([]int, error)
}
