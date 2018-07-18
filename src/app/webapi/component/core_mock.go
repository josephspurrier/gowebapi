package component

import (
	"app/webapi/internal/bind"
	"app/webapi/internal/response"
	"app/webapi/internal/testutil"
	"app/webapi/pkg/database"
	"app/webapi/pkg/passhash"
	"app/webapi/pkg/query"
)

// NewCoreMock returns all mocked dependencies.
func NewCoreMock(db *database.DBW) (Core, *CoreMock) {
	ml := new(testutil.MockLogger)
	mq := query.New(db)
	mt := new(testutil.MockToken)
	resp := response.New()
	binder := bind.New()
	p := passhash.New()

	core := NewCore(ml, db, mq, binder, resp, mt, p)
	m := &CoreMock{
		Log:      ml,
		DB:       db,
		Q:        mq,
		Bind:     binder,
		Response: resp,
		Token:    mt,
		Password: p,
	}
	return core, m
}

// CoreMock contains all the mocked dependencies.
type CoreMock struct {
	Log      *testutil.MockLogger
	DB       IDatabase
	Q        IQuery
	Bind     IBind
	Response IResponse
	Token    *testutil.MockToken
	Password IPassword
}
