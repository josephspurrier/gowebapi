package webtoken

import "time"

// IClock represents a system clock.
type IClock interface {
	Now() time.Time
}

// clock is the standard system clock.
type clock struct{}

// Now returns the current time.
func (c *clock) Now() time.Time {
	return time.Now()
}
