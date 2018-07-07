package router_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"app/webapi/pkg/router"

	"github.com/stretchr/testify/assert"
)

func TestParams(t *testing.T) {
	mux := router.New()
	mux.Get("/user/:name", router.Handler(
		func(w http.ResponseWriter, r *http.Request) (status int, err error) {
			assert.Equal(t, "john", router.Params(r, "name"))
			return http.StatusOK, nil
		}))

	r := httptest.NewRequest("GET", "/user/john", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
}

func TestInstance(t *testing.T) {
	mux := router.New()

	mux.Get("/user/:name", router.Handler(
		func(w http.ResponseWriter, r *http.Request) (status int, err error) {
			assert.Equal(t, "john", router.Params(r, "name"))
			return http.StatusOK, nil
		}))

	r := httptest.NewRequest("GET", "/user/john", nil)
	w := httptest.NewRecorder()

	mux.Instance().ServeHTTP(w, r)
}

func TestPostForm(t *testing.T) {
	mux := router.New()

	form := url.Values{}
	form.Add("username", "jsmith")

	mux.Post("/user", router.Handler(
		func(w http.ResponseWriter, r *http.Request) (status int, err error) {
			r.ParseForm()
			assert.Equal(t, "jsmith", r.FormValue("username"))
			return http.StatusOK, nil
		}))

	r := httptest.NewRequest("POST", "/user", strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
}

func TestPostJSON(t *testing.T) {
	mux := router.New()

	j, err := json.Marshal(map[string]interface{}{
		"username": "jsmith",
	})
	assert.Nil(t, err)

	mux.Post("/user", router.Handler(
		func(w http.ResponseWriter, r *http.Request) (status int, err error) {
			b, err := ioutil.ReadAll(r.Body)
			assert.Nil(t, err)
			r.Body.Close()
			assert.Equal(t, `{"username":"jsmith"}`, string(b))
			return http.StatusOK, nil
		}))

	r := httptest.NewRequest("POST", "/user", bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
}

func TestGet(t *testing.T) {
	mux := router.New()

	called := false

	mux.Get("/user", router.Handler(
		func(w http.ResponseWriter, r *http.Request) (status int, err error) {
			called = true
			return http.StatusOK, nil
		}))

	r := httptest.NewRequest("GET", "/user", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, true, called)
}

func TestDelete(t *testing.T) {
	mux := router.New()

	called := false

	mux.Delete("/user", router.Handler(
		func(w http.ResponseWriter, r *http.Request) (status int, err error) {
			called = true
			return http.StatusOK, nil
		}))

	r := httptest.NewRequest("DELETE", "/user", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, true, called)
}

func TestHead(t *testing.T) {
	mux := router.New()

	called := false

	mux.Head("/user", router.Handler(
		func(w http.ResponseWriter, r *http.Request) (status int, err error) {
			called = true
			return http.StatusOK, nil
		}))

	r := httptest.NewRequest("HEAD", "/user", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, true, called)
}

func TestOptions(t *testing.T) {
	mux := router.New()

	called := false

	mux.Options("/user", router.Handler(
		func(w http.ResponseWriter, r *http.Request) (status int, err error) {
			called = true
			return http.StatusOK, nil
		}))

	r := httptest.NewRequest("OPTIONS", "/user", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, true, called)
}

func TestPatch(t *testing.T) {
	mux := router.New()

	called := false

	mux.Patch("/user", router.Handler(
		func(w http.ResponseWriter, r *http.Request) (status int, err error) {
			called = true
			return http.StatusOK, nil
		}))

	r := httptest.NewRequest("PATCH", "/user", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, true, called)
}

func TestPut(t *testing.T) {
	mux := router.New()

	called := false

	mux.Put("/user", router.Handler(
		func(w http.ResponseWriter, r *http.Request) (status int, err error) {
			called = true
			return http.StatusOK, nil
		}))

	r := httptest.NewRequest("PUT", "/user", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, true, called)
}
