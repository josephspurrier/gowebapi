# gowebapi

[![Go Report Card](https://goreportcard.com/badge/github.com/josephspurrier/gowebapi)](https://goreportcard.com/report/github.com/josephspurrier/gowebapi)
[![Build Status](https://travis-ci.org/josephspurrier/gowebapi.svg)](https://travis-ci.org/josephspurrier/gowebapi)
[![Coverage Status](https://coveralls.io/repos/github/josephspurrier/gowebapi/badge.svg?branch=master&timestamp=20180718-01)](https://coveralls.io/github/josephspurrier/gowebapi?branch=master)

[![Swagger Validator](http://online.swagger.io/validator?url=https://raw.githubusercontent.com/josephspurrier/gowebapi/master/src/app/webapi/swagger.json)](http://petstore.swagger.io/?url=https://raw.githubusercontent.com/josephspurrier/gowebapi/master/src/app/webapi/swagger.json)

## Testable Web API in Go with Swagger

This project demonstrates how to structure and build an API using the Go language without a framework. Only carefully chosen packages are included. Dredd is used to test the generated Swagger spec against the API to ensure it's correct.

**Older Version:** The previous version that was around for a while was 0.1-alpha. If you want to see that code, you can view the [tag](https://github.com/josephspurrier/gowebapi/releases/tag/0.1-alpha). The current version is a significant refactor that follows better practices.

**Note:** You cannot use `go get` with this repository. You should perform a `git clone` then set your GOPATH to the folder that git clone created called `gowebapi`. This allows you to easily fork the repository and build your own applications without rewritting any import paths.

You must use Go 1.7 or newer because this project uses the http context.

## Quick Start with Docker Compose

You can build a docker image from this repository and set it up along with a MySQL container using docker compose.

```bash
# Create a docker image.
docker build -t webapi:latest .

# Launch MySQL and the webapi with docker compose.
docker-compose up

# Open your browser to http://localhost:8080 and you should be able to access the API.
# Try using the Swagger spec:
# http://petstore.swagger.io/?url=https://raw.githubusercontent.com/josephspurrier/gowebapi/master/src/app/webapi/swagger.json

# Shutdown the containers.
docker-compose down
```

## Manual Start

Use the following commands to start a MySQL container with Docker:

```bash
# Start MySQL without a password.
docker run -d --name=mysql57 -p 3306:3306 -e MYSQL_ALLOW_EMPTY_PASSWORD=yes mysql:5.7
# or start MySQL with a password.
docker run -d --name=mysql57 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=somepassword mysql:5.7

# Create the database via docker exec.
docker exec mysql57 sh -c 'exec mysql -uroot -e "CREATE DATABASE IF NOT EXISTS webapi DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;"'
# Or create the database manually.
CREATE DATABASE webapi DEFAULT CHARSET = utf8 COLLATE = utf8_unicode_ci;

# CD to the CLI tool.
cd src/app/webapi/cmd/cliapp

# Build the CLI tool.
go build

# Apply the database migrations without a password.
DB_USERNAME=root DB_HOSTNAME=127.0.0.1 DB_PORT=3306 DB_DATABASE=webapi ./cliapp migrate all ../../../../../migration/mysql-v0.sql
# or apply the database migrations with a password.
DB_USERNAME=root DB_PASSWORD=somepassword DB_HOSTNAME=127.0.0.1 DB_PORT=3306 DB_DATABASE=webapi ./cliapp migrate all ../../../../../migration/mysql-v0.sql
```

Using the database connection information above, follow the steps to set up the `config.json` file:

```bash
# Copy the config.json from the root of the project to the CLI app folder.
cp config.json src/app/webapi/cmd/webapi/config.json

# Edit the `Database` section so the connection information matches your MySQL instance.
# The database password is read from the `config.json` file, but is overwritten by the environment variable, `DB_PASSWORD`, if it is set.

# Generate a base64 encoded secret.
./cliapp generate

# Use the encoded secret above to replace the `JWT.Secret` value in the config.
```

Now you can start the API.

```bash
# CD to the webapi app folder.
cd src/app/webapi/cmd/webapi

# Build the app.
go build

# Run the app.
./webapi

# Open your browser to this URL to see the **welcome** message and status **OK**: http://localhost:8080/v1
```

To interact with the API, open your favorite REST client.

You'll need to authenticate with at http://localhost:8080/v1/auth before you can use any of the user endpoints. Once you have a token, add it to the request header with a name of `Authorization` and with a value of `Bearer {TOKEN HERE}`. To create a user, send a POST request to http://localhost:8080/v1/user with the following fields: first_name, last_name, email, and password.

Currently, only a Content-Type of `application/x-www-form-urlencoded` is supported when sending to the API.

## Available Endpoints

The following endpoints are available:

```
* POST   /v1/user           - Create a new user
* GET	 /v1/user/{user_id} - Retrieve a user by ID
* GET	 /v1/user           - Retrieve a list of all users
* PUT	 /v1/user/{user_id} - Update a user by ID
* DELETE /v1/user/{user_id} - Delete a user by ID
* DELETE /v1/user           - Delete all users
```

## Swagger

This projects uses [Swagger v2](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md) to document the API. The entire Swagger spec is generated from the code in this repository.

The Swagger UI linked back to this project can be viewed
[here](http://petstore.swagger.io/?url=https://raw.githubusercontent.com/josephspurrier/gowebapi/master/src/app/webapi/swagger.json).

The Swagger spec JSON file is available
[here](https://github.com/josephspurrier/gowebapi/blob/master/src/app/webapi/swagger.json).

### Install Swagger

This tool will generate the Swagger spec from annotations in the Go code. It will read the comments in the code and will pull types from structs.

```bash
go get github.com/go-swagger/go-swagger/cmd/swagger
```

### Generate Swagger Spec

```bash
# CD to the webapi folder.
cd src/app/webapi

# Generate the swagger spec.
swagger generate spec -o ./swagger.json

# Replace 'example' with 'x-example' in the swagger spec.
## MacOS
sed -i '' -e 's/example/x\-example/' ./swagger.json
## Linux
sed -i'' -e 's/example/x\-example/' ./swagger.json

# Validate the swagger spec.
swagger validate ./swagger.json

# Serve the spec for the browser.
swagger serve -F=swagger ./swagger.json
```

## Dredd

This projects uses [Dredd](https://github.com/apiaryio/dredd) to test the Swagger spec against the API. Since the Swagger spec is generated from annotations in the Go code, it's good to ensure there are no discrepancies.

The Go documentation for Dredd is [here](https://dredd.readthedocs.io/en/latest/hooks-go.html).

Sample output from Dredd is [here](https://github.com/josephspurrier/gowebapi/wiki/Dredd-Sample-Output).

### Install Dredd

You must have MySQL running for these tests to pass.

```bash
# Install dredd.
npm install -g dredd

# Get the goodman package for Go hooks.
go get github.com/snikch/goodman/cmd/goodman

# CD to the webapi folder.
cd src/app/webapi

# Copy the testdata/config.json to the current directory.
cp testdata/config.json ./config.json

# Build the hooks app to load the test data.
go build -o ./cmd/hooks/hooks app/webapi/cmd/hooks

# Start MySQL without a password.
docker run -d --name=mysql57 -p 3306:3306 -e MYSQL_ALLOW_EMPTY_PASSWORD=yes mysql:5.7

# Run a test with dredd.
dredd
```

## Vendoring

This project uses [dep](https://github.com/golang/dep). The `dep init` command was run from inside the `src/app/webapi` folder.

These packages are used in the project:
- MySQL Driver: [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
- SQL to Struct: [github.com/jmoiron/sqlx](https://github.com/jmoiron/sqlx)
- Routing: [github.com/matryer/way](https://github.com/matryer/way)
- Request Validation: [github.com/go-playground/validator](https://github.com/go-playground/validator)
- JSON Web Tokens (JWT): [github.com/dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)
- CLI and Flag Parser: [gopkg.in/alecthomas/kingpin.v2](https://gopkg.in/alecthomas/kingpin.v2)
- Password Hashing: [golang.org/x/crypto/bcrypt](https://golang.org/x/crypto/bcrypt)

## Folder Structure

All the Go code is inside the `src` folder. This allows you to easily fork this project to use and test it. You'll just need to set your GOPATH to the `gowebapi` folder after you do a `git clone` (don't do a `go get`, it will not work).

In the `src/app/webapi` folder, you see a few top level folders:
- **cmd** - contains the main function and a static folder for the favicon.
- **component** - contains sets of related endpoints and database code.
- **internal** - contains project specific packages with dependencies.
- **middleware** - contains http wrappers for logging and CORS.
- **model** - contains the files with JSON structs that will outputted by the API.
- **pkg** - contains generic packages without project specific dependencies - these can be safely moved to other projects without internal dependencies.
- **store** - contains the files with SQL used to query the database.

## Components

In the root of the `src/app/webapi/component` folder, you see:
- **core.go** - contains the dependencies shared by all the components: logger, database connection, request bind/validation, and the responses.
- **core_mock.go** - contains the mocked dependencies which can be used by tests to modify the mocked dependencies.
- **interface.go** - contains all the interfaces for the dependencies so you can easily mock out each one for testing purposes.

Inside each component, you see a `component.go` file which contains the main struct and all the routes. You'll also see individual files for each endpoint with Swagger annotations and the tests for each endpoint.

## Store

In the `store` folder, you see `user.go` which has the SQL queries. Notice how `IDatabase` and the `IQuery` are passed into each store. This provides a unified way to run database queries and also provides a base set of simple SQL queries so you don't have to rewrite them for every table:
- FindOneByID(dest query.IRecord, ID string) (found bool, err error)
- FindAll(dest query.IRecord) (total int, err error)
- ExistsByID(db query.IRecord, s string) (found bool, err error)
- ExistsByField(db query.IRecord, field string, value string) (found bool, ID string, err error)
- DeleteOneByID(dest query.IRecord, ID string) (affected int, err error)
- DeleteAll(dest query.IRecord) (affected int, err error)

This is not an ORM - it just provides you with a simple query builder. Since the struct has an anonymous field, `component.IQuery`, you can overwrite any of the functions.

For instance, to retrieve a single user from the database, you would use this code:

```go
// Create the store.
u := store.NewUser(p.DB, p.Q)

// Get a user.
exists, err := u.FindOneByID(u, req.UserID)
```

The code for the generic `FindOneByID()` is in the `pkg/query/query.go` file:

```go
// FindOneByID will find a record by string ID.
func (q *Q) FindOneByID(dest IRecord, ID string) (exists bool, err error) {
	err = q.db.Get(dest, fmt.Sprintf(`
		SELECT * FROM %s
		WHERE %s = ?
		LIMIT 1`, dest.Table(), dest.PrimaryKey()),
		ID)
	return recordExists(err)
}
```

If you wanted to change the query so it excludes deleted users, you could add a new function to the `store/user.go` file so it looks like this:

```go
// FindOneByID will find a record by string ID excluding deleted records.
func (x *User) FindOneByID(dest query.IRecord, ID string) (exists bool, err error) {
	err = x.db.Get(dest, fmt.Sprintf(`
		SELECT * FROM %s
		WHERE %s = ?
		AND deleted_at  IS NULL
		LIMIT 1`, dest.Table(), dest.PrimaryKey()),
		ID)

	if err == nil {
		return true, nil
	} else if err == sql.ErrNoRows {
		return false, nil
	}
	return false, err
}
```

This allows you to standardize on how to interact with your database models throughout the team.

## Endpoint HTTP Handlers

In order to make the endpoints error driven, all the http handler functions must return an `int` and an `error`. This allows error handling to be centralized in the `webapi.go` file by setting the `router.ServeHTTP` variable. You can see the routes in the `component/user/component.go` file:

```go
// Routes will set up the endpoints.
func (p *Endpoint) Routes(router component.IRouter) {
	router.Post("/v1/user", p.Create)
	router.Get("/v1/user/:user_id", p.Show)
	router.Get("/v1/user", p.Index)
	router.Put("/v1/user/:user_id", p.Update)
	router.Delete("/v1/user/:user_id", p.Destroy)
	router.Delete("/v1/user", p.DestroyAll)
}
```

The endpoints are separated into files under each component folder and they look like this:

```go
func (p *Endpoint) DestroyAll(w http.ResponseWriter, r *http.Request) (int, error) {
	// Create the store.
	u := store.NewUser(p.DB, p.Q)

	// Delete all items.
	count, err := u.DeleteAll(u)
	if err != nil {
		return http.StatusInternalServerError, err
	} else if count < 1 {
		return http.StatusBadRequest, errors.New("no users to delete")
	}

	return p.Response.OK(w, "users deleted")
}
```

## Request Validation

The `app/webapi/internal/bind` is a wrapper around the `github.com/go-playground/validator` package so it can validate structs. You can view the `user/create.go` file to see where the email validation and the required validation is specified in the tags:

```go
// swagger:parameters UserCreate
type request struct {
	// in: formData
	// Required: true
	FirstName string `json:"first_name" validate:"required"`
	// in: formData
	// Required: true
	LastName string `json:"last_name" validate:"required"`
	// in: formData
	// Required: true
	Email string `json:"email" validate:"required,email"`
	// in: formData
	// Required: true
	Password string `json:"password" validate:"required"`
}

// Request validation.
req := new(request)
if err := p.Bind.FormUnmarshal(req, r); err != nil {
	return http.StatusBadRequest, err
} else if err = p.Bind.Validate(req); err != nil {
	return http.StatusBadRequest, err
}
```

## Reflection

The `app/webapi/internal/bind` and the `app/webapi/pkg/structcopy` packages use reflection. The `bind` package will take the form parameters from the request object and map them to a struct. The `structcopy` package will copy the values from the SQL store structs and set the fields on the JSON model structs based on the JSON tags.

```go
func (p *Endpoint) Index(w http.ResponseWriter, r *http.Request) (int, error) {
	// Create the DB store.
	u := store.NewUser(p.DB, p.Q)

	// Get all items.
	results := make(store.UserGroup, 0)
	_, err := u.FindAll(&results)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Copy the items to the JSON model.
	arr := make([]model.UserIndexResponseData, 0)
	for _, u := range results {
		item := new(model.UserIndexResponseData)
		err = structcopy.ByTag(&u, "db", item, "json")
		if err != nil {
			return http.StatusInternalServerError, err
		}
		arr = append(arr, *item)
	}

	// Send the response.
	resp := new(model.UserIndexResponse)
	resp.Body.Status = http.StatusText(http.StatusOK)
	resp.Body.Data = arr
	return p.Response.JSON(w, resp.Body)
}
```

## Logging

You can disable logging on the server by setting an environment variable: `WEBAPI_LOG_LEVEL=none`

## Testing

All the tests use a database called: `webapitest`. The quickest way to get it set up is:

```bash
# Launch MySQL in docker container.
docker run -d --name=mysql57 -p 3306:3306 -e MYSQL_ALLOW_EMPTY_PASSWORD=yes mysql:5.7

# Create the database via docker exec.
docker exec mysql57 sh -c 'exec mysql -uroot -e "CREATE DATABASE IF NOT EXISTS webapitest DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;"'

# Or create the database manually.
CREATE DATABASE webapitest DEFAULT CHARSET = utf8 COLLATE = utf8_unicode_ci;
```

You can use these commands to run tests:

```bash
# CD to the folder.
cd src/app/webapi

# Test all the packages.
go test ./...

# Get coverage of all tests.
go test -coverpkg=all ./...

# Get the coverage map of the current folder.
go test -coverprofile cover.out && go tool cover -html=cover.out -o cover.html && open cover.html && sleep 5 && rm cover.html && rm cover.out

# Get the coverage map of all the packages.
go test -coverprofile cover.out ./... && go tool cover -html=cover.out -o cover.html && open cover.html && sleep 5 && rm cover.html && rm cover.out

# Get the total code coverage - this only takes into consideration packages that
# have a test file in them.
go test ./... -coverprofile cover.out; go tool cover -func cover.out
```

## Conventions

Rules for mapping HTTP methods to CRUD:

```
POST   - Create (add record into database)
GET    - Read (get record from the database)
PUT    - Update (edit record in the database)
DELETE - Delete (remove record from the database)
```

Rules for HTTP status codes:

```
* Create something            - 201 (Created)
* Read something              - 200 (OK)
* Update something            - 200 (OK)
* Delete something            - 200 (OK)
* Missing request information - 400 (Bad Request)
* Unauthorized operation      - 401 (Unauthorized)
* Any other error             - 500 (Internal Server Error)
```

## Resources

These are all good reads:

Custom HTTP Handlers:

- https://blog.golang.org/error-handling-and-go
- https://mwholt.blogspot.com/2015/05/handling-errors-in-http-handlers-in-go.html
- https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831

Package Layout:

- https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
