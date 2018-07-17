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

func TestDestroy(t *testing.T) {
	db, unique := testutil.LoadDatabase()
	core, _ := component.NewCoreMock(db)

	u := store.NewUser(core.DB, core.Q)
	ID, err := u.Create("John", "Smith", "jsmith@example.com", "password")
	assert.Nil(t, err)

	form := url.Values{}
	form.Add("first_name", "John")
	form.Add("last_name", "Smith")
	form.Add("email", "jsmith@example.com")
	form.Add("password", "password")

	w := testrequest.SendForm(t, core, "DELETE", "/v1/user/"+ID, form)

	r := new(model.OKResponse)
	err = json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", r.Body.Status)
	assert.Equal(t, "user deleted", r.Body.Message)

	found, err := u.FindOneByID(u, ID)
	assert.Nil(t, err)
	assert.False(t, found)

	testutil.TeardownDatabase(unique)
}
