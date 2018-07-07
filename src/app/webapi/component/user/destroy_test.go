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

func TestDestroy(t *testing.T) {
	testutil.LoadDatabase(t)
	core, _ := component.NewCoreMock()

	mux := router.New()
	user.New(core).Routes(mux)

	u := store.NewUser(core.DB, core.Q)
	ID, err := u.Create("John", "Smith", "jsmith@example.com", "password")
	assert.Nil(t, err)

	form := url.Values{}
	form.Add("first_name", "John")
	form.Add("last_name", "Smith")
	form.Add("email", "jsmith@example.com")
	form.Add("password", "password")

	r := httptest.NewRequest("DELETE", "/v1/user/"+ID, strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `{"status":"OK","message":"user deleted"}`)

	found, err := u.FindOneByID(u, ID)
	assert.Nil(t, err)
	assert.False(t, found)
}
