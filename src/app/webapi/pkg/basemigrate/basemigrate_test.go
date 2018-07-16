package basemigrate_test

import (
	"os"
	"testing"

	"app/webapi/internal/testutil"
	"app/webapi/pkg/basemigrate"

	"github.com/stretchr/testify/assert"
)

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
func TestMigration(t *testing.T) {
	setEnv()
	defer unsetEnv()

	testutil.PrepDatabase()
	db := testutil.ConnectDatabase(true)

	// Run migration.
	err := basemigrate.Migrate("testdata/success.sql", 0, false)
	assert.Nil(t, err)

	// Count the records.
	rows := 0
	err = db.Get(&rows, `SELECT count(*) from databasechangelog`)
	assert.Nil(t, err)
	assert.Equal(t, 3, rows)

	// Run migration again.
	err = basemigrate.Migrate("testdata/success.sql", 0, false)
	assert.Nil(t, err)

	// Remove all migrations.
	err = basemigrate.Reset("testdata/success.sql", 0, false)
	assert.Nil(t, err)

	rows = 0
	err = db.Get(&rows, `SELECT count(*) from databasechangelog`)
	assert.Nil(t, err)
	assert.Equal(t, 0, rows)

	// Remove all migrations again.
	err = basemigrate.Reset("testdata/success.sql", 0, false)
	assert.Nil(t, err)

	// Run 2 migrations.
	err = basemigrate.Migrate("testdata/success.sql", 2, false)
	assert.Nil(t, err)

	rows = 0
	err = db.Get(&rows, `SELECT count(*) from databasechangelog`)
	assert.Nil(t, err)
	assert.Equal(t, 2, rows)

	// Remove 1 migration.
	err = basemigrate.Reset("testdata/success.sql", 1, false)
	assert.Nil(t, err)

	rows = 0
	err = db.Get(&rows, `SELECT count(*) from databasechangelog`)
	assert.Nil(t, err)
	assert.Equal(t, 1, rows)
}

func TestMigrationFailDuplicate(t *testing.T) {
	setEnv()
	defer unsetEnv()

	testutil.PrepDatabase()
	db := testutil.ConnectDatabase(true)

	err := basemigrate.Migrate("testdata/fail-duplicate.sql", 0, false)
	assert.Contains(t, err.Error(), "checksum does not match")

	rows := 0
	err = db.Get(&rows, `SELECT count(*) from databasechangelog`)
	assert.Nil(t, err)
	assert.Equal(t, 2, rows)
}

func TestParse(t *testing.T) {
	setEnv()
	defer unsetEnv()

	arr, err := basemigrate.ParseFileArray("testdata/success.sql")
	assert.Nil(t, err)
	assert.Equal(t, 3, len(arr))

	m, err := basemigrate.ParseFileMap("testdata/success.sql")
	assert.Nil(t, err)
	assert.Equal(t, 3, len(m))
}
