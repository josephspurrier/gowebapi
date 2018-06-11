package user

import (
	"database/sql"
	"time"

	"app/webapi/component"
	"app/webapi/pkg/uuid"
)

// TUser is user table that contains users.
type TUser struct {
	ID        string     `db:"id" json:"id"`
	FirstName string     `db:"first_name" json:"first_name" require:"true"`
	LastName  string     `db:"last_name" json:"last_name" require:"true"`
	Email     string     `db:"email" json:"email" require:"true"`
	Password  string     `db:"password" json:"password" require:"true"`
	StatusID  uint8      `db:"status_id" json:"status_id"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}

// *****************************************************************************
// Create
// *****************************************************************************

// Create adds a new user.
func Create(db component.IDatabase, firstName, lastName, email,
	password string) (string, error) {
	// Generate a UUID.
	uuid, err := uuid.Generate()
	if err != nil {
		return "", err
	}

	// Create the user.
	_, err = db.Exec(`
		INSERT INTO user
		(id, first_name, last_name, email, password, status_id)
		VALUES
		(?,?,?,?,?,?)
		`,
		uuid, firstName, lastName, email, password, 1)

	return uuid, err
}

// *****************************************************************************
// Read
// *****************************************************************************

// One returns one user with the matching ID.
func One(db component.IDatabase, ID string) (p TUser, exists bool, err error) {
	err = db.Get(&p, `
		SELECT * FROM user
		WHERE id = ?
		LIMIT 1`,
		ID)
	return p, (err != sql.ErrNoRows), db.Error(err)
}

// All returns all users.
func All(db component.IDatabase) ([]TUser, error) {
	result := make([]TUser, 0)
	err := db.Select(&result, `SELECT * FROM user`)
	return result, err
}

// ExistsEmail determines if a user exists by email.
func ExistsEmail(db component.IDatabase, s string) (exists bool, ID string,
	err error) {
	var p TUser
	err = db.Get(&p, `
		SELECT id FROM user
		WHERE email = ?
		LIMIT 1`,
		s)
	return db.ExistsString(err, p.ID)
}

// ExistsID determines if a user exists by ID.
func ExistsID(db component.IDatabase, s string) (exists bool, ID string,
	err error) {
	var p TUser
	err = db.Get(&p, `
		SELECT id FROM user
		WHERE id = ?
		LIMIT 1`,
		s)
	return db.ExistsString(err, p.ID)
}

// *****************************************************************************
// Update
// *****************************************************************************

// Update makes changes to one entity.
func Update(db component.IDatabase, ID, firstName, lastName, email,
	password string) (err error) {
	// Update the entity.
	_, err = db.Exec(`
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

// *****************************************************************************
// Delete
// *****************************************************************************

// Delete removes one entity.
func Delete(db component.IDatabase, ID string) (int, error) {
	result, err := db.Exec("DELETE FROM user WHERE id = ? LIMIT 1", ID)
	if err != nil {
		return 0, err
	}

	return db.AffectedRows(result), err
}

// DeleteAll removes all entities.
func DeleteAll(db component.IDatabase) (int, error) {
	result, err := db.Exec(`DELETE FROM user`)
	if err != nil {
		return 0, err
	}

	return db.AffectedRows(result), err
}
