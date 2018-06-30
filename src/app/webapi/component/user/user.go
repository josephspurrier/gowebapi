package user

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
	component.IQuery
	db component.IDatabase

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

//TODO: This should be used as an example.
// FindOneByID will find the user by string ID.
/*func (x *TUser) FindOneByID(dest component.IRecord, ID string) (exists bool, err error) {
	ID = "2"
	err = x.db.Get(dest, fmt.Sprintf(`
		SELECT * FROM %s
		WHERE %s = ?
		LIMIT 1`, x.Table(), x.PrimaryKey()),
		ID)
	return (err != sql.ErrNoRows), x.db.Error(err)
}*/

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
// Read
// *****************************************************************************

// One returns one user with the matching ID.
/*func One(db component.IDatabase, ID string) (p TUser, exists bool, err error) {
	err = db.Get(&p, `
		SELECT * FROM user
		WHERE id = ?
		LIMIT 1`,
		ID)
	return p, (err != sql.ErrNoRows), db.Error(err)
}*/

// All returns all users.
/*func All(db component.IDatabase) (result []TUser, total int, err error) {
	result = make([]TUser, 0)

	err = db.Get(&total, `
		SELECT COUNT(DISTINCT id)
		FROM user
		WHERE deleted_at IS NULL;`)
	if err != nil {
		return result, total, db.Error(err)
	}

	err = db.Select(&result, `SELECT * FROM user`)
	return result, total, err
}*/

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
