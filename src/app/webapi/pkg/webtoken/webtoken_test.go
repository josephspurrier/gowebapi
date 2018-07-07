package webtoken_test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"
	"time"

	"app/webapi/pkg/webtoken"

	"github.com/stretchr/testify/assert"
)

type NowFn func() time.Time

type MockClock struct {
	nowfn NowFn
}

func (c *MockClock) SetNow(fn NowFn) {
	c.nowfn = fn
}

func (c *MockClock) Now() time.Time {
	if c.nowfn == nil {
		return time.Now()
	}
	return c.nowfn()
}

func TestValidJWT(t *testing.T) {
	mc := new(MockClock)

	secret := []byte("0123456789ABCDEF0123456789ABCDEF")

	// Generate a token.
	token := webtoken.New(secret)
	token.SetClock(mc)
	ss, err := token.Generate("jsmith", 999999*time.Hour)
	assert.Nil(t, err)
	assert.NotEmpty(t, ss)

	// Verify the token.
	s, err := token.Verify(ss)
	assert.Nil(t, err)
	assert.Equal(t, "jsmith", s)
}

func TestInvalidSecret(t *testing.T) {
	mc := new(MockClock)

	secret := []byte("0123456789ABCDEF0123456789ABCDEF")
	secret2 := []byte("0123456789ABCDEF0123456789ABCDEF3")

	// Generate a token.
	token := webtoken.New(secret)
	token.SetClock(mc)
	ss, err := token.Generate("jsmith", 999999*time.Hour)
	assert.Nil(t, err)
	assert.NotEmpty(t, ss)

	// Verify the token.
	token2 := webtoken.New(secret2)
	token.SetClock(mc)
	s, err := token2.Verify(ss)
	assert.Equal(t, webtoken.ErrSignatureInvalid, err)
	assert.Equal(t, "", s)
}

func TestNoSecret(t *testing.T) {
	mc := new(MockClock)

	// Generate a token.
	token := webtoken.New([]byte(""))
	token.SetClock(mc)
	ss, err := token.Generate("jsmith", 999999*time.Hour)
	assert.Equal(t, webtoken.ErrSecretTooShort, err)
	assert.Equal(t, "", ss)

	// Verify the token.
	s, err := token.Verify(ss)
	assert.Equal(t, webtoken.ErrMalformed, err)
	assert.Equal(t, "", s)
}

func TestFutureJWT(t *testing.T) {
	mc := new(MockClock)

	// Set the clock in the future.
	mc.SetNow(func() time.Time {
		return time.Now().Add(5 * time.Minute)
	})

	secret := []byte("0123456789ABCDEF0123456789ABCDEF")

	// Generate a token.
	token := webtoken.New(secret)
	token.SetClock(mc)
	ss, err := token.Generate("jsmith", 24*time.Hour)
	assert.Nil(t, err)
	assert.NotEmpty(t, ss)

	// Set the the clock for now.
	mc.SetNow(func() time.Time {
		return time.Now()
	})

	// Verify the token.
	s, err := token.Verify(ss)
	assert.Equal(t, webtoken.ErrNotValidYet, err)
	assert.Equal(t, "", s)
}

func TestPastJWT(t *testing.T) {
	mc := new(MockClock)

	// Set the clock in the past.
	mc.SetNow(func() time.Time {
		return time.Now().Add(-5 * time.Minute)
	})

	secret := []byte("0123456789ABCDEF0123456789ABCDEF")

	// Generate a token.
	token := webtoken.New(secret)
	token.SetClock(mc)
	ss, err := token.Generate("jsmith", 1*time.Minute)
	assert.Nil(t, err)
	assert.NotEmpty(t, ss)

	// Set the the clock for now.
	mc.SetNow(func() time.Time {
		return time.Now()
	})

	// Verify the token.
	s, err := token.Verify(ss)
	assert.Equal(t, webtoken.ErrExpired, err)
	assert.Equal(t, "", s)
}

func TestErrorJWT(t *testing.T) {
	mc := new(MockClock)

	secret := []byte("0123456789ABCDEF0123456789ABCDEF")

	// Generate a token object.
	token := webtoken.New(secret)
	token.SetClock(mc)

	// Random text in three sections.
	s, err := token.Verify("this.is.randomtext")
	assert.Equal(t, webtoken.ErrMalformed, err)
	assert.Equal(t, "", s)

	// Random text.
	s, err = token.Verify("this is randomtext")
	assert.Equal(t, webtoken.ErrMalformed, err)
	assert.Equal(t, "", s)

	// Invalid signature.
	txt := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJqc3B1cnJpZXIiLCJleHAiOjE1Mjk0NTA4MDMsImp0aSI6ImRhNWI0NzZjLTM2ZGYtMzkxNS0yMjU2LTJlYjg1MWYxZjMzMyIsImlhdCI6MTUyOTM2NDQwMywibmJmIjoxNTI5MzY0NDAzfQ.YCcAp7QQ9L0F_OIzEFWQu4v4fiERGvWCAJANO5S229`
	s, err = token.Verify(txt)
	assert.Equal(t, webtoken.ErrSignatureInvalid, err)
	assert.Equal(t, "", s)

	// Invalid expiration.
	txt = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJqc3B1cnJpZXIiLCJqdGkiOiJjOTU5ZTYzMC1lOWU5LTAwZjYtOWU1OS01ZDAzYTViMjczNDkiLCJpYXQiOjE1MjkzNzAzNDAsIm5iZiI6MTUyOTM3MDM0MH0.SGdZ50vAcBq_EuW8UkqGmjpBkQJJWGwLmOdMw1hcH2I`
	s, err = token.Verify(txt)
	assert.Equal(t, webtoken.ErrExpirationInvalid, err)
	assert.Equal(t, "", s)

	// Invalid not before date.
	txt = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJqc3B1cnJpZXIiLCJleHAiOjUxMjkzNjY3NjUsImp0aSI6IjQ1MjRiZjBlLTkwZDYtOTQwMS1hZTc4LTc2YmFlMjlhMmZmOSIsImlhdCI6MTUyOTM3MDM2NX0.VE7bPbSh3ZKLkWjJPTbVyqFQF4dIo8NBBPZIWJ92ch4`
	s, err = token.Verify(txt)
	assert.Equal(t, webtoken.ErrNotBeforeInvalid, err)
	assert.Equal(t, "", s)

	// Invalid issued at date.
	txt = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJqc3B1cnJpZXIiLCJleHAiOjUxMjkzNjY3ODksImp0aSI6IjI4YWUxOGVlLTM4NmItZDY1ZC1hZDNjLTZiYjFlNmVlNzNlNCIsIm5iZiI6MTUyOTM3MDM4OX0.muVGPA1nsXrZkG2RauY5aoFEdsr7gLObzYkBJyLj0l4`
	s, err = token.Verify(txt)
	assert.Equal(t, webtoken.ErrIssuedAtInvalid, err)
	assert.Equal(t, "", s)

	// Invalid audience.
	txt = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjUxMjkzNjY4MTcsImp0aSI6IjZmYTcwZjdkLWQ0ODMtZGJhZC0xNzkzLTQwMzc5ZDJjM2UxNiIsImlhdCI6MTUyOTM3MDQxNywibmJmIjoxNTI5MzcwNDE3fQ.VTsuIKug9LGfP7oz8LVKT8iBwCUsNyTfV8ftAuT5jn0`
	s, err = token.Verify(txt)
	assert.Equal(t, webtoken.ErrAudienceInvalid, err)
	assert.Equal(t, "", s)
}

func TestUnmarshal(t *testing.T) {
	type container struct {
		JWT webtoken.Configuration `json:"JWT"`
	}
	config := new(container)
	path := filepath.Join("testdata", "config.json")
	b, err := ioutil.ReadFile(path)
	assert.Nil(t, err)

	err = json.Unmarshal(b, &config)
	assert.Nil(t, err)

	assert.Equal(t, "0123456789ABCDEF0123456789ABCDEF", string(config.JWT.Secret))
}
