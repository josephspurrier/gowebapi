package query

import (
	"database/sql"
	"fmt"
)

// New returns a new query object.
func New(db IDatabase) *Q {
	return &Q{
		db: db,
	}
}

// Q is a database wrapper that provides helpful utilities.
type Q struct {
	db IDatabase
}

// IDatabase provides data query capabilities.
type IDatabase interface {
	Get(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)
	Select(dest interface{}, query string, args ...interface{}) error
	QueryRowScan(dest interface{}, query string, args ...interface{}) error

	Error(err error) error
	AffectedRows(result sql.Result) int

	ExistsString(err error, s string) (bool, string, error)
}

// IRecord provides table information.
type IRecord interface {
	Table() string
	PrimaryKey() string
}

// FindOneByID will find a record by string ID.
func (q *Q) FindOneByID(dest IRecord, ID string) (exists bool, err error) {
	err = q.db.Get(dest, fmt.Sprintf(`
		SELECT * FROM %s
		WHERE %s = ?
		LIMIT 1`, dest.Table(), dest.PrimaryKey()),
		ID)
	return (err != sql.ErrNoRows), q.db.Error(err)
}

// FindAll returns all users.
func (q *Q) FindAll(dest IRecord) (total int, err error) {
	//TODO: Add in something to handle soft deletes.
	//WHERE deleted_at IS NULL

	err = q.db.QueryRowScan(&total, fmt.Sprintf(`
		SELECT COUNT(DISTINCT %s)
		FROM %s
		`, dest.PrimaryKey(), dest.Table()))

	if err != nil {
		return total, q.db.Error(err)
	}

	err = q.db.Select(dest, fmt.Sprintf(`SELECT * FROM %s`, dest.Table()))
	return total, err
}

// *****************************************************************************
// Delete
// *****************************************************************************

// DeleteOneByID removes one record by ID.
func (q *Q) DeleteOneByID(dest IRecord, ID string) (affected int, err error) {
	result, err := q.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE %s = ? LIMIT 1",
		dest.Table(), dest.PrimaryKey()), ID)
	if err != nil {
		return 0, err
	}

	return q.db.AffectedRows(result), err
}

// DeleteAll removes all records.
func (q *Q) DeleteAll(dest IRecord) (affected int, err error) {
	result, err := q.db.Exec(fmt.Sprintf(`DELETE FROM %s`, dest.Table()))
	if err != nil {
		return 0, err
	}

	return q.db.AffectedRows(result), err
}

// *****************************************************************************
// Exists
// *****************************************************************************

// ExistsByID determines if a records exists by ID.
func (q *Q) ExistsByID(db IRecord, value string) (found bool, err error) {
	err = q.db.Get(db, fmt.Sprintf(`
		SELECT %s FROM %s
		WHERE %s = ?
		LIMIT 1`, db.PrimaryKey(), db.Table(), db.PrimaryKey()),
		value)
	return recordExists(err)
}

// ExistsByField determines if a records exists by a specified field and
// returns the ID.
func (q *Q) ExistsByField(db IRecord, field string, value string) (found bool, ID string, err error) {
	err = q.db.QueryRowScan(&ID, fmt.Sprintf(`
		SELECT %s FROM %s
		WHERE %s = ?
		LIMIT 1`, db.PrimaryKey(), db.Table(), field),
		value)

	//TODO: Add this to a function so it can be reused.
	if err == nil {
		return true, ID, nil
	} else if err == sql.ErrNoRows {
		return false, "", nil
	}
	return false, "", err

	/*err = q.db.Get(db, fmt.Sprintf(`
		SELECT %s FROM %s
		WHERE %s = ?
		LIMIT 1`, db.PrimaryKey(), db.Table(), field),
		value)
	return recordExists(err)*/
}

// recordExists returns if the record exists or not.
func recordExists(err error) (bool, error) {
	if err == nil {
		return true, nil
	} else if err == sql.ErrNoRows {
		return false, nil
	}
	return false, err
}
