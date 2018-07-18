package passhash_test

import (
	"testing"

	"app/webapi/pkg/passhash"

	"github.com/stretchr/testify/assert"
)

// TestStringString tests string to string hash.
func TestStringString(t *testing.T) {
	ph := passhash.New()
	plainText := "This is a test."
	hash, err := ph.HashString(plainText)
	assert.Nil(t, err)
	assert.True(t, ph.MatchString(hash, plainText))

	plainText2 := "This is a test2."
	hash, err = ph.HashString(plainText2)
	assert.Nil(t, err)
	assert.False(t, ph.MatchString(hash, plainText))
}

// TestByteByte tests byte to byte hash.
func TestByteByte(t *testing.T) {
	ph := passhash.New()
	plainText := []byte("This is a test.")
	hash, err := ph.HashBytes(plainText)
	assert.Nil(t, err)
	assert.True(t, ph.MatchBytes(hash, plainText))

	plainText2 := []byte("This is a test2.")
	hash, err = ph.HashBytes(plainText2)
	assert.Nil(t, err)
	assert.False(t, ph.MatchBytes(hash, plainText))
}

// TestStringByte tests string to byte hash.
func TestStringByte(t *testing.T) {
	ph := passhash.New()
	plainText := "This is a test."
	hash, err := ph.HashString(plainText)
	assert.Nil(t, err)
	assert.True(t, ph.MatchBytes([]byte(hash), []byte(plainText)))
}

// TestByteString tests byte to string hash.
func TestByteString(t *testing.T) {
	ph := passhash.New()
	plainText := []byte("This is a test.")
	hash, err := ph.HashBytes(plainText)
	assert.Nil(t, err)
	assert.True(t, ph.MatchString(string(hash), string(plainText)))
}

// TestHashStringEmpty tests empty string which should pass fine.
func TestHashStringEmpty(t *testing.T) {
	ph := passhash.New()
	plainText := ""
	hash, err := ph.HashString(plainText)
	assert.Nil(t, err)
	assert.True(t, ph.MatchString(hash, plainText))
}
