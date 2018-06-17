package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetDefaults(t *testing.T) {
	c := Connection{}
	c = c.setDefaults()

	assert.Equal(t, "utf8mb4", c.Charset)
	assert.Equal(t, "utf8mb4_unicode_ci", c.Collation)
}
