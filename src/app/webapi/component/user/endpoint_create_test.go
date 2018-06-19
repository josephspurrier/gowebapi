package user_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"app/webapi/component"
	"app/webapi/component/user"
	"app/webapi/pkg/router"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	core, m := component.NewCoreMock()

	mux := router.New()
	user.New(core).Routes(mux)

	m.DB.SetRecordExistsString(func() (exists bool, ID string, err error) {
		return false, "", nil
	})

	m.DB.SetAddRecordString(func() (ID string, err error) {
		return "123", nil
	})

	form := url.Values{}
	form.Add("first_name", "John")
	form.Add("last_name", "Smith")
	form.Add("email", "jsmith@example.com")
	form.Add("password", "password")

	r := httptest.NewRequest("POST", "/v1/user", strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `{"status":"Created","record_id":"123"}`)
}

func TestCreateUserAlreadyExists(t *testing.T) {
	core, m := component.NewCoreMock()

	mux := router.New()
	user.New(core).Routes(mux)

	m.DB.SetRecordExistsString(func() (exists bool, ID string, err error) {
		return true, "", nil
	})

	form := url.Values{}
	form.Add("first_name", "John")
	form.Add("last_name", "Smith")
	form.Add("email", "jsmith@example.com")
	form.Add("password", "password")

	r := httptest.NewRequest("POST", "/v1/user", strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `{"status":"Bad Request","message":"user already exists"}`)
}

func TestCreateUserInternalError(t *testing.T) {
	core, m := component.NewCoreMock()

	mux := router.New()
	user.New(core).Routes(mux)

	m.DB.SetRecordExistsString(func() (exists bool, ID string, err error) {
		return true, "", errors.New("bad error")
	})

	form := url.Values{}
	form.Add("first_name", "John")
	form.Add("last_name", "Smith")
	form.Add("email", "jsmith@example.com")
	form.Add("password", "password")

	r := httptest.NewRequest("POST", "/v1/user", strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), `{"status":"Internal Server Error","message":"bad error"}`)
}

func TestCreateUserInternalError2(t *testing.T) {
	core, m := component.NewCoreMock()

	mux := router.New()
	user.New(core).Routes(mux)

	m.DB.SetRecordExistsString(func() (exists bool, ID string, err error) {
		return false, "", nil
	})

	m.DB.SetAddRecordString(func() (ID string, err error) {
		return "", errors.New("bad error 2")
	})

	form := url.Values{}
	form.Add("first_name", "John")
	form.Add("last_name", "Smith")
	form.Add("email", "jsmith@example.com")
	form.Add("password", "password")

	r := httptest.NewRequest("POST", "/v1/user", strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), `{"status":"Internal Server Error","message":"bad error 2"}`)
}

func TestCreateMissingItem(t *testing.T) {
	core, _ := component.NewCoreMock()

	mux := router.New()
	user.New(core).Routes(mux)

	form := url.Values{}
	form.Add("first_name", "John")
	form.Add("last_name", "Smith")
	form.Add("email", "jsmith@example.com")

	r := httptest.NewRequest("POST", "/v1/user", strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `Error:Field validation for 'Password' failed on the 'required' tag`)
}

func TestCreateBadEmail(t *testing.T) {
	core, _ := component.NewCoreMock()

	mux := router.New()
	user.New(core).Routes(mux)

	form := url.Values{}
	form.Add("first_name", "John")
	form.Add("last_name", "Smith")
	form.Add("email", "jsmith@bademail")
	form.Add("password", "password")

	r := httptest.NewRequest("POST", "/v1/user", strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `Error:Field validation for 'Email' failed on the 'email' tag`)
}
