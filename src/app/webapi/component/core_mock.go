package component

import (
	"log"

	"app/webapi/internal/bind"
	"app/webapi/internal/response"
	"app/webapi/internal/testutil"
	"app/webapi/pkg/database"
	"app/webapi/pkg/query"
)

// TestDatabase returns a test database.
func TestDatabase(dbSpecificDB bool) *database.DBW {
	dbc := new(database.Connection)
	dbc.Hostname = "127.0.0.1"
	dbc.Port = 3306
	dbc.Username = "root"
	dbc.Password = ""
	dbc.Database = "webapitest"
	dbc.Parameter = "parseTime=true&allowNativePasswords=true"

	connection, err := dbc.Connect(dbSpecificDB)
	if err != nil {
		log.Println("DB Error:", err)
	}

	dbw := database.New(connection)

	return dbw
}

// NewCoreMock returns all mocked dependencies.
func NewCoreMock() (Core, *CoreMock) {
	ml := new(testutil.MockLogger)
	//md := new(testutil.MockDatabase)
	md := TestDatabase(true)
	mq := query.New(md)
	mt := new(testutil.MockToken)
	resp := response.New()
	binder := bind.New()

	core := NewCore(ml, md, mq, binder, resp, mt)
	m := &CoreMock{
		Log:     ml,
		DB:      md,
		Q:       mq,
		Bind:    binder,
		Reponse: resp,
		Token:   mt,
	}
	return core, m
}

// CoreMock contains all the mocked dependencies.
type CoreMock struct {
	Log     *testutil.MockLogger
	DB      IDatabase
	Q       IQuery
	Bind    IBind
	Reponse IResponse
	Token   *testutil.MockToken
}
