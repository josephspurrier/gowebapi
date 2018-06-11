# gowebapi

[![Go Report Card](https://goreportcard.com/badge/github.com/josephspurrier/gowebapi)](https://goreportcard.com/report/github.com/josephspurrier/gowebapi)

This project demonstrates how to structure and build an API using the Go language without a framework.
The API is still a work-in-progress, but it's designed to be easy to troubleshoot and easy to modify.
Everyone structures their API differently, but ultimately consistency is key. The more
consistent your API is, the easier it will be for other people to interact with it.

**Older Version:** The previous version that was around for a while was
0.1-alpha. If you want to see that code, you can view the
[tag](https://github.com/josephspurrier/gowebapi/releases/tag/0.1-alpha).
The current version is significant refactor following better practices.

You cannot use `go get` with this repository. You should perform a `git clone`
then set your GOPATH to the folder that git clone created called `gowebapi`.
This allows you to easily fork the repository and build your own applications
without rewritting any import paths.

You must use Go 1.7 or newer because it uses the http context.

## Swagger

This projects uses Swagger v2 to document the API:
https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md

### Install Swagger

This tool will generate the Swagger spec from annotations in the Go code. It
will read the comments in the code and will pull types from structs.

```bash
go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

### Generate Swagger Spec

```bash
# Change to the proper directory.
cd src/app/webapi

# Generate the swagger spec.
swagger generate spec -o ./swagger.json

# Serve the spec for the browser.
swagger serve -F=swagger ./swagger.json
```

## Vendoring

This project uses [dep](https://github.com/golang/dep). The `dep init` command
was run from inside the `src/app/webapi` folder.

The packages used in this project are:
- MySQL Driver: github.com/go-sql-driver/mysql
- SQL to Struct: github.com/jmoiron/sqlx
- Routing: github.com/matryer/way
- Request Validation: github.com/go-playground/validator

## Quick Start with MySQL

Start MySQL and import `migration/mysql.sql` to create the database and tables.

Copy `config.json` to `src/app/webapi/cmd/webapi/config.json` and edit the
**Database** section so the connection information matches your MySQL instance.

Build and run from the root directory. Open your REST client to:
http://localhost. You should see the **ok** message and status 200.

To create a user, send a POST request to http://localhost/user with the
following fields: first_name, last_name, email, and password.

Currently, only a Content-Type of `application/x-www-form-urlencoded` is
supported.

## Available Endpoints

The following endpoints are available:

```
* POST   http://localhost/v1/user           - Create a new user
* GET	 http://localhost/v1/user/{user_id} - Retrieve a user by ID
* GET	 http://localhost/v1/user           - Retrieve a list of all users
* PUT	 http://localhost/v1/user/{user_id} - Update a user by ID
* DELETE http://localhost/v1/user/{user_id} - Delete a user by ID
* DELETE http://localhost/v1/user           - Delete all users
```

## Interesting Files

The files that are the most interesting are:

* [src/app/webapi/controller/user.go](https://github.com/josephspurrier/gowebapi/blob/master/src/app/webapi/controller/user.go) - Controller with the routes for /users
* [src/app/webapi/model/user/user.go](https://github.com/josephspurrier/gowebapi/blob/master/src/app/webapi/model/user/user.go) - Model with all the MySQL logic
* [src/app/webapi/pkg/response/response.go](https://github.com/josephspurrier/gowebapi/blob/master/src/app/webapi/pkg/response/response.go) - Handles the output of JSON in 3 different formats

## Rules for Consistency

Rules for mapping HTTP methods to CRUD:

```
POST   - Create (add record into database)
GET    - Read (get record from the database)
PUT    - Update (edit record in the database)
DELETE - Delete (remove record from the database)
```

Rules for status codes:

```
* Create something - 201 (Created)
* Read something - 200 (OK)
* Update something - 200 (OK)
* Delete something - 200 (OK)
* Create but missing info - 400 (Bad Request)
* Any other error - 500 (Internal Server Error)
```

Rules for messages:

```
* 201 - item created
* 200 - item found; no items to find; items deleted; no items to delete; etc
* 400 - [field] is missing; [field] needs to be type: [type]
* 500 - an error occurred, please try again later (should also log error because it's a programming or server issue)
```