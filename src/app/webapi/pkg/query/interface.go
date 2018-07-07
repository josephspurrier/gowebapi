package query

import "database/sql"

// IDatabase provides data query capabilities.
type IDatabase interface {
	Get(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)
	Select(dest interface{}, query string, args ...interface{}) error
	QueryRowScan(dest interface{}, query string, args ...interface{}) error
}

// IRecord provides table information.
type IRecord interface {
	Table() string
	PrimaryKey() string
}
