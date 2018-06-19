package user_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"app/webapi/component"
	"app/webapi/component/user"
	"app/webapi/internal/testutil"
	"app/webapi/pkg/router"

	"github.com/stretchr/testify/assert"
)

func TestIndexEmpty(t *testing.T) {
	core, m := component.NewCoreMock()

	mux := router.New()
	user.New(core).Routes(mux)

	m.DB.SetPaginatedResults(testutil.PaginatedResultsEmpty)

	r := httptest.NewRequest("GET", "/v1/user", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `{"status":"OK","data":[]}`)
}

func TestIndexOne(t *testing.T) {
	core, m := component.NewCoreMock()

	mux := router.New()
	user.New(core).Routes(mux)

	m.DB.SetPaginatedResults(func() (interface{}, int, error) {
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
