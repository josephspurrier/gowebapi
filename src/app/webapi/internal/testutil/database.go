package testutil

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"app/webapi/internal/basemigrate"
	"app/webapi/pkg/database"
	"app/webapi/pkg/env"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func setEnv(unique string) {
	os.Setenv(unique+"DB_HOSTNAME", "127.0.0.1")
	os.Setenv(unique+"DB_PORT", "3306")
	os.Setenv(unique+"DB_USERNAME", "root")
	os.Setenv(unique+"DB_PASSWORD", "")
	os.Setenv(unique+"DB_DATABASE", "webapitest"+unique)
	os.Setenv(unique+"DB_PARAMETER", "parseTime=true&allowNativePasswords=true")
}

func unsetEnv(unique string) {
	os.Unsetenv(unique + "DB_HOSTNAME")
	os.Unsetenv(unique + "DB_PORT")
	os.Unsetenv(unique + "DB_USERNAME")
	os.Unsetenv(unique + "DB_PASSWORD")
	os.Unsetenv(unique + "DB_DATABASE")
	os.Unsetenv(unique + "DB_PARAMETER")
}

// connectDatabase returns a test database connection.
func connectDatabase(dbSpecificDB bool, unique string) *database.DBW {
	dbc := new(database.Connection)
	err := env.Unmarshal(dbc, unique)
	if err != nil {
		fmt.Println("DB ENV Error:", err)
	}

	connection, err := dbc.Connect(dbSpecificDB)
	if err != nil {
		fmt.Println("DB Error:", err)
	}

	dbw := database.New(connection)

	return dbw
}

// SetupDatabase will create the test database and set the environment
// variables.
func SetupDatabase() (*database.DBW, string) {
	unique := "T" + fmt.Sprint(rand.Intn(500))
	setEnv(unique)

	db := connectDatabase(false, unique)
	_, err := db.Exec(`DROP DATABASE IF EXISTS webapitest` + unique)
	if err != nil {
		fmt.Println("DB DROP SETUP Error:", err)
	}
	_, err = db.Exec(`CREATE DATABASE webapitest` + unique + ` DEFAULT CHARSET = utf8 COLLATE = utf8_unicode_ci`)
	if err != nil {
		fmt.Println("DB CREATE Error:", err)
	}

	return connectDatabase(true, unique), unique
}

// TeardownDatabase will destroy the test database and unset the environment
// variables.
func TeardownDatabase(unique string) {
	db := connectDatabase(false, unique)
	_, err := db.Exec(`DROP DATABASE IF EXISTS webapitest` + unique)
	if err != nil {
		fmt.Println("DB DROP TEARDOWN Error:", err)
	}

	unsetEnv(unique)
}

// LoadDatabase will set up the DB and apply migrations for the tests.
func LoadDatabase() (*database.DBW, string) {
	return LoadDatabaseFromFile("../../../../../migration/mysql-v0.sql", true)
}

// LoadDatabaseFromFile will set up the DB for the tests.
func LoadDatabaseFromFile(file string, usePrefix bool) (*database.DBW, string) {
	unique := ""
	var db *database.DBW

	if usePrefix {
		db, unique = SetupDatabase()
	} else {
		setEnv(unique)
		db = connectDatabase(false, unique)
		_, err := db.Exec(`DROP DATABASE IF EXISTS webapitest`)
		if err != nil {
			fmt.Println("DB DROP SETUP Error:", err)
		}
		_, err = db.Exec(`CREATE DATABASE webapitest DEFAULT CHARSET = utf8 COLLATE = utf8_unicode_ci`)
		if err != nil {
			fmt.Println("DB CREATE Error:", err)
		}

		db = connectDatabase(true, unique)
	}

	err := basemigrate.Migrate(file, unique, 0, false)
	if err != nil {
		log.Println("DB Error:", err)
	}

	return db, unique
}
