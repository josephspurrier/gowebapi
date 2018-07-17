package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"app/webapi/component"
	"app/webapi/component/user"
	"app/webapi/internal/testutil"
	"app/webapi/pkg/router"
	"app/webapi/store"

	"github.com/snikch/goodman/hooks"
	trans "github.com/snikch/goodman/transaction"
)

/*
Example transaction.

&transaction.Transaction{
	Id:"POST (400) /v1/user"
	Name:"user > /v1/user > Create a user. > 400 > application/json"
	Host:"127.0.0.1"
	Port:"8080"
	Protocol:"http:"
	FullPath:"/v1/user"
	Request:(*struct {
		Body string "json:\"body,omitempty\"";
		Headers map[string]interface {} "json:\"headers,omitempty\"";
		URI string "json:\"uri,omitempty\"";
		Method string "json:\"method,omitempty\"" })(0xc420150780),
	Expected:(*struct { StatusCode string "json:\"statusCode,omitempty\"";
		Body string "json:\"body,omitempty\"";
		Headers map[string]interface {} "json:\"headers,omitempty\"";
		Schema *json.RawMessage "json:\"bodySchema,omitempty\"" })(0xc4201464e0),
	Real:(*struct { Body string "json:\"body\"";
		Headers map[string]interface {} "json:\"headers\"";
		StatusCode int "json:\"statusCode\"" })(nil),
	Origin:(*json.RawMessage)(0xc4201584a0),
	Test:(*json.RawMessage)(nil),
	Results:(*json.RawMessage)(nil),
	Skip:true, Fail:interface {}(nil),
	TestOrder:[]string(nil)}
*/

// Response returns 200.
type response struct {
	// in: body
	Body struct {
		// Required: true
		Status string `json:"status"`
		// Required: true
		Data struct {
			// Required: true
			Token string `json:"token"`
		} `json:"data"`
	}
}

func main() {
	h := hooks.NewHooks()
	server := hooks.NewServer(hooks.NewHooksRunner(h))
	token := ""

	h.BeforeAll(func(t []*trans.Transaction) {
		// Get the auth token.
		r, err := http.Get(fmt.Sprintf("%v//%v:%v/v1/auth", t[0].Protocol, t[0].Host, t[0].Port))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Decode the response.
		rs := new(response)
		err = json.NewDecoder(r.Body).Decode(&rs.Body)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		token = rs.Body.Data.Token
	})

	h.BeforeEach(func(t *trans.Transaction) {
		// Set the Authorization header.
		t.Request.Headers["Authorization"] = "Bearer " + token

		// Load the database with test data.
		db, _ := testutil.LoadDatabaseFromFile("../../../migration/mysql-v0.sql", false)
		core, _ := component.NewCoreMock(db)

		mux := router.New()
		user.New(core).Routes(mux)

		// Create a new user.
		u := store.NewUser(core.DB, core.Q)
		id1, err := u.Create("John", "Smith", "jsmith@example.com", "password")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Change the email to a real email.
		if strings.Contains(t.Request.Body, "email") {
			u, err := url.ParseQuery(t.Request.Body)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			u.Set("email", "jsmith2@example.com")
			t.Request.Body = u.Encode()
		}

		// Update the URL for the requests so they have the ID.
		if t.Request.URI == "/v1/user/USERID" {
			t.FullPath = "/v1/user/" + id1
		}

		//testutil.TeardownDatabase(unique)
	})

	if false {
		h.BeforeAll(func(t []*trans.Transaction) {
			fmt.Println("before all modification")
		})
		h.BeforeEach(func(t *trans.Transaction) {
			fmt.Println("before each modification")
		})
		h.Before("user > /v1/user/{user_id} > Return one user.", func(t *trans.Transaction) {
			fmt.Println("before modification")
		})
		h.BeforeEachValidation(func(t *trans.Transaction) {
			fmt.Println("before each validation modification")
		})
		h.BeforeValidation("/message > GET", func(t *trans.Transaction) {
			fmt.Println("before validation modification")
		})
		h.After("/message > GET", func(t *trans.Transaction) {
			fmt.Println("after modification")
		})
		h.AfterEach(func(t *trans.Transaction) {
			fmt.Println("after each modification")
		})
		h.AfterAll(func(t []*trans.Transaction) {
			fmt.Println("after all modification")
		})
	}

	server.Serve()
	defer server.Listener.Close()
}
