package testutil

import (
	"database/sql"
	"net/http"
)

// MockLogger .
type MockLogger struct{}

// MockSQLResponse .
type MockSQLResponse struct{}

// LastInsertId .
func (ms *MockSQLResponse) LastInsertId() (int64, error) {
	return 0, nil
}

// RowsAffected .
func (ms *MockSQLResponse) RowsAffected() (int64, error) {
	return 0, nil
}

// MockDatabase .
type MockDatabase struct{}

// Select .
func (d *MockDatabase) Select(dest interface{}, query string, args ...interface{}) error {
	return nil
}

// Get .
func (d *MockDatabase) Get(dest interface{}, query string, args ...interface{}) error {
	return nil
}

// Exec .
func (d *MockDatabase) Exec(query string, args ...interface{}) (sql.Result, error) {
	ms := new(MockSQLResponse)
	return ms, nil
}

// ExistsString .
func (d *MockDatabase) ExistsString(err error, s string) (bool, string, error) {
	return false, "", nil
}

// Error .
func (d *MockDatabase) Error(err error) error {
	return nil
}

// AffectedRows .
func (d *MockDatabase) AffectedRows(result sql.Result) int {
	return 0
}

// *****************************************************************************

type recordExistsIntFunc func() (exists bool, ID int64, err error)

var (
	recordExistsInt recordExistsIntFunc

	// RecordExistsIntNot returns false, 0, nil.
	RecordExistsIntNot = func() (exists bool, ID int64, err error) {
		return false, 0, nil
	}
)

// SetRecordExistsInt will set the function.
func (d *MockDatabase) SetRecordExistsInt(fn recordExistsIntFunc) {
	recordExistsInt = fn
}

// RecordExistsInt returns the ID if a record exists.
func (d *MockDatabase) RecordExistsInt(fn func() (exists bool, ID int64, err error)) (
	exists bool, ID int64, err error) {
	// Use the default.
	fnInternal := recordExistsInt
	if fnInternal == nil {
		fnInternal = fn
	}

	return fnInternal()
}

// *****************************************************************************

type recordExistsStringFunc func() (exists bool, ID string, err error)

var (
	recordExistsString recordExistsStringFunc

	// RecordExistsStringNot returns false, "", nil.
	RecordExistsStringNot = func() (exists bool, ID string, err error) {
		return false, "", nil
	}
)

// SetRecordExistsString will set the function.
func (d *MockDatabase) SetRecordExistsString(fn recordExistsStringFunc) {
	recordExistsString = fn
}

// RecordExistsString returns the ID if a record exists.
func (d *MockDatabase) RecordExistsString(fn func() (exists bool, ID string, err error)) (
	exists bool, ID string, err error) {
	// Use the default.
	fnInternal := recordExistsString
	if fnInternal == nil {
		fnInternal = fn
	}

	return fnInternal()
}

// *****************************************************************************

type addRecordIntFunc func() (ID int64, err error)

var (
	addRecordInt addRecordIntFunc

	// AddRecordIntDefault returns 0, nil.
	AddRecordIntDefault = func() (ID int64, err error) {
		return 0, nil
	}
)

// SetAddRecordInt will set the function.
func (d *MockDatabase) SetAddRecordInt(fn addRecordIntFunc) {
	addRecordInt = fn
}

// AddRecordInt returns the ID if the record is created.
func (d *MockDatabase) AddRecordInt(fn func() (ID int64, err error)) (ID int64, err error) {
	// Use the default.
	fnInternal := addRecordInt
	if fnInternal == nil {
		fnInternal = fn
	}
	return fnInternal()
}

// *****************************************************************************

type addRecordStringFunc func() (ID string, err error)

var (
	addRecordString addRecordStringFunc

	// AddRecordStringDefault returns "", nil.
	AddRecordStringDefault = func() (ID string, err error) {
		return "", nil
	}
)

// SetAddRecordString will set the function.
func (d *MockDatabase) SetAddRecordString(fn addRecordStringFunc) {
	addRecordString = fn
}

// AddRecordString returns the ID if the record is created.
func (d *MockDatabase) AddRecordString(fn func() (ID string, err error)) (ID string, err error) {
	// Use the default.
	fnInternal := addRecordString
	if fnInternal == nil {
		fnInternal = fn
	}
	return fnInternal()
}

// *****************************************************************************

// MockBind .
type MockBind struct{}

// FormUnmarshal .
func (mb *MockBind) FormUnmarshal(i interface{}, r *http.Request) (err error) {
	return nil
}

// Validate .
func (mb *MockBind) Validate(s interface{}) error {
	return nil
}

// MockResponse .
type MockResponse struct{}

// Created .
func (mr *MockResponse) Created(w http.ResponseWriter, recordID string) (int, error) {
	return 0, nil
}

// Results .
func (mr *MockResponse) Results(w http.ResponseWriter, body interface{}, data interface{}) (int, error) {
	return 0, nil
}

// OK .
func (mr *MockResponse) OK(w http.ResponseWriter, message string) (int, error) {
	return 0, nil
}
