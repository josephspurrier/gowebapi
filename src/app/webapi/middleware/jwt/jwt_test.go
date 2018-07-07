package jwt_test

import (
	"app/webapi/middleware/jwt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhitelistAllowed(t *testing.T) {
	for _, v := range []string{
		"GET /v1",
		"GET /v1/auth",
	} {
		arr := strings.Split(v, " ")

		mux := http.NewServeMux()

		whitelist := []string{
			"GET /v1",
			"GET /v1/auth",
		}

		token := jwt.New([]byte("secret"), whitelist)
		h := token.Handler(mux)

		r := httptest.NewRequest(arr[0], arr[1], nil)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, r)

		assert.Equal(t, http.StatusNotFound, w.Code)
	}
}
func TestWhitelistNotAllowed(t *testing.T) {
	for _, v := range []string{
		"POST /v1",
		"POST /v1/auth",
		"POST /v1/user",
		"DELETE /v1/user/1",
	} {
		arr := strings.Split(v, " ")

		mux := http.NewServeMux()

		whitelist := []string{
			"GET /v1",
			"GET /v1/auth",
		}

		token := jwt.New([]byte("secret"), whitelist)
		h := token.Handler(mux)

		r := httptest.NewRequest(arr[0], arr[1], nil)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, r)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), `authorization token is missing`)
	}
}

func TestWhitelistBadBearer(t *testing.T) {
	mux := http.NewServeMux()

	whitelist := []string{
		"GET /v1",
		"GET /v1/auth",
	}

	token := jwt.New([]byte("secret"), whitelist)
	h := token.Handler(mux)

	r := httptest.NewRequest("POST", "/v1/user", nil)
	w := httptest.NewRecorder()
	r.Header.Set("Authorization", "Bearer bad")
	h.ServeHTTP(w, r)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), `authorization token is invalid`)
}

func TestIsWhitelisted(t *testing.T) {
	assert.Equal(t, true, jwt.IsWhitelisted("GET", "/v1", []string{
		"GET /v1",
	}))

	assert.Equal(t, true, jwt.IsWhitelisted("GET", "/v1", []string{
		"*",
	}))

	assert.Equal(t, true, jwt.IsWhitelisted("GET", "/v1", []string{
		"POST /v1",
		"*",
	}))

	// Bad spacing.
	assert.Equal(t, false, jwt.IsWhitelisted("GET", "/v1", []string{
		"POST /v1",
		"* ",
	}))

	// Not in the list.
	assert.Equal(t, false, jwt.IsWhitelisted("GET", "/v2", []string{
		"POST /v1",
		"GET /v1",
	}))
}
