package fake_test

import (
	"log"
	"testing"

	"app/webapi/component/user"
	"app/webapi/pkg/database"
	"app/webapi/pkg/query"
)

// docker run --name=mysql57 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -e -d ct-mysql:5.7.18

func TestMain(t *testing.T) {
	dbc := new(database.Connection)
	dbc.Hostname = "127.0.0.1"
	dbc.Port = 3306
	dbc.Username = "root"
	dbc.Password = "password"
	dbc.Database = "webapi"
	dbc.Parameter = "parseTime=true&allowNativePasswords=true"

	connection, err := dbc.Connect(true)
	if err != nil {
		log.Println("DB Error:", err)
	}

	dbw := database.New(connection)

	q := query.New(dbw)

	u := user.NewUser(dbw, q)

	exists, err := u.FindOneByID(u, "1")

	log.Println("Exists:", exists)
	log.Println("Record:", u)
	log.Println("First Name:", u.FirstName)
	log.Println("Error:", err)

	arr := u.NewGroup()
	total, err := u.FindAll(arr)

	log.Println("Total:", total)
	log.Println("Arr:", arr)
	log.Println("Error:", err)
}
