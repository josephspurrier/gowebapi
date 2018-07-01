package testdb

import (
	"app/webapi/component"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

// SetupTest will set up the DB for the testss.
func SetupTest(t *testing.T) {
	db := component.TestDatabase(false)
	db.Exec(`DROP DATABASE IF EXISTS webapitest`)
	db.Exec(`CREATE DATABASE webapitest DEFAULT CHARSET = utf8 COLLATE = utf8_unicode_ci`)

	db = component.TestDatabase(true)
	b, err := ioutil.ReadFile("../../../../../migration/tables-only.sql")
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

	//exit := m.Run()
	//db.Exec(`DROP DATABASE IF EXISTS webapitest`)
	//os.Exit(exit)
}
