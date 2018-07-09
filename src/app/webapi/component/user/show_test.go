package user_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"app/webapi/component"
	"app/webapi/internal/testrequest"
	"app/webapi/internal/testutil"
	"app/webapi/model"
	"app/webapi/store"

	"github.com/stretchr/testify/assert"
)

func TestShowOne(t *testing.T) {
	testutil.LoadDatabase(t)
	core, _ := component.NewCoreMock()

	u := store.NewUser(core.DB, core.Q)
	ID, err := u.Create("John", "Smith", "jsmith@example.com", "password")
	assert.Nil(t, err)

	w := testrequest.SendForm(t, core, "GET", "/v1/user/"+ID, nil)

	r := new(model.UserShowResponse)
	err = json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", r.Body.Status)
	assert.Equal(t, 1, len(r.Body.Data))
	assert.Equal(t, "John", r.Body.Data[0].FirstName)
	assert.Equal(t, "Smith", r.Body.Data[0].LastName)
	assert.Equal(t, "jsmith@example.com", r.Body.Data[0].Email)
}

func TestShowNotFound(t *testing.T) {
	testutil.LoadDatabase(t)
	core, _ := component.NewCoreMock()

	w := testrequest.SendForm(t, core, "GET", "/v1/user/1", nil)

	r := new(model.BadRequestResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Bad Request", r.Body.Status)
	assert.Equal(t, "user not found", r.Body.Message)
}
