package main_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	s := `first_name=consectetur&last_name=ipsum&email=Ut%20cillum%20in&password=proident`

	u, err := url.ParseQuery(s)
	assert.Nil(t, err)
	assert.Equal(t, "", u.Get("email"))
}
