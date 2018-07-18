package basemigrate

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// parseFileToArray will parse a file into changesets.
func parseFileToArray(filename string) ([]*Changeset, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return parseToArray(f, filename)
}

// parseToArray will split the SQL migration into an ordered array.
func parseToArray(r io.Reader, filename string) ([]*Changeset, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	// Array of changesets.
	arr := make([]*Changeset, 0)

	for scanner.Scan() {
		// Get the line without leading or trailing spaces.
		line := strings.TrimSpace(scanner.Text())

		// Skip blank lines.
		if len(line) == 0 {
			continue
		}

		// Determine if the line is an `include`.
		if strings.HasPrefix(line, elementInclude) {
			// Load the file and add to the array.
			fp := strings.TrimPrefix(line, elementInclude)
			rfp := filepath.Join(filepath.Dir(filename), fp)
			cs, err := parseFileToArray(rfp)
			if err != nil {
				return nil, err
			}
			arr = append(arr, cs...)
			continue
		}

		// Start recording the changeset.
		if strings.HasPrefix(line, elementChangeset) {
			// Create a new changeset.
			cs := new(Changeset)
			cs.ParseHeader(strings.TrimPrefix(line, elementChangeset))
			cs.SetFileInfo(path.Base(filename), "sql", appVersion)
			arr = append(arr, cs)
			continue
		}

		// If the length of the array is 0, then the first changeset is missing.
		if len(arr) == 0 {
			return nil, ErrInvalidFormat
		}

		// Determine if the line is a rollback.
		if strings.HasPrefix(line, elementRollback) {
			cs := arr[len(arr)-1]
			cs.AddRollback(strings.TrimPrefix(line, elementRollback))
			continue
		}

		// Determine if the line is comment, ignore it.
		if strings.HasPrefix(line, "--") {
			continue
		}

		// Add the line as a changeset.
		cs := arr[len(arr)-1]
		cs.AddChange(line)
	}

	return arr, nil
}

// parseFileToMap will parse a file into a map.
func parseFileToMap(filename string) (map[string]Changeset, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return parseToMap(f, filename)
}

// parseToMap will parse a reader to a map.
func parseToMap(r io.Reader, filename string) (map[string]Changeset, error) {
	arr, err := parseToArray(r, filename)
	if err != nil {
		return nil, err
	}

	m := make(map[string]Changeset)

	for _, cs := range arr {
		id := fmt.Sprintf("%v:%v:%v", cs.author, cs.id, cs.filename)
		if _, found := m[id]; found {
			return nil, errors.New("Duplicate entry found: " + id)
		}

		m[id] = *cs
	}

	return m, nil
}
