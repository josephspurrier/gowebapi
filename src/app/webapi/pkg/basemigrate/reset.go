package basemigrate

import (
	"errors"
	"fmt"
	"strings"
)

// DBChangeset represents the database table records.
type DBChangeset struct {
	ID            string `db:"id"`
	Author        string `db:"author"`
	Filename      string `db:"filename"`
	OrderExecuted int    `db:"orderexecuted"`
}

// Reset will remove all migrations.
func Reset(filename string, verbose bool) (err error) {
	db, err := connect()
	if err != nil {
		return err
	}

	// Get the changesets in a map.
	m, err := ParseFileMap(filename)
	if err != nil {
		return err
	}

	// Get each changeset from the database.
	results := make([]DBChangeset, 0)
	err = db.Select(&results, `
		SELECT id, author, filename, orderexecuted
		FROM databasechangelog
		ORDER BY orderexecuted DESC;`)
	if err != nil {
		return err
	}

	if len(results) == 0 {
		if verbose {
			fmt.Println("No rollbacks to perform.")
			return nil
		}
	}

	// Loop through each changeset.
	for _, r := range results {
		id := fmt.Sprintf("%v:%v", r.Author, r.ID)

		cs, ok := m[id]
		if !ok {
			return errors.New("changeset is missing: " + id)
		}

		arrQueries := strings.Split(cs.Rollbacks(), ";")

		// Loop through each rollback.
		for _, q := range arrQueries {
			if len(q) == 0 {
				continue
			}

			// Execute the query.
			_, err = db.Exec(q)
			if err != nil {
				return fmt.Errorf("sql error on rollback %v:%v - %v", cs.author, cs.id, err.Error())
			}
		}

		// Delete the record.
		_, err = db.Exec(`
			DELETE FROM databasechangelog
			WHERE id = ? AND author = ? AND filename = ?
			LIMIT 1
			`, cs.id, cs.author, cs.filename)
		if err != nil {
			return err
		}

		if verbose {
			fmt.Printf("Rollback applied: %v:%v\n", cs.author, cs.id)
		}
	}

	return
}
