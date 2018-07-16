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
func connect() (*sqlx.DB, error) {
	dbc := new(database.Connection)

	// Load the struct from environment variables.
	err := env.Unmarshal(dbc)
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

// Debug will return the SQL file.
func Debug(arr []*Changeset) {
	for _, cs := range arr {
		fmt.Printf("%v%v:%v\n", elementChangeset, cs.author, cs.id)
		fmt.Println(cs.Changes())
		fmt.Printf("%v%v\n", elementRollback, cs.Rollbacks())
		fmt.Println("--md5", cs.Checksum())
		break
	}

	fmt.Println("Total:", len(arr))
}
