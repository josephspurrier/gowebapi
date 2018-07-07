package auth_test

import (
	"encoding/base64"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"app/webapi/component"
	"app/webapi/component/auth"
	"app/webapi/pkg/router"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	core, m := component.NewCoreMock()

	m.Token.SetGenerate(func(userID string, duration time.Duration) (string, error) {
		b := []byte("0123456789ABCDEF0123456789ABCDEF")
		enc := base64.StdEncoding.EncodeToString(b)
		return enc, nil
	})

	mux := router.New()
	auth.New(core).Routes(mux)

	r := httptest.NewRequest("GET", "/v1/auth", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"status":"","data":{"token":"MDEyMzQ1Njc4OUFCQ0RFRjAxMjM0NTY3ODlBQkNERUY="}}`+"\n", w.Body.String())
}

func TestIndexError(t *testing.T) {
	core, m := component.NewCoreMock()

	m.Token.SetGenerate(func(userID string, duration time.Duration) (string, error) {
		return "", errors.New("generate error")
	})

	mux := router.New()
	auth.New(core).Routes(mux)

	r := httptest.NewRequest("GET", "/v1/auth", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, `generate error`+"\n", w.Body.String())
}
