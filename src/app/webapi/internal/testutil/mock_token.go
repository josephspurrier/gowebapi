package testutil

import "time"

// MockToken is a mocked webtoken.
type MockToken struct{}

type generateFunc func(userID string, duration time.Duration) (string, error)

var (
	generate generateFunc

	// GenerateDefault returns "", nil.
	GenerateDefault = func(userID string, duration time.Duration) (string, error) {
		return "", nil
	}
)

// SetGenerate will set the function.
func (mt *MockToken) SetGenerate(fn generateFunc) {
	generate = fn
}

// Generate .
func (mt *MockToken) Generate(userID string, duration time.Duration) (string, error) {
	// Use the default.
	fnInternal := generate
	if fnInternal == nil {
		fnInternal = GenerateDefault
	}
	return fnInternal(userID, duration)
}
