package query

import "database/sql"

// recordExists returns if the record exists or not.
func recordExists(err error) (bool, error) {
	if err == nil {
		return true, nil
	} else if err == sql.ErrNoRows {
		return false, nil
	}
	return false, err
}

// recordExistsString returns the proper string is the record exists.
func recordExistsString(err error, s string) (bool, string, error) {
	if err == nil {
		return true, s, nil
	} else if err == sql.ErrNoRows {
		return false, "", nil
	}
	return false, "", err
}

// suppressNoRowsError will return nil if the error is sql.ErrNoRows.
func suppressNoRowsError(err error) error {
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

// affectedRows returns the number of rows affected by the query.
func affectedRows(result sql.Result) int {
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
