package user_test

import (
	"app/webapi/model"
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"app/webapi/component"
	"app/webapi/internal/testrequest"
	"app/webapi/internal/testutil"
	"app/webapi/store"

	"github.com/stretchr/testify/assert"
)

func TestDestroyAll(t *testing.T) {
	testutil.LoadDatabase(t)
	core, _ := component.NewCoreMock()

	u := store.NewUser(core.DB, core.Q)
	_, err := u.Create("John", "Smith", "jsmith@example.com", "password")
	assert.Nil(t, err)

	w := testrequest.SendForm(t, core, "DELETE", "/v1/user", nil)

	r := new(model.OKResponse)
	err = json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", r.Body.Status)
	assert.Equal(t, "users deleted", r.Body.Message)
}

func TestDestroyAllNoUsers(t *testing.T) {
	testutil.LoadDatabase(t)
	core, _ := component.NewCoreMock()

	w := testrequest.SendForm(t, core, "DELETE", "/v1/user", nil)

	r := new(model.BadRequestResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Bad Request", r.Body.Status)
	assert.Equal(t, "no users to delete", r.Body.Message)
}

func TestDestroyValidation(t *testing.T) {
	for _, v := range []string{
		"DELETE /v1/user/1",
	} {
		testutil.LoadDatabase(t)
		core, _ := component.NewCoreMock()

		arr := strings.Split(v, " ")

		w := testrequest.SendForm(t, core, arr[0], arr[1], nil)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	}
}
