package structcopy_test

import (
	"testing"
	"time"

	"app/webapi/pkg/structcopy"

	"github.com/stretchr/testify/assert"
)

type User struct {
	ID        string     `db:"id"`
	FirstName string     `db:"first_name"`
	LastName  string     `db:"last_name"`
	Email     string     `db:"email"`
	Password  string     `db:"password"`
	StatusID  uint8      `db:"status_id"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type UserJSON struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	StatusID  uint8  `json:"status_id"`
}

func TestCopySuccess(t *testing.T) {
	src := new(User)
	src.ID = "1"
	src.FirstName = "John"
	src.LastName = "Smith"
	src.Email = "jsmith@example.com"
	src.StatusID = 5

	dst := new(UserJSON)

	err := structcopy.ByTag(src, "db", dst, "json")
	assert.Nil(t, err)

	assert.Equal(t, src.ID, dst.ID)
	assert.Equal(t, src.FirstName, dst.FirstName)
	assert.Equal(t, src.LastName, dst.LastName)
	assert.Equal(t, src.Email, dst.Email)
	assert.Equal(t, src.StatusID, dst.StatusID)
}

func TestCopyFailNoPointerSrc(t *testing.T) {
	src := User{}
	src.ID = "1"
	src.FirstName = "John"
	src.LastName = "Smith"
	src.Email = "jsmith@example.com"
	src.StatusID = 5

	dst := new(UserJSON)

	err := structcopy.ByTag(src, "db", dst, "json")
	assert.Contains(t, err.Error(), "src type not pointer")

	assert.NotEqual(t, src.ID, dst.ID)
	assert.NotEqual(t, src.FirstName, dst.FirstName)
	assert.NotEqual(t, src.LastName, dst.LastName)
	assert.NotEqual(t, src.Email, dst.Email)
	assert.NotEqual(t, src.StatusID, dst.StatusID)
}

func TestCopyFailNoPointerDst(t *testing.T) {
	src := new(User)
	src.ID = "1"
	src.FirstName = "John"
	src.LastName = "Smith"
	src.Email = "jsmith@example.com"
	src.StatusID = 5

	dst := UserJSON{}

	err := structcopy.ByTag(src, "db", dst, "json")
	assert.Contains(t, err.Error(), "dst type not pointer")

	assert.NotEqual(t, src.ID, dst.ID)
	assert.NotEqual(t, src.FirstName, dst.FirstName)
	assert.NotEqual(t, src.LastName, dst.LastName)
	assert.NotEqual(t, src.Email, dst.Email)
	assert.NotEqual(t, src.StatusID, dst.StatusID)
}

type UserDifferentType struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	StatusID  string `json:"status_id"`
}

func TestCopyDifferentType(t *testing.T) {
	src := new(User)
	src.ID = "1"
	src.FirstName = "John"
	src.LastName = "Smith"
	src.Email = "jsmith@example.com"
	src.StatusID = 5

	dst := new(UserDifferentType)

	err := structcopy.ByTag(src, "db", dst, "json")
	assert.Contains(t, err.Error(), "field types do not match")

	assert.Equal(t, src.ID, dst.ID)
	assert.Equal(t, src.FirstName, dst.FirstName)
	assert.Equal(t, src.LastName, dst.LastName)
	assert.Equal(t, src.Email, dst.Email)
	assert.NotEqual(t, src.StatusID, dst.StatusID)
}

func TestCopyFailStringSrc(t *testing.T) {
	src := ""
	dst := UserJSON{}

	err := structcopy.ByTag(src, "db", dst, "json")
	assert.Contains(t, err.Error(), "src type not pointer")
}

func TestCopyFailPointerStringSrc(t *testing.T) {
	src := ""
	dst := UserJSON{}

	err := structcopy.ByTag(&src, "db", dst, "json")
	assert.Contains(t, err.Error(), "src type not struct")
}

func TestCopyFailStringDst(t *testing.T) {
	src := new(User)
	dst := 0

	err := structcopy.ByTag(src, "db", dst, "json")
	assert.Contains(t, err.Error(), "dst type not pointer")
}

func TestCopyFailPointerStringDst(t *testing.T) {
	src := new(User)
	dst := 0

	err := structcopy.ByTag(src, "db", &dst, "json")
	assert.Contains(t, err.Error(), "dst type not struct")
}
