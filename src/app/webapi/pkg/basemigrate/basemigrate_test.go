package basemigrate_test

import (
	"testing"

	"app/webapi/pkg/basemigrate"

	"github.com/stretchr/testify/assert"
)

func TestMigration(t *testing.T) {
	err := basemigrate.Migrate("testdata/success.sql", false)
	assert.Nil(t, err)
}

func TestMigrationFailDuplicate(t *testing.T) {
	err := basemigrate.Migrate("testdata/fail-duplicate.sql", false)
	assert.Contains(t, err.Error(), "checksum does not match")
}

func TestParse(t *testing.T) {
	arr, err := basemigrate.ParseFile("testdata/success.sql")
	assert.Nil(t, err)
	assert.Equal(t, 3, len(arr))

	//basemigrate.Debug(arr)

	/*for _, v := range arr {
		fmt.Println(v.Changes())
		fmt.Println(v.Rollbacks())
		fmt.Println("MD5:", v.Checksum())
		break
	}

	fmt.Println("Total:", len(arr))*/
}
