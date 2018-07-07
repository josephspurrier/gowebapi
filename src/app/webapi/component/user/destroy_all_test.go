package user_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"app/webapi/component"
	"app/webapi/component/user"
	"app/webapi/internal/testutil"
	"app/webapi/pkg/router"
	"app/webapi/store"

	"github.com/stretchr/testify/assert"
)

func TestDestroyAll(t *testing.T) {
	testutil.LoadDatabase(t)
	core, _ := component.NewCoreMock()

	mux := router.New()
	user.New(core).Routes(mux)

	u := store.NewUser(core.DB, core.Q)
	_, err := u.Create("John", "Smith", "jsmith@example.com", "password")
	assert.Nil(t, err)

	r := httptest.NewRequest("DELETE", "/v1/user", nil)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `{"status":"OK","message":"users deleted"}`)
}

func TestDestroyAllNoUsers(t *testing.T) {
	testutil.LoadDatabase(t)
	core, _ := component.NewCoreMock()

	mux := router.New()
	user.New(core).Routes(mux)

	r := httptest.NewRequest("DELETE", "/v1/user", nil)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `no users to delete`)
}
