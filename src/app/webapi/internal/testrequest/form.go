package testrequest

import (
	"io"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"app/webapi"
	"app/webapi/component"
)

// SendForm is a helper to quickly make a form request.
func SendForm(t *testing.T, core component.Core, method string, target string,
	v url.Values) *httptest.ResponseRecorder {
	mux := webapi.Routes(core)

	var body io.Reader
	if v != nil {
		body = strings.NewReader(v.Encode())
	}

	r := httptest.NewRequest(method, target, body)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	return w
}
