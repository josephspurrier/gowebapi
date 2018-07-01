package store

import (
	"time"

	"app/webapi/component"
	"app/webapi/pkg/securegen"
)

// NewUser returns a new query object.
func NewUser(db component.IDatabase, q component.IQuery) *TUser {
	return &TUser{
		IQuery: q,
		db:     db,
	}
}

// TUser represents a user.
type TUser struct {
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
func (x *TUser) Table() string {
	return "user"
}

// PrimaryKey returns the primary key field.
func (x *TUser) PrimaryKey() string {
	return "id"
}

// NewGroup returns an empty group.
func (x *TUser) NewGroup() *TUserGroup {
	group := make(TUserGroup, 0)
	return &group
}

// TUserGroup represents a group of users.
type TUserGroup []TUser

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
func (x *TUser) Create(firstName, lastName, email, password string) (string, error) {
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
func (x *TUser) Update(ID, firstName, lastName, email, password string) (err error) {
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
