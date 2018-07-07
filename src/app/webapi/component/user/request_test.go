package user_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"app/webapi/component"
	"app/webapi/component/user"
	"app/webapi/internal/testutil"
	"app/webapi/pkg/router"

	"github.com/stretchr/testify/assert"
)

func TestRequestValidation(t *testing.T) {
	for _, v := range []string{
		"POST /v1/user",
		"DELETE /v1/user/1",
	} {
		arr := strings.Split(v, " ")

		testutil.LoadDatabase(t)
		core, _ := component.NewCoreMock()

		mux := router.New()
		user.New(core).Routes(mux)

		r := httptest.NewRequest(arr[0], arr[1], nil)
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	}
}
