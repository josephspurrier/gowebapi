package store

import (
	"time"

	"app/webapi/component"
	"app/webapi/pkg/securegen"
)

// NewUser returns a new query object.
func NewUser(db component.IDatabase, q component.IQuery) *User {
	return &User{
		IQuery: q,
		db:     db,
	}
}

// User is a user of the system.
type User struct {
	component.IQuery
	db component.IDatabase

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

// Table returns the table name.
func (x *User) Table() string {
	return "user"
}

// PrimaryKey returns the primary key field.
func (x *User) PrimaryKey() string {
	return "id"
}

// NewGroup returns an empty group.
func (x *User) NewGroup() *UserGroup {
	group := make(UserGroup, 0)
	return &group
}

// UserGroup represents a group of users.
type UserGroup []User

// Table returns the table name.
func (x UserGroup) Table() string {
	return "user"
}

// PrimaryKey returns the primary key field.
func (x UserGroup) PrimaryKey() string {
	return "id"
}

// Create adds a new user.
func (x *User) Create(firstName, lastName, email, password string) (string, error) {
	uuid, err := securegen.UUID()
	if err != nil {
		return "", err
	}

	_, err = x.db.Exec(`
		INSERT INTO user
		(id, first_name, last_name, email, password, status_id)
		VALUES
		(?,?,?,?,?,?)
		`,
		uuid, firstName, lastName, email, password, 1)

	return uuid, err
}

// Update makes changes to a user.
func (x *User) Update(ID, firstName, lastName, email, password string) (err error) {
	_, err = x.db.Exec(`
		UPDATE user
		SET
			first_name = ?,
			last_name = ?,
			email = ?,
			password = ?
		WHERE id = ?
		`,
		firstName, lastName, email, password, ID)
	return
}
