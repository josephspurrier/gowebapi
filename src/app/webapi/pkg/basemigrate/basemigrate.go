package basemigrate

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	sqlChangelog = `CREATE TABLE IF NOT EXISTS databasechangelog (
	id varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
	author varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
	filename varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
	dateexecuted datetime NOT NULL,
	orderexecuted int(11) NOT NULL,
	md5sum varchar(35) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
	description varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
	tag varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
	version varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`
)

const (
	appVersion       = "1.0"
	elementChangeset = "--changeset "
	elementRollback  = "--rollback "
)

var (
	// ErrInvalidHeader is when the changeset header is invalid.
	ErrInvalidHeader = errors.New("invalid changeset header")
	// ErrInvalidFormat is when a changeset is not found.
	ErrInvalidFormat = errors.New("invalid changeset format")
)

// ParseFile will parse a file into changesets.
func ParseFile(filename string) ([]*Changeset, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return parse(f, filename)
}

// Migrate will migrate a file.
func Migrate(filename string, verbose bool) (err error) {
	db, err := connect()
	if err != nil {
		return err
	}

	// Create the DATABASECHANGELOG.
	_, err = db.Exec(sqlChangelog)

	// Get the changesets.
	arr, err := ParseFile(filename)
	if err != nil {
		return err
	}

	// Loop through each changeset.
	for _, cs := range arr {
		checksum := ""
		newChecksum := cs.Checksum()

		// Determine if the changeset was already applied.
		// Count the number of rows.
		err = db.Get(&checksum, `SELECT md5sum
			FROM databasechangelog
			WHERE id = ?
			AND author = ?`, cs.id, cs.author)
		if err == nil {
			// Determine if the checksums match.
			if checksum != newChecksum {
				return fmt.Errorf("checksum does not match - existing changeset %v:%v has checksum %v, but new changeset has checksum %v",
					cs.author, cs.id, checksum, newChecksum)
			}

			if verbose {
				fmt.Printf("Changeset already applied: %v:%v\n", cs.author, cs.id)
			}
			continue
		} else if err != nil && err != sql.ErrNoRows {
			return fmt.Errorf("internal error on changeset %v:%v - %v", cs.author, cs.id, err.Error())
		}

		arrQueries := strings.Split(cs.Changes(), ";")
		// Loop through each change.
		for _, q := range arrQueries {
			if len(q) == 0 {
				continue
			}

			// Execute the query.
			_, err = db.Exec(q)
			if err != nil {
				return fmt.Errorf("sql error on changeset %v:%v - %v", cs.author, cs.id, err.Error())
			}
		}

		// Count the number of rows.
		count := 0
		err = db.Get(&count, `SELECT COUNT(*) FROM databasechangelog`)
		if err != nil {
			return err
		}

		// Insert the record.
		_, err = db.Exec(`
			INSERT INTO databasechangelog
			(id,author,filename,dateexecuted,orderexecuted,md5sum,description,version)
			VALUES(?,?,?,NOW(),?,?,?,?)
			`, cs.id, cs.author, cs.filename, count+1, newChecksum, cs.description, cs.version)
		if err != nil {
			return err
		}

		if verbose {
			fmt.Printf("Changeset applied: %v:%v\n", cs.author, cs.id)
		}
	}

	return
}
