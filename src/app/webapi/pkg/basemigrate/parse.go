package basemigrate

import (
	"bufio"
	"io"
	"path"
	"strings"
)

// parse will split the SQL migration into pieces.
func parse(r io.Reader, filename string) ([]*Changeset, error) {
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

		// Start recording the changeset.
		if strings.HasPrefix(line, elementChangeset) {
			// Create a new changeset.
			cs := new(Changeset)
			cs.ParseHeader(strings.TrimLeft(line, elementChangeset))
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
			cs.AddRollback(strings.TrimLeft(line, elementRollback))
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
