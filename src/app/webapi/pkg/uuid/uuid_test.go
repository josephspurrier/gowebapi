package uuid_test

import (
	"strings"
	"testing"

	"app/webapi/pkg/uuid"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	iterations := 500000

	m := make(map[string]bool)

	for i := 0; i < iterations; i++ {
		s, err := uuid.Generate()
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
