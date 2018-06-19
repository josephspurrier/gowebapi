package securegen

import (
	"crypto/rand"
	"fmt"
)

// UUID generates a UUID for use as an ID.
func UUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}

// Bytes return a []byte of random values.
func Bytes(length int) ([]byte, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	return b, err
}
