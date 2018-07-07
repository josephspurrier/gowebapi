package bind_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"app/webapi/internal/bind"
	"app/webapi/pkg/router"

	"github.com/stretchr/testify/assert"
)

func TestSuccess(t *testing.T) {
	called := false

	mux := router.New()

	mux.Post("/user/:user_id", router.Handler(
		func(w http.ResponseWriter, r *http.Request) (status int, err error) {
			called = true

			// swagger:parameters UserCreate
			type request struct {
				// in: path
				UserID string `json:"user_id" validate:"required"`
				// in: formData
				// Required: true
				FirstName string `json:"first_name" validate:"required"`
				// in: formData
				// Required: true
				LastName string `json:"last_name" validate:"required"`
			}

			req := new(request)
			b := bind.New()

			assert.Nil(t, b.FormUnmarshal(&req, r))
			assert.Nil(t, b.Validate(req))

			assert.Equal(t, "10", req.UserID)
			assert.Equal(t, "john", req.FirstName)
			assert.Equal(t, "smith", req.LastName)
			return http.StatusOK, nil
		}))

	form := url.Values{}
	form.Add("first_name", "john")
	form.Add("last_name", "smith")

	r := httptest.NewRequest("POST", "/user/10", strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, true, called)
}

func TestMissingPointer(t *testing.T) {
	called := false

	mux := router.New()

	mux.Post("/user/:user_id", router.Handler(
		func(w http.ResponseWriter, r *http.Request) (status int, err error) {
			called = true

			// swagger:parameters UserCreate
			type request struct {
				// in: path
				UserID string `json:"user_id" validate:"required"`
				// in: formData
				// Required: true
				FirstName string `json:"first_name" validate:"required"`
				// in: formData
				// Required: true
				LastName string `json:"last_name" validate:"required"`
			}

			req := request{}
			b := bind.New()

			assert.NotNil(t, b.FormUnmarshal(req, r))

			assert.Equal(t, "", req.UserID)
			assert.Equal(t, "", req.FirstName)
			assert.Equal(t, "", req.LastName)
			return http.StatusOK, nil
		}))

	form := url.Values{}
	form.Add("first_name", "john")
	form.Add("last_name", "smith")

	r := httptest.NewRequest("POST", "/user/10", strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, true, called)
}
