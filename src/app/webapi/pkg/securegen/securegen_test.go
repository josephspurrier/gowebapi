package securegen_test

import (
	"app/webapi/pkg/securegen"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUUID(t *testing.T) {
	iterations := 50000

	m := make(map[string]bool)

	for i := 0; i < iterations; i++ {
		s, err := securegen.UUID()
		assert.Nil(t, err)

		// Ensure the lengths are consistent.
		arr := strings.Split(s, "-")
		assert.Equal(t, 5, len(arr))
		assert.Equal(t, len(arr[0]), 8)
		assert.Equal(t, len(arr[1]), 4)
		assert.Equal(t, len(arr[2]), 4)
		assert.Equal(t, len(arr[3]), 4)
		assert.Equal(t, len(arr[4]), 12)

		m[s] = false
	}

	// Ensure the randomness is accurate. Every entry should be unique making
	// the len the same number as the iterations.
	assert.Len(t, m, iterations)
}

func TestBytes(t *testing.T) {
	iterations := 50000

	m := make(map[string]bool)

	for i := 0; i < iterations; i++ {
		s, err := securegen.Bytes(32)
		assert.Nil(t, err)

		// Ensure the lengths are consistent.
		assert.Equal(t, 32, len(s))

		m[string(s)] = false
	}

	// Ensure the randomness is accurate. Every entry should be unique making
	// the len the same number as the iterations.
	assert.Len(t, m, iterations)
}
