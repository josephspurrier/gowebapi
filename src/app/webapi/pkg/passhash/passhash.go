// Package passhash provides password hashing functionality using bcrypt.
package passhash

import (
	"golang.org/x/crypto/bcrypt"
)

// New returns a password hashing tool.
func New() *Passhash {
	return &Passhash{}
}

// Passhash is a password hashing tool.
type Passhash struct{}

// HashString returns a hashed string and an error.
func (p *Passhash) HashString(password string) (string, error) {
	key, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(key), nil
}

// HashBytes returns a hashed byte array and an error.
func (p *Passhash) HashBytes(password []byte) ([]byte, error) {
	key, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return key, nil
}

// MatchString returns true if the hash matches the password.
func (p *Passhash) MatchString(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == nil {
		return true
	}

	return false
}

// MatchBytes returns true if the hash matches the password.
func (p *Passhash) MatchBytes(hash, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err == nil {
		return true
	}

	return false
}
