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
	component.IQuery `json:"-"`
	db               component.IDatabase

	ID        string     `db:"id" json:"id"`
	FirstName string     `db:"first_name" json:"first_name"`
	LastName  string     `db:"last_name" json:"last_name"`
	Email     string     `db:"email" json:"email"`
	Password  string     `db:"password" json:"password"`
	StatusID  uint8      `db:"status_id" json:"status_id"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
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
func (x *User) NewGroup() *TUserGroup {
	group := make(TUserGroup, 0)
	return &group
}

// UserGroup represents a group of users.
type TUserGroup []User

// Table returns the table name.
func (x TUserGroup) Table() string {
	return "user"
}

// PrimaryKey returns the primary key field.
func (x TUserGroup) PrimaryKey() string {
	return "id"
}

// *****************************************************************************
// Create
// *****************************************************************************

// Create adds a new user.
func (x *User) Create(firstName, lastName, email, password string) (string, error) {
	// Generate a UUID.
	uuid, err := securegen.UUID()
	if err != nil {
		return "", err
	}

	// Create the user.
	_, err = x.db.Exec(`
		INSERT INTO user
		(id, first_name, last_name, email, password, status_id)
		VALUES
		(?,?,?,?,?,?)
		`,
		uuid, firstName, lastName, email, password, 1)

	return uuid, err
}

// *****************************************************************************
// Update
// *****************************************************************************

// Update makes changes to one entity.
func (x *User) Update(ID, firstName, lastName, email, password string) (err error) {
	// Update the entity.
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
