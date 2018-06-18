package user_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"app/webapi/component"
	"app/webapi/component/user"
	"app/webapi/internal/bind"
	"app/webapi/internal/response"
	"app/webapi/internal/testutil"
	"app/webapi/pkg/router"

	"github.com/stretchr/testify/assert"
)

func TestIndexEmpty(t *testing.T) {
	ml := new(testutil.MockLogger)
	md := new(testutil.MockDatabase)

	binder := bind.New()
	resp := response.New()

	mux := router.New()
	core := component.New(ml, md, binder, resp)
	user.New(core).Routes(mux)

	md.SetPaginatedResults(testutil.PaginatedResultsEmpty)

	r := httptest.NewRequest("GET", "/v1/user", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `{"status":"OK","data":[]}`)
}

func TestIndexOne(t *testing.T) {
	ml := new(testutil.MockLogger)
	md := new(testutil.MockDatabase)

	binder := bind.New()
	resp := response.New()

	mux := router.New()
	core := component.New(ml, md, binder, resp)
	user.New(core).Routes(mux)

	md.SetPaginatedResults(func() (interface{}, int, error) {
		results := make([]user.TUser, 0)
		results = append(results, user.TUser{
			FirstName: "John",
			LastName:  "Smith",
		})
		return results, 1, nil
	})

	r := httptest.NewRequest("GET", "/v1/user", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"first_name":"John","last_name":"Smith"`)
}
