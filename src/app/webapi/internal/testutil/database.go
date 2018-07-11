package testutil

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"app/webapi/pkg/database"
)

// ConnectDatabase returns a test database connection.
func ConnectDatabase(dbSpecificDB bool) *database.DBW {
	dbc := new(database.Connection)
	dbc.Hostname = "127.0.0.1"
	dbc.Port = 3306
	dbc.Username = "root"
	dbc.Password = ""
	dbc.Database = "webapitest"
	dbc.Parameter = "parseTime=true&allowNativePasswords=true"

	connection, err := dbc.Connect(dbSpecificDB)
	if err != nil {
		log.Println("DB Error:", err)
	}

	dbw := database.New(connection)

	return dbw
}

// ResetDatabase will drop and create the test database.
func ResetDatabase() {
	db := ConnectDatabase(false)
	db.Exec(`DROP DATABASE IF EXISTS webapitest`)
	db.Exec(`CREATE DATABASE webapitest DEFAULT CHARSET = utf8 COLLATE = utf8_unicode_ci`)
}

// LoadDatabase will set up the DB for the tests.
func LoadDatabase(t *testing.T) {
	ResetDatabase()

	db := ConnectDatabase(true)
	b, err := ioutil.ReadFile("../../../../../migration/tables-only.sql")
	if err != nil {
		t.Error(err)
	}

	// Split each statement.
	stmts := strings.Split(string(b), ";")
	for i, s := range stmts {
		if i == len(stmts)-1 {
			break
		}
		_, err = db.Exec(s)
		if err != nil {
			t.Error(err)
		}
	}
}

// LoadDatabaseFromFile will set up the DB for the tests.
func LoadDatabaseFromFile(file string) {
	ResetDatabase()

	db := ConnectDatabase(true)
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Split each statement.
	stmts := strings.Split(string(b), ";")
	for i, s := range stmts {
		if i == len(stmts)-1 {
			break
		}
		_, err = db.Exec(s)
		if err != nil {
			log.Println(err)
		}
	}
}
