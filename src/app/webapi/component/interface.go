package component

import (
	"database/sql"
	"net/http"
	"time"

	"app/webapi/pkg/query"
)

// IDatabase provides data query capabilities.
type IDatabase interface {
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRowScan(dest interface{}, query string, args ...interface{}) error

	ExistsString(err error, s string) (bool, string, error)
	Error(err error) error
	AffectedRows(result sql.Result) int

	PaginatedResults(results interface{}, fn func() (results interface{}, total int, err error)) (total int, err error)
	RecordExistsInt(fn func() (exists bool, ID int64, err error)) (exists bool, ID int64, err error)
	RecordExistsString(fn func() (exists bool, ID string, err error)) (exists bool, ID string, err error)
	AddRecordInt(fn func() (ID int64, err error)) (ID int64, err error)
	AddRecordString(fn func() (ID string, err error)) (ID string, err error)
}

// IQuery provides default queries.
type IQuery interface {
	FindOneByID(dest query.IRecord, ID string) (found bool, err error)
	FindAll(dest query.IRecord) (total int, err error)
	ExistsByID(db query.IRecord, s string) (found bool, err error)
	ExistsByField(db query.IRecord, field string, value string) (found bool, ID string, err error)
	DeleteOneByID(dest query.IRecord, ID string) (affected int, err error)
	DeleteAll(dest query.IRecord) (affected int, err error)
}

// IRecord provides table information.
type IRecord interface {
	Table() string
	PrimaryKey() string
}

/*
// IQuery provides data query capabilities.
type IQuery interface {
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)

	ExistsString(err error, s string) (bool, string, error)
	Error(err error) error
	AffectedRows(result sql.Result) int
}*/

// IDatabase provides data query capabilities.
/*type IDatabase interface {
	//LastInsertID(r sql.Result, err error) (int64, error)
	//MySQLTimestamp(t time.Time) string
	//GoTimestamp(s string) (time.Time, error)

	//ExistsID(err error, ID int64) (bool, int64, error)

	PaginatedResults(results interface{}, fn func() (results interface{}, total int, err error)) (total int, err error)
	RecordExistsInt(fn func() (exists bool, ID int64, err error)) (exists bool, ID int64, err error)
	RecordExistsString(fn func() (exists bool, ID string, err error)) (exists bool, ID string, err error)
	AddRecordInt(fn func() (ID int64, err error)) (ID int64, err error)
	AddRecordString(fn func() (ID string, err error)) (ID string, err error)
	//ExecQuery(fn func() (err error)) (err error)
}*/

// ILogger provides logging capabilities.
type ILogger interface {
	//ControllerError(r *http.Request, err error, a ...interface{})
	//Fatalf(format string, v ...interface{})
	//Printf(format string, v ...interface{})
}

// IRouter provides routing capabilities.
type IRouter interface {
	Delete(path string, fn http.Handler)
	Get(path string, fn http.Handler)
	Head(path string, fn http.Handler)
	Options(path string, fn http.Handler)
	Patch(path string, fn http.Handler)
	Post(path string, fn http.Handler)
	Put(path string, fn http.Handler)
}

// IBind provides bind and validation for requests.
type IBind interface {
	FormUnmarshal(i interface{}, r *http.Request) (err error)
	Validate(s interface{}) error
}

// IResponse provides outputs for data.
type IResponse interface {
	Created(w http.ResponseWriter, recordID string) (int, error)
	Results(w http.ResponseWriter, body interface{}, data interface{}) (int, error)
	OK(w http.ResponseWriter, message string) (int, error)
}

// IToken provides outputs for the JWT.
type IToken interface {
	Generate(userID string, duration time.Duration) (string, error)
}
