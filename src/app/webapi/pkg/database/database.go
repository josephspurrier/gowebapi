package database

import (
	"database/sql"

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

// QueryRowScan returns a single result.
func (d *DBW) QueryRowScan(dest interface{}, query string, args ...interface{}) error {
	return d.db.QueryRow(query, args...).Scan(dest)
}

/*
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
}*/
