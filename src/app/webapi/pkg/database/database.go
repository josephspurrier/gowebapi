package database

import (
	"database/sql"
	"errors"
	"reflect"
	"time"

	"github.com/jmoiron/sqlx"
)

// New returns a new database wrapper.
func New(db *sqlx.DB) *DBW {
	return &DBW{
		db: db,
	}
}

// DBW is a database wrapper that provides helpful utilities.
type DBW struct {
	db *sqlx.DB
}

// Select using this DB.
// Any placeholder parameters are replaced with supplied args.
func (d *DBW) Select(dest interface{}, query string, args ...interface{}) error {
	return d.db.Select(dest, query, args...)
}

// Get using this DB.
// Any placeholder parameters are replaced with supplied args.
// An error is returned if the result set is empty.
func (d *DBW) Get(dest interface{}, query string, args ...interface{}) error {
	return d.db.Get(dest, query, args...)
}

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
func (d *DBW) Exec(query string, args ...interface{}) (sql.Result, error) {
	return d.db.Exec(query, args...)
}

// LastInsertID returns the last inserted ID.
func (d *DBW) LastInsertID(r sql.Result, err error) (int64, error) {
	if err != nil {
		return 0, err
	}

	return r.LastInsertId()
}

// MySQLTimestamp returns a MySQL timestamp.
func (d *DBW) MySQLTimestamp(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// GoTimestamp returns a Go timestamp.
func (d *DBW) GoTimestamp(s string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", s)
}

// Exists returns the proper ID and other values based on the query results.
func (d *DBW) Exists(err error, ID int64) (bool, int64, error) {
	if err == nil {
		return true, ID, nil
	} else if err == sql.ErrNoRows {
		return false, 0, nil
	}
	return false, 0, err
}

// Error will return nil if the error is sql.ErrNoRows.
func (d *DBW) Error(err error) error {
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

// AffectedRows returns the number of rows affected by the query.
func (d *DBW) AffectedRows(result sql.Result) int {
	if result == nil {
		return 0
	}

	// If successful, get the number of affected rows.
	count, err := result.RowsAffected()
	if err != nil {
		return 0
	}

	return int(count)
}

// PaginatedResults returns the paginated results of a query.
func (d *DBW) PaginatedResults(i interface{}, fn func() (interface{}, int,
	error)) (int, error) {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Ptr {
		return 0, errors.New("must pass a pointer, not a value")
	}

	results, d2, d3 := fn()
	if results == nil {
		return d2, d3
	}

	arrPtr := reflect.ValueOf(i)
	value := arrPtr.Elem()
	itemPtr := reflect.ValueOf(results)
	value.Set(itemPtr)

	return d2, d3
}

// RecordExists returns the ID if a record exists.
func (d *DBW) RecordExists(fn func() (exists bool, ID int64, err error)) (
	exists bool, ID int64, err error) {
	return fn()
}

// AddRecord returns the ID if the record is created.
func (d *DBW) AddRecord(fn func() (ID int64, err error)) (ID int64, err error) {
	return fn()
}

// ExecQuery returns an error if the query failed.
func (d *DBW) ExecQuery(fn func() (err error)) (err error) {
	return fn()
}
