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

func TestIndexEmpty(t *testing.T) {
	db, unique := testutil.LoadDatabase()
	core, _ := component.NewCoreMock(db)

	w := testrequest.SendForm(t, core, "GET", "/v1/user", nil)

	r := new(model.UserIndexResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", r.Body.Status)
	assert.Equal(t, 0, len(r.Body.Data))

	testutil.TeardownDatabase(unique)
}

func TestIndexOne(t *testing.T) {
	db, unique := testutil.LoadDatabase()
	core, _ := component.NewCoreMock(db)

	u := store.NewUser(core.DB, core.Q)
	_, err := u.Create("John", "Smith", "jsmith@example.com", "password")
	assert.Nil(t, err)

	w := testrequest.SendForm(t, core, "GET", "/v1/user", nil)

	r := new(model.UserIndexResponse)
	err = json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 1, len(r.Body.Data))
	assert.Equal(t, "John", r.Body.Data[0].FirstName)
	assert.Equal(t, "Smith", r.Body.Data[0].LastName)
	assert.Equal(t, "jsmith@example.com", r.Body.Data[0].Email)

	testutil.TeardownDatabase(unique)
}
