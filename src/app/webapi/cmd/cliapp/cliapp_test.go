package main

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"testing"

	"app/webapi/internal/testutil"
	"app/webapi/pkg/database"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	// Set the arguments.
	os.Args = make([]string, 2)
	os.Args[0] = "cliapp"
	os.Args[1] = "generate"

	// Redirect stdout.
	backupd := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the application.
	main()

	// Get the output.
	w.Close()
	out, err := ioutil.ReadAll(r)
	assert.Nil(t, err)
	os.Stdout = backupd

	// Decode the output.
	b, err := base64.StdEncoding.DecodeString(string(out))
	assert.Nil(t, err)
	s := string(b)

	// Ensure the length is 32 bytes.
	assert.Equal(t, 32, len(s))
}

func TestMigrationAll(t *testing.T) {
	_, unique := migrateAll(t)
	testutil.TeardownDatabase(unique)
}

func migrateAll(t *testing.T) (*database.DBW, string) {
	db, unique := testutil.SetupDatabase()

	// Set the arguments.
	os.Args = make([]string, 6)
	os.Args[0] = "cliapp"
	os.Args[1] = "migrate"
	os.Args[2] = "all"
	os.Args[3] = "testdata/success.sql"
	os.Args[4] = "--envprefix"
	os.Args[5] = unique

	// Redirect stdout.
	backupd := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the application.
	main()

	// Get the output.
	w.Close()
	out, err := ioutil.ReadAll(r)
	assert.Nil(t, err)
	os.Stdout = backupd

	assert.Contains(t, string(out), "Changeset applied")

	// Count the records.
	rows := 0
	err = db.Get(&rows, `SELECT count(*) from databasechangelog`)
	assert.Nil(t, err)
	assert.Equal(t, 3, rows)

	return db, unique
}

func TestMigrationReset(t *testing.T) {
	db, unique := migrateAll(t)

	// Set the arguments.
	os.Args = make([]string, 6)
	os.Args[0] = "cliapp"
	os.Args[1] = "migrate"
	os.Args[2] = "reset"
	os.Args[3] = "testdata/success.sql"
	os.Args[4] = "--envprefix"
	os.Args[5] = unique

	// Redirect stdout.
	backupd := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the application.
	main()

	// Get the output.
	w.Close()
	out, err := ioutil.ReadAll(r)
	assert.Nil(t, err)
	os.Stdout = backupd

	assert.Contains(t, string(out), "Rollback applied")

	// Count the records.
	rows := 0
	err = db.Get(&rows, `SELECT count(*) from databasechangelog`)
	assert.Nil(t, err)
	assert.Equal(t, 0, rows)

	testutil.TeardownDatabase(unique)
}

func TestMigrationUp(t *testing.T) {
	_, unique := migrateUp(t)
	testutil.TeardownDatabase(unique)
}

func migrateUp(t *testing.T) (*database.DBW, string) {
	db, unique := testutil.SetupDatabase()

	// Set the arguments.
	os.Args = make([]string, 7)
	os.Args[0] = "cliapp"
	os.Args[1] = "migrate"
	os.Args[2] = "up"
	os.Args[3] = "2"
	os.Args[4] = "testdata/success.sql"
	os.Args[5] = "--envprefix"
	os.Args[6] = unique

	// Redirect stdout.
	backupd := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the application.
	main()

	// Get the output.
	w.Close()
	out, err := ioutil.ReadAll(r)
	assert.Nil(t, err)
	os.Stdout = backupd

	assert.Contains(t, string(out), "Changeset applied")

	// Count the records.
	rows := 0
	err = db.Get(&rows, `SELECT count(*) from databasechangelog`)
	assert.Nil(t, err)
	assert.Equal(t, 2, rows)

	return db, unique
}

func TestMigrationDown(t *testing.T) {
	db, unique := migrateUp(t)

	// Set the arguments.
	os.Args = make([]string, 7)
	os.Args[0] = "cliapp"
	os.Args[1] = "migrate"
	os.Args[2] = "down"
	os.Args[3] = "1"
	os.Args[4] = "testdata/success.sql"
	os.Args[5] = "--envprefix"
	os.Args[6] = unique

	// Redirect stdout.
	backupd := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the application.
	main()

	// Get the output.
	w.Close()
	out, err := ioutil.ReadAll(r)
	assert.Nil(t, err)
	os.Stdout = backupd

	assert.Contains(t, string(out), "Rollback applied")

	// Count the records.
	rows := 0
	err = db.Get(&rows, `SELECT count(*) from databasechangelog`)
	assert.Nil(t, err)
	assert.Equal(t, 1, rows)

	testutil.TeardownDatabase(unique)
}
