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

func TestUpdateUserAllFields(t *testing.T) {
	testutil.LoadDatabase(t)
	core, _ := component.NewCoreMock()

	mux := router.New()
	user.New(core).Routes(mux)

	u := store.NewUser(core.DB, core.Q)
	ID, err := u.Create("John", "Smith", "jsmith@example.com", "password")
	assert.Nil(t, err)

	form := url.Values{}
	form.Add("first_name", "John1")
	form.Add("last_name", "Smith2")
	form.Add("email", "jsmith3@example.com")
	form.Add("password", "password4")

	r := httptest.NewRequest("PUT", "/v1/user/"+ID, strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `{"status":"OK","message":"user updated"}`)

	found, err := u.FindOneByID(u, ID)
	assert.Nil(t, err)
	assert.True(t, found)
	assert.Equal(t, "John1", u.FirstName)
	assert.Equal(t, "Smith2", u.LastName)
	assert.Equal(t, "jsmith3@example.com", u.Email)
	assert.Equal(t, "password4", u.Password)
}

func TestUpdateMissingFields(t *testing.T) {
	testutil.LoadDatabase(t)
	core, _ := component.NewCoreMock()

	mux := router.New()
	user.New(core).Routes(mux)

	u := store.NewUser(core.DB, core.Q)
	ID, err := u.Create("John", "Smith", "jsmith@example.com", "password")
	assert.Nil(t, err)

	form := url.Values{}
	form.Add("first_name", "John1")

	r := httptest.NewRequest("PUT", "/v1/user/"+ID, strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `Error:Field validation`)

	found, err := u.FindOneByID(u, ID)
	assert.Nil(t, err)
	assert.True(t, found)
	assert.Equal(t, "John", u.FirstName)
	assert.Equal(t, "Smith", u.LastName)
	assert.Equal(t, "jsmith@example.com", u.Email)
	assert.Equal(t, "password", u.Password)
}
