package user_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"app/webapi/component"
	"app/webapi/component/user"
	"app/webapi/internal/testutil"
	"app/webapi/pkg/router"
	"app/webapi/store"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	testutil.LoadDatabase(t)
	core, _ := component.NewCoreMock()

	mux := router.New()
	user.New(core).Routes(mux)

	form := url.Values{}
	form.Add("first_name", "John")
	form.Add("last_name", "Smith")
	form.Add("email", "jsmith@example.com")
	form.Add("password", "password")

	r := httptest.NewRequest("POST", "/v1/user", strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), `{"status":"Created","record_id"`)
}

func TestCreateUserAlreadyExists(t *testing.T) {
	testutil.LoadDatabase(t)
	core, _ := component.NewCoreMock()

	mux := router.New()
	user.New(core).Routes(mux)

	u := store.NewUser(core.DB, core.Q)
	_, err := u.Create("John", "Smith", "jsmith@example.com", "password")
	assert.Nil(t, err)

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
	assert.Contains(t, w.Body.String(), `user already exists`)
}

func TestCreateBadEmail(t *testing.T) {
	testutil.LoadDatabase(t)
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
