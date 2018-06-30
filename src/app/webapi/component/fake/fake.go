package fake

// TUser represents users.
/*type TUser struct {
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

// NewUser returns a new query object.
func NewUser(db component.IDatabase, q component.IQuery) *TUser {
	return &TUser{
		IQuery: q,
		db:     db,
	}
}*/

// FindOneByID will find the user by string ID.
/*func (x *TUser) FindOneByID(dest component.IRecord, ID string) (exists bool, err error) {
	err = x.db.Get(dest, fmt.Sprintf(`
		SELECT * FROM %s
		WHERE %s = ?
		LIMIT 1`, x.Table(), x.PrimaryKey()),
		ID)
	return (err != sql.ErrNoRows), x.db.Error(err)
}*/

//////////////////////

// NewDatabase returns a new database wrapper.
/*func NewDatabase(db *sqlx.DB) *DBW {
	return &DBW{
		db: db,
	}
}

// DBW is a database wrapper that provides helpful utilities.
type DBW struct {
	db *sqlx.DB
}

// Get using this DB.
// Any placeholder parameters are replaced with supplied args.
// An error is returned if the result set is empty.
func (d *DBW) Get(dest interface{}, query string, args ...interface{}) error {
	return d.db.Get(dest, query, args...)
}

// Error will return nil if the error is sql.ErrNoRows.
func (d *DBW) Error(err error) error {
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

// FindOneByID will find the user by string ID.
func (d *DBW) FindOneByID(dest component.IRecord, ID string) (exists bool, err error) {
	err = d.Get(dest, fmt.Sprintf(`
		SELECT * FROM %s
		WHERE %s = ?
		LIMIT 1`, dest.Table(), dest.PrimaryKey()),
		ID)
	return (err != sql.ErrNoRows), d.Error(err)
}*/
