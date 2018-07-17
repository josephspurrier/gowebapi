package basemigrate_test

import (
	"testing"

	"app/webapi/internal/basemigrate"
	"app/webapi/internal/testutil"

	"github.com/stretchr/testify/assert"
)

func TestMigration(t *testing.T) {
	testutil.ResetDatabase()
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
	testutil.ResetDatabase()
	db := testutil.ConnectDatabase(true)

	err := basemigrate.Migrate("testdata/fail-duplicate.sql", 0, false)
	assert.Contains(t, err.Error(), "checksum does not match")

	rows := 0
	err = db.Get(&rows, `SELECT count(*) from databasechangelog`)
	assert.Nil(t, err)
	assert.Equal(t, 2, rows)

	testutil.ResetDatabase()
}

func TestInclude(t *testing.T) {
	testutil.ResetDatabase()
	db := testutil.ConnectDatabase(true)

	// Run migration.
	err := basemigrate.Migrate("testdata/parent.sql", 0, false)
	assert.Nil(t, err)

	// Count the records.
	rows := 0
	err = db.Get(&rows, `SELECT count(*) from databasechangelog`)
	assert.Nil(t, err)
	assert.Equal(t, 3, rows)

	// Run migration again.
	err = basemigrate.Migrate("testdata/parent.sql", 0, false)
	assert.Nil(t, err)

	// Remove all migrations.
	err = basemigrate.Reset("testdata/parent.sql", 0, false)
	assert.Nil(t, err)

	rows = 0
	err = db.Get(&rows, `SELECT count(*) from databasechangelog`)
	assert.Nil(t, err)
	assert.Equal(t, 0, rows)

	// Remove all migrations again.
	err = basemigrate.Reset("testdata/parent.sql", 0, false)
	assert.Nil(t, err)

	// Run 2 migrations.
	err = basemigrate.Migrate("testdata/parent.sql", 2, false)
	assert.Nil(t, err)

	rows = 0
	err = db.Get(&rows, `SELECT count(*) from databasechangelog`)
	assert.Nil(t, err)
	assert.Equal(t, 2, rows)

	// Remove 1 migration.
	err = basemigrate.Reset("testdata/parent.sql", 1, false)
	assert.Nil(t, err)

	rows = 0
	err = db.Get(&rows, `SELECT count(*) from databasechangelog`)
	assert.Nil(t, err)
	assert.Equal(t, 1, rows)
}
