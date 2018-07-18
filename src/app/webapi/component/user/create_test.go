package user_test

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"app/webapi/component"
	"app/webapi/internal/testrequest"
	"app/webapi/internal/testutil"
	"app/webapi/model"
	"app/webapi/store"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	db, unique := testutil.LoadDatabase()
	core, _ := component.NewCoreMock(db)

	form := url.Values{}
	form.Add("first_name", "John")
	form.Add("last_name", "Smith")
	form.Add("email", "jsmith@example.com")
	form.Add("password", "password")

	w := testrequest.SendForm(t, core, "POST", "/v1/user", form)

	r := new(model.CreatedResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "Created", r.Body.Status)
	assert.Equal(t, 36, len(r.Body.RecordID))

	testutil.TeardownDatabase(unique)
}

func TestCreateUserAlreadyExists(t *testing.T) {
	db, unique := testutil.LoadDatabase()
	core, _ := component.NewCoreMock(db)

	u := store.NewUser(core.DB, core.Q)
	_, err := u.Create("John", "Smith", "jsmith@example.com", "password")
	assert.Nil(t, err)

	form := url.Values{}
	form.Add("first_name", "John")
	form.Add("last_name", "Smith")
	form.Add("email", "jsmith@example.com")
	form.Add("password", "password")

	w := testrequest.SendForm(t, core, "POST", "/v1/user", form)

	r := new(model.BadRequestResponse)
	err = json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Bad Request", r.Body.Status)
	assert.Equal(t, "user already exists", r.Body.Message)

	testutil.TeardownDatabase(unique)
}

func TestCreateBadEmail(t *testing.T) {
	db, unique := testutil.LoadDatabase()
	core, _ := component.NewCoreMock(db)

	form := url.Values{}
	form.Add("first_name", "John")
	form.Add("last_name", "Smith")
	form.Add("email", "jsmith@bademail")
	form.Add("password", "password")

	w := testrequest.SendForm(t, core, "POST", "/v1/user", form)

	r := new(model.BadRequestResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Bad Request", r.Body.Status)
	assert.Contains(t, w.Body.String(), `failed`)

	testutil.TeardownDatabase(unique)
}

func TestCreateValidation(t *testing.T) {
	for _, v := range []string{
		"POST /v1/user",
	} {
		db, unique := testutil.LoadDatabase()
		core, _ := component.NewCoreMock(db)

		arr := strings.Split(v, " ")

		w := testrequest.SendForm(t, core, arr[0], arr[1], nil)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		testutil.TeardownDatabase(unique)
	}
}
