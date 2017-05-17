package user

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"app/shared/database"
)

// Name of the table
const tableName = "user"

// Errors
var (
	ErrNoResult = errors.New("no result")
	ErrExists   = errors.New("already exists")
	ErrNotExist = errors.New("does not exist")
)

// Entity information
type Entity struct {
	ID        string    `db:"id" json:"id"`
	FirstName string    `db:"first_name" json:"first_name" require:"true"`
	LastName  string    `db:"last_name" json:"last_name" require:"true"`
	Email     string    `db:"email" json:"email" require:"true"`
	Password  string    `db:"password" json:"password" require:"true"`
	StatusID  uint8     `db:"status_id" json:"status_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt time.Time `db:"deleted_at" json:"deleted_at"`
}

// Group of entities
type Group []Entity

// New entity
func New() (*Entity, error) {
	var err error
	entity := &Entity{}

	// Set the default parameters
	entity.StatusID = 1
	entity.ID, err = database.UUID()
	// If error on UUID generation
	if err != nil {
		return entity, err
	}

	return entity, nil
}

// *****************************************************************************
// Create
// *****************************************************************************

// Create will add a new entity
func (u *Entity) Create() (int, error) {
	// Check for existing entity
	_, err := readOneByField("email", u.Email)

	// If entity exists
	if err != ErrNoResult {
		return 0, ErrExists
	}

	// Create the entity
	_, err = database.SQL.Exec(fmt.Sprintf(`
		INSERT INTO %v
		(id, first_name, last_name, email, password, status_id)
		VALUES
		(?,?,?,?,?,?)
		`, tableName),
		u.ID,
		u.FirstName,
		u.LastName,
		u.Email,
		u.Password,
		u.StatusID)

	// If error occurred error
	if err != nil {
		return 0, err
	}

	return 1, nil
}

// *****************************************************************************
// Read
// *****************************************************************************

// Read returns one entity with the matching ID
// If no result, it will return sql.ErrNoRows
func Read(ID string) (*Entity, error) {
	return readOneByField("id", ID)
}

// ReadAll returns all entities
func ReadAll() (Group, error) {
	var result Group
	err := database.SQL.Select(&result, fmt.Sprintf("SELECT * FROM %v", tableName))
	return result, err
}

// readOneByField returns the entity that matches the field value
// If no result, it will return ErrNoResult
func readOneByField(name string, value string) (*Entity, error) {
	result := &Entity{}
	err := database.SQL.Get(result, fmt.Sprintf("SELECT * FROM %v WHERE %v = ? LIMIT 1", tableName, name), value)
	if err == sql.ErrNoRows {
		err = ErrNoResult
	}
	return result, err
}

// readAllByField returns entities matching a field value
// If no result, it will return an empty group
func readAllByField(name string, value string) (Group, error) {
	var result Group
	err := database.SQL.Select(&result, fmt.Sprintf("SELECT * FROM %v WHERE %v = ?", tableName, name), value)
	return result, err
}

// *****************************************************************************
// Update
// *****************************************************************************

// Update makes changes to one entity
func (u *Entity) Update() (int, error) {
	// Check for existing entity
	_, err := readOneByField("id", u.ID)

	// If entity does NOT exists
	if err == ErrNoResult {
		return 0, ErrNotExist
	}

	// Update the entity
	_, err = database.SQL.Exec(fmt.Sprintf(`
		UPDATE %v SET
		first_name = ?,
		last_name = ?,
		email = ?,
		password = ?
		WHERE id = ?
		`, tableName),
		u.FirstName,
		u.LastName,
		u.Email,
		u.Password,
		u.ID)

	// If error occurred error
	if err != nil {
		return 0, err
	}

	return 1, nil
}

// *****************************************************************************
// Delete
// *****************************************************************************

// DeleteOne removes one entity
func Delete(ID string) (int, error) {
	result, err := database.SQL.Exec(fmt.Sprintf("DELETE FROM %v WHERE id = ? LIMIT 1", tableName), ID)
	if err != nil {
		return 0, err
	}

	return database.AffectedRows(result), err
}

// DeleteAll removes all entities
func DeleteAll() (int, error) {
	result, err := database.SQL.Exec(fmt.Sprintf("DELETE FROM %v", tableName))
	if err != nil {
		return 0, err
	}

	return database.AffectedRows(result), err
}

// deleteOneByField deletes an entity matching a field value
func deleteOneByField(name string, value string) (int, error) {
	result, err := database.SQL.Exec(fmt.Sprintf("DELETE FROM %v WHERE %v = ? LIMIT 1", tableName, name), value)
	if err != nil {
		return 0, err
	}

	return database.AffectedRows(result), err
}

// deleteAllByField deletes all entities matching a field value
func deleteAllByField(name string, value string) (int, error) {
	result, err := database.SQL.Exec(fmt.Sprintf("DELETE FROM %v WHERE %v = ?", tableName, name), value)
	if err != nil {
		return 0, err
	}

	return database.AffectedRows(result), err
}
