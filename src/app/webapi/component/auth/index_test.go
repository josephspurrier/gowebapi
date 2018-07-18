package auth_test

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"testing"
	"time"

	"app/webapi/component"
	"app/webapi/internal/testrequest"
	"app/webapi/internal/testutil"
	"app/webapi/model"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	db, unique := testutil.LoadDatabase()
	core, m := component.NewCoreMock(db)

	m.Token.GenerateFunc = func(userID string, duration time.Duration) (string, error) {
		b := []byte("0123456789ABCDEF0123456789ABCDEF")
		enc := base64.StdEncoding.EncodeToString(b)
		return enc, nil
	}

	w := testrequest.SendForm(t, core, "GET", "/v1/auth", nil)

	r := new(model.AuthIndexResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", r.Body.Status)
	assert.Equal(t, "MDEyMzQ1Njc4OUFCQ0RFRjAxMjM0NTY3ODlBQkNERUY=", r.Body.Data.Token)

	testutil.TeardownDatabase(unique)
}

func TestIndexError(t *testing.T) {
	db, unique := testutil.LoadDatabase()
	core, m := component.NewCoreMock(db)

	m.Token.GenerateFunc = func(userID string, duration time.Duration) (string, error) {
		return "", errors.New("generate error")
	}

	w := testrequest.SendForm(t, core, "GET", "/v1/auth", nil)

	r := new(model.InternalServerErrorResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "Internal Server Error", r.Body.Status)
	assert.Equal(t, "generate error", r.Body.Message)

	testutil.TeardownDatabase(unique)
}
