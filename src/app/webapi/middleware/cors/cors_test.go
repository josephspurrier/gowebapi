package cors

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerGet(t *testing.T) {
	called := false

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		called = true
	})

	r := httptest.NewRequest("GET", "/", nil)
	r.Header.Set(origin, "localhost")
	w := httptest.NewRecorder()

	Handler(mux).ServeHTTP(w, r)

	assert.Equal(t, true, called)
	assert.Equal(t, http.StatusOK, w.Code)

	wo := w.Header().Get(allowOrigin)
	ro := r.Header.Get(origin)
	assert.Equal(t, ro, wo, "Wrong Allow-Origin")

	wm := w.Header().Get(allowMethods)
	assert.Equal(t, methods, wm, "Wrong Allow-Methods")

	wh := w.Header().Get(allowHeaders)
	assert.Equal(t, headers, wh, "Wrong Allow-Headers")
}

func TestHandlerOptions(t *testing.T) {
	called := false

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		called = false
	})

	r := httptest.NewRequest("OPTIONS", "/", nil)
	r.Header.Set(origin, "localhost")
	w := httptest.NewRecorder()

	Handler(mux).ServeHTTP(w, r)

	assert.Equal(t, false, called)
	assert.Equal(t, http.StatusOK, w.Code)

	wo := w.Header().Get(allowOrigin)
	ro := r.Header.Get(origin)
	assert.Equal(t, ro, wo, "Wrong Allow-Origin")

	wm := w.Header().Get(allowMethods)
	assert.Equal(t, methods, wm, "Wrong Allow-Methods")

	wh := w.Header().Get(allowHeaders)
	assert.Equal(t, headers, wh, "Wrong Allow-Headers")
}

func TestHandlerNoOrigin(t *testing.T) {
	called := false

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		called = false
	})

	r := httptest.NewRequest("OPTIONS", "/", nil)
	w := httptest.NewRecorder()

	Handler(mux).ServeHTTP(w, r)

	assert.Equal(t, false, called)
	assert.Equal(t, http.StatusOK, w.Code)

	wo := w.Header().Get(allowOrigin)
	assert.Equal(t, "*", wo, "Wrong Allow-Origin")

	wm := w.Header().Get(allowMethods)
	assert.Equal(t, methods, wm, "Wrong Allow-Methods")

	wh := w.Header().Get(allowHeaders)
	assert.Equal(t, headers, wh, "Wrong Allow-Headers")
}
