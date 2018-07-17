package testutil

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"app/webapi/internal/basemigrate"
	"app/webapi/pkg/env"

	"app/webapi/pkg/database"
)

func setEnv() {
	os.Setenv("DB_HOSTNAME", "127.0.0.1")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USERNAME", "root")
	os.Setenv("DB_PASSWORD", "")
	os.Setenv("DB_DATABASE", "webapitest")
	os.Setenv("DB_PARAMETER", "parseTime=true&allowNativePasswords=true")
}

// ConnectDatabase returns a test database connection.
func ConnectDatabase(dbSpecificDB bool) *database.DBW {
	dbc := new(database.Connection)
	setEnv()

	err := env.Unmarshal(dbc)
	if err != nil {
		log.Println("DB ENV Error:", err)
	}

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
	_, err := db.Exec(`DROP DATABASE IF EXISTS webapitest`)
	if err != nil {
		fmt.Println(err)
	}
	_, err = db.Exec(`CREATE DATABASE webapitest DEFAULT CHARSET = utf8 COLLATE = utf8_unicode_ci`)
	if err != nil {
		fmt.Println(err)
	}
}

// LoadDatabase will set up the DB for the tests.
func LoadDatabase(t *testing.T) {
	ResetDatabase()

	err := basemigrate.Migrate("../../../../../migration/mysql-v0.sql", 0, false)
	if err != nil {
		log.Println("DB Error:", err)
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
