package main

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"testing"

	"app/webapi/internal/testutil"

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

func setEnv() {
	os.Setenv("DB_HOSTNAME", "127.0.0.1")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USERNAME", "root")
	os.Setenv("DB_PASSWORD", "")
	os.Setenv("DB_DATABASE", "webapitest")
	os.Setenv("DB_PARAMETER", "parseTime=true&allowNativePasswords=true")
}

func unsetEnv() {
	os.Unsetenv("DB_HOSTNAME")
	os.Unsetenv("DB_PORT")
	os.Unsetenv("DB_USERNAME")
	os.Unsetenv("DB_PASSWORD")
	os.Unsetenv("DB_DATABASE")
	os.Unsetenv("DB_PARAMETER")
}
func TestMigrationAll(t *testing.T) {
	setEnv()
	defer unsetEnv()

	testutil.ResetDatabase()
	db := testutil.ConnectDatabase(true)

	// Set the arguments.
	os.Args = make([]string, 4)
	os.Args[0] = "cliapp"
	os.Args[1] = "migrate"
	os.Args[2] = "all"
	os.Args[3] = "testdata/success.sql"

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
}

func TestMigrationReset(t *testing.T) {
	TestMigrationAll(t)

	setEnv()
	defer unsetEnv()

	db := testutil.ConnectDatabase(true)

	// Set the arguments.
	os.Args = make([]string, 4)
	os.Args[0] = "cliapp"
	os.Args[1] = "migrate"
	os.Args[2] = "reset"
	os.Args[3] = "testdata/success.sql"

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
}

func TestMigrationUp(t *testing.T) {
	setEnv()
	defer unsetEnv()

	testutil.ResetDatabase()
	db := testutil.ConnectDatabase(true)

	// Set the arguments.
	os.Args = make([]string, 5)
	os.Args[0] = "cliapp"
	os.Args[1] = "migrate"
	os.Args[2] = "up"
	os.Args[3] = "2"
	os.Args[4] = "testdata/success.sql"

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
}

func TestMigrationDown(t *testing.T) {
	TestMigrationUp(t)

	setEnv()
	defer unsetEnv()

	db := testutil.ConnectDatabase(true)

	// Set the arguments.
	os.Args = make([]string, 5)
	os.Args[0] = "cliapp"
	os.Args[1] = "migrate"
	os.Args[2] = "down"
	os.Args[3] = "1"
	os.Args[4] = "testdata/success.sql"

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
}
