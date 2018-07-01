package user_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"app/webapi/component"
	"app/webapi/component/user"
	"app/webapi/internal/testdb"
	"app/webapi/pkg/router"
	"app/webapi/store"

	"github.com/stretchr/testify/assert"
)

func TestIndexEmpty(t *testing.T) {
	testdb.SetupTest(t)
	core, _ := component.NewCoreMock()

	mux := router.New()
	user.New(core).Routes(mux)

	r := httptest.NewRequest("GET", "/v1/user", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `{"status":"OK","data":[]}`)
}

func TestIndexOne(t *testing.T) {
	testdb.SetupTest(t)
	core, _ := component.NewCoreMock()

	mux := router.New()
	user.New(core).Routes(mux)

	u := store.NewUser(core.DB, core.Q)
	_, err := u.Create("John", "Smith", "jsmith@example.com", "password")
	assert.Nil(t, err)

	r := httptest.NewRequest("GET", "/v1/user", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"first_name":"John","last_name":"Smith"`)
}
