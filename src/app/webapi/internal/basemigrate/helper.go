package basemigrate

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"

	"app/webapi/pkg/database"
	"app/webapi/pkg/env"

	"github.com/jmoiron/sqlx"
)

// connect will connect to the database.
func connect(prefix string) (*sqlx.DB, error) {
	dbc := new(database.Connection)

	// Load the struct from environment variables.
	err := env.Unmarshal(dbc, prefix)
	if err != nil {
		return nil, err
	}

	return dbc.Connect(true)
}

// md5sum will return a checksum from a string.
func md5sum(s string) string {
	h := md5.New()
	r := bytes.NewReader([]byte(s))
	_, _ = io.Copy(h, r)
	return fmt.Sprintf("%x", h.Sum(nil))
}
