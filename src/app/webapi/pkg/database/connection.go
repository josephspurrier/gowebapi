package database

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/jmoiron/sqlx"
)

// Connection holds the details for the MySQL connection.
type Connection struct {
	Username  string `json:"Username" env:"DB_USERNAME"`
	Password  string `json:"Password" env:"DB_PASSWORD"`
	Database  string `json:"Database" env:"DB_DATABASE"`
	Charset   string `json:"Charset" env:"DB_CHARSET"`
	Collation string `json:"Collation" env:"DB_COLLATION"`
	Hostname  string `json:"Hostname" env:"DB_HOSTNAME"`
	Port      int    `json:"Port" env:"DB_PORT"`
	Parameter string `json:"Parameter" env:"DB_PARAMETER"`
}

// *****************************************************************************
// Database Handling
// *****************************************************************************

// Connect to the database.
func (c Connection) Connect(specificDatabase bool) (*sqlx.DB, error) {
	// Connect to database and ping
	return sqlx.Connect("mysql", c.dsn(specificDatabase))
}

// Create a new database.
func (c Connection) Create(sql *sqlx.DB) error {
	// Set defaults
	ci := c.setDefaults()

	// Create the database
	_, err := sql.Exec(fmt.Sprintf(`CREATE DATABASE %v
				DEFAULT CHARSET = %v
				COLLATE = %v
				;`, ci.Database,
		ci.Charset,
		ci.Collation))
	return err
}

// Drop a database.
func (c Connection) Drop(sql *sqlx.DB) error {
	// Drop the database
	_, err := sql.Exec(fmt.Sprintf(`DROP DATABASE %v;`, c.Database))
	return err
}

// *****************************************************************************
// MySQL Specific
// *****************************************************************************

// DSN returns the Data Source Name.
func (c Connection) dsn(includeDatabase bool) string {
	// Set defaults
	ci := c.setDefaults()

	// Build parameters
	param := ci.Parameter

	// If parameter is specified, add a question mark
	// Don't add one if a question mark is already there
	if len(ci.Parameter) > 0 && !strings.HasPrefix(ci.Parameter, "?") {
		param = "?" + ci.Parameter
	}

	// Add collation
	if !strings.Contains(param, "collation") {
		if len(param) > 0 {
			param += "&collation=" + ci.Collation
		} else {
			param = "?collation=" + ci.Collation
		}
	}

	// Add charset
	if !strings.Contains(param, "charset") {
		if len(param) > 0 {
			param += "&charset=" + ci.Charset
		} else {
			param = "?charset=" + ci.Charset
		}
	}

	// Example: root:password@tcp(localhost:3306)/test
	s := fmt.Sprintf("%v:%v@tcp(%v:%d)/%v", ci.Username, ci.Password, ci.Hostname, ci.Port, param)

	if includeDatabase {
		s = fmt.Sprintf("%v:%v@tcp(%v:%d)/%v%v", ci.Username, ci.Password, ci.Hostname, ci.Port, ci.Database, param)
	}

	return s
}

// setDefaults sets the charset and collation if they are not set.
func (c Connection) setDefaults() Connection {
	ci := c

	if len(ci.Charset) == 0 {
		ci.Charset = "utf8mb4"
	}
	if len(ci.Collation) == 0 {
		ci.Collation = "utf8mb4_unicode_ci"
	}

	return ci
}
