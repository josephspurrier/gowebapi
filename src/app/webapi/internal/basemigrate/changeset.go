package basemigrate

import (
	"strings"
)

// Changeset is a SQL changeset.
type Changeset struct {
	id          string
	author      string
	filename    string
	md5         string
	description string
	version     string

	change   []string
	rollback []string
}

// ParseHeader will parse the header information.
func (cs *Changeset) ParseHeader(line string) error {
	arr := strings.Split(line, ":")
	if len(arr) != 2 {
		return ErrInvalidHeader
	}

	cs.author = arr[0]
	cs.id = arr[1]

	return nil
}

// SetFileInfo will set the file information.
func (cs *Changeset) SetFileInfo(filename string, description string, version string) {
	cs.filename = filename
	cs.description = description
	cs.version = version
}

// AddRollback will add a rollback command.
func (cs *Changeset) AddRollback(line string) {
	if len(cs.rollback) == 0 {
		cs.rollback = make([]string, 0)
	}
	cs.rollback = append(cs.rollback, line)
}

// AddChange will add a change command.
func (cs *Changeset) AddChange(line string) {
	if len(cs.change) == 0 {
		cs.change = make([]string, 0)
	}
	cs.change = append(cs.change, line)
}

// Changes will return all the changes.
func (cs *Changeset) Changes() string {
	return strings.Join(cs.change, "\n")
}

// Rollbacks will return all the rollbacks.
func (cs *Changeset) Rollbacks() string {
	return strings.Join(cs.rollback, "\n")
}

// Checksum returns an MD5 checksum for the changeset.
func (cs *Changeset) Checksum() string {
	return md5sum(cs.Changes())
}
