package root_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"app/webapi/component"
	"app/webapi/internal/testrequest"
	"app/webapi/internal/testutil"
	"app/webapi/model"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	db, unique := testutil.LoadDatabase()
	core, _ := component.NewCoreMock(db)

	w := testrequest.SendForm(t, core, "GET", "/v1", nil)

	r := new(model.OKResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", r.Body.Status)
	assert.Equal(t, "hello", r.Body.Message)

	testutil.TeardownDatabase(unique)
}
