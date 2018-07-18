package user_test

import (
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"app/webapi/component"
	"app/webapi/internal/testrequest"
	"app/webapi/internal/testutil"
	"app/webapi/model"
	"app/webapi/store"

	"github.com/stretchr/testify/assert"
)

func TestUpdateUserAllFields(t *testing.T) {
	db, unique := testutil.LoadDatabase()
	core, _ := component.NewCoreMock(db)

	u := store.NewUser(core.DB, core.Q)
	ID, err := u.Create("John", "Smith", "jsmith@example.com", "password")
	assert.Nil(t, err)

	form := url.Values{}
	form.Add("first_name", "John1")
	form.Add("last_name", "Smith2")
	form.Add("email", "jsmith3@example.com")
	form.Add("password", "password4")

	w := testrequest.SendForm(t, core, "PUT", "/v1/user/"+ID, form)

	r := new(model.OKResponse)
	err = json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", r.Body.Status)
	assert.Equal(t, "user updated", r.Body.Message)

	found, err := u.FindOneByID(u, ID)
	assert.Nil(t, err)
	assert.True(t, found)
	assert.Equal(t, "John1", u.FirstName)
	assert.Equal(t, "Smith2", u.LastName)
	assert.Equal(t, "jsmith3@example.com", u.Email)
	assert.True(t, core.Password.MatchString(u.Password, "password4"))

	testutil.TeardownDatabase(unique)
}

func TestUpdateMissingFields(t *testing.T) {
	db, unique := testutil.LoadDatabase()
	core, _ := component.NewCoreMock(db)

	u := store.NewUser(core.DB, core.Q)
	ID, err := u.Create("John", "Smith", "jsmith@example.com", "password")
	assert.Nil(t, err)

	form := url.Values{}
	form.Add("first_name", "John1")

	w := testrequest.SendForm(t, core, "PUT", "/v1/user/"+ID, form)

	r := new(model.BadRequestResponse)
	err = json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Bad Request", r.Body.Status)
	assert.Contains(t, r.Body.Message, "failed")

	found, err := u.FindOneByID(u, ID)
	assert.Nil(t, err)
	assert.True(t, found)
	assert.Equal(t, "John", u.FirstName)
	assert.Equal(t, "Smith", u.LastName)
	assert.Equal(t, "jsmith@example.com", u.Email)
	assert.Equal(t, "password", u.Password)

	testutil.TeardownDatabase(unique)
}
