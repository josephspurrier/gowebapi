package testutil

import "time"

// MockToken is a mocked webtoken.
type MockToken struct {
	GenerateFunc GenerateFuncType
}

// GenerateFuncType .
type GenerateFuncType func(userID string, duration time.Duration) (string, error)

// GenerateFuncDefault .
var GenerateFuncDefault = func(userID string, duration time.Duration) (string, error) {
	return "", nil
}

// Generate .
func (mt *MockToken) Generate(userID string, duration time.Duration) (string, error) {
	if mt.GenerateFunc != nil {
		return mt.GenerateFunc(userID, duration)
	}
	return GenerateFuncDefault(userID, duration)
}
