package root_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"app/webapi/component"
	"app/webapi/component/root"
	"app/webapi/internal/response"
	"app/webapi/internal/testutil"
	"app/webapi/pkg/router"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	ml := new(testutil.MockLogger)
	md := new(testutil.MockDatabase)
	mb := new(testutil.MockBind)
	resp := response.New()

	mux := router.New()
	core := component.New(ml, md, mb, resp)
	root.New(core).Routes(mux)

	r := httptest.NewRequest("GET", "/v1", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"status":"OK","message":"hello"}`+"\n", w.Body.String())
}
