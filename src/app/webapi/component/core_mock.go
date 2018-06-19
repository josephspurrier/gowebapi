package component

import (
	"app/webapi/internal/bind"
	"app/webapi/internal/response"
	"app/webapi/internal/testutil"
)

// NewCoreMock returns all mocked dependencies.
func NewCoreMock() (Core, *CoreMock) {
	ml := new(testutil.MockLogger)
	md := new(testutil.MockDatabase)
	mt := new(testutil.MockToken)
	resp := response.New()
	binder := bind.New()

	core := NewCore(ml, md, binder, resp, mt)
	m := &CoreMock{
		Log:     ml,
		DB:      md,
		Bind:    binder,
		Reponse: resp,
		Token:   mt,
	}
	return core, m
}

// CoreMock contains all the mocked dependencies.
type CoreMock struct {
	Log     *testutil.MockLogger
	DB      *testutil.MockDatabase
	Bind    IBind
	Reponse IResponse
	Token   *testutil.MockToken
}
