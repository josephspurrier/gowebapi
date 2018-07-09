package webapi_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"app/webapi"
	"app/webapi/pkg/database"
	"app/webapi/pkg/jsonconfig"

	"github.com/stretchr/testify/assert"
)

type MockLogger struct {
	FatalfOutput string
	PrintfOutput string
}

func (l *MockLogger) Fatalf(format string, v ...interface{}) {
	l.FatalfOutput = fmt.Sprintf(format, v...)
}

func (l *MockLogger) Printf(format string, v ...interface{}) {
	l.PrintfOutput = fmt.Sprintf(format, v...)
}

func TestDatabase(t *testing.T) {
	dbc := new(database.Connection)
	dbc.Hostname = "127.0.0.1"
	dbc.Port = 3306
	dbc.Username = "root"
	dbc.Password = ""
	dbc.Database = ""
	dbc.Parameter = "parseTime=true&allowNativePasswords=true"

	ml := new(MockLogger)

	dbw := webapi.Database(*dbc, ml)
	assert.NotNil(t, dbw)
	assert.Empty(t, ml.PrintfOutput)
}

func TestDatabaseFail(t *testing.T) {
	dbc := new(database.Connection)
	dbc.Hostname = "127.0.0.1"
	dbc.Port = 3306
	dbc.Username = "root"
	dbc.Password = "bad-password"
	dbc.Database = ""
	dbc.Parameter = "parseTime=true&allowNativePasswords=true"

	ml := new(MockLogger)

	dbw := webapi.Database(*dbc, ml)
	assert.NotNil(t, dbw)
	assert.Contains(t, ml.PrintfOutput, "Access denied")
}

func TestDatabaseFailEnv(t *testing.T) {
	dbc := new(database.Connection)
	dbc.Hostname = "127.0.0.1"
	dbc.Port = 3306
	dbc.Username = "root"
	dbc.Password = ""
	dbc.Database = ""
	dbc.Parameter = "parseTime=true&allowNativePasswords=true"

	ml := new(MockLogger)

	os.Setenv("DB_PASSWORD", "bad-password")

	dbw := webapi.Database(*dbc, ml)
	assert.NotNil(t, dbw)
	assert.Contains(t, ml.PrintfOutput, "Access denied")

	os.Setenv("DB_PASSWORD", "")
}

func TestRoutes200(t *testing.T) {
	c := new(webapi.AppConfig)

	dbc := new(database.Connection)
	dbc.Hostname = "127.0.0.1"
	dbc.Port = 3306
	dbc.Username = "root"
	dbc.Password = ""
	dbc.Database = ""
	dbc.Parameter = "parseTime=true&allowNativePasswords=true"

	c.Database = *dbc
	ml := new(MockLogger)
	core := webapi.Services(c, ml)
	mux := webapi.Routes(core)

	r := httptest.NewRequest("GET", "/v1", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"status":"OK","message":"hello"}`+"\n", w.Body.String())
}

func TestRoutes404(t *testing.T) {
	c := new(webapi.AppConfig)

	dbc := new(database.Connection)
	dbc.Hostname = "127.0.0.1"
	dbc.Port = 3306
	dbc.Username = "root"
	dbc.Password = ""
	dbc.Database = ""
	dbc.Parameter = "parseTime=true&allowNativePasswords=true"

	c.Database = *dbc
	ml := new(MockLogger)
	core := webapi.Services(c, ml)
	mux := webapi.Routes(core)

	r := httptest.NewRequest("GET", "/v1/nope", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), `{"status":"Not Found"}`)
}

func TestRoutes500(t *testing.T) {
	c := new(webapi.AppConfig)

	dbc := new(database.Connection)
	dbc.Hostname = "127.0.0.1"
	dbc.Port = 3306
	dbc.Username = "root"
	dbc.Password = ""
	dbc.Database = ""
	dbc.Parameter = "parseTime=true&allowNativePasswords=true"

	c.Database = *dbc
	ml := new(MockLogger)
	core := webapi.Services(c, ml)
	mux := webapi.Routes(core)

	r := httptest.NewRequest("GET", "/v1/user", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), `Internal Server Error`)
	assert.Contains(t, ml.PrintfOutput, "No database selected")
}

func TestParseAppConfig(t *testing.T) {
	config := new(webapi.AppConfig)
	err := jsonconfig.Load("../../../config.json", config)
	assert.Nil(t, err)

	assert.Equal(t, "127.0.0.1", config.Database.Hostname)
	assert.Equal(t, 8080, config.Server.HTTPPort)
}
