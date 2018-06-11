package component

import (
	"database/sql"
	"net/http"
	"time"
)

// IDatabase provides data query capabilities.
type IDatabase interface {
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)

	LastInsertID(r sql.Result, err error) (int64, error)
	MySQLTimestamp(t time.Time) string
	GoTimestamp(s string) (time.Time, error)
	Exists(err error, ID int64) (bool, int64, error)

	Error(err error) error
	AffectedRows(result sql.Result) int

	PaginatedResults(results interface{}, fn func() (results interface{}, total int, err error)) (total int, err error)
	RecordExists(fn func() (exists bool, ID int64, err error)) (exists bool, ID int64, err error)
	AddRecord(fn func() (ID int64, err error)) (ID int64, err error)
	ExecQuery(fn func() (err error)) (err error)
}

// ILogger provides logging capabilities.
type ILogger interface {
	ControllerError(r *http.Request, err error, a ...interface{})
	Fatalf(format string, v ...interface{})
	Printf(format string, v ...interface{})
}

// IRouter provides routing capabilities.
type IRouter interface {
	Delete(path string, fn http.HandlerFunc)
	Get(path string, fn http.HandlerFunc)
	Head(path string, fn http.HandlerFunc)
	Options(path string, fn http.HandlerFunc)
	Patch(path string, fn http.HandlerFunc)
	Post(path string, fn http.HandlerFunc)
	Put(path string, fn http.HandlerFunc)
}

// IBind provides bind and validation for requests.
type IBind interface {
	FormUnmarshal(i interface{}, r *http.Request) (err error)
	Validate(s interface{}) error
}

// IResponse provides outputs for data.
type IResponse interface {
	Send(w http.ResponseWriter, status http.ConnState, message string, count int, results interface{})
	SendError(w http.ResponseWriter, status http.ConnState, message string)
}
