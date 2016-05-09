# GoWebApi

[![Go Report Card](https://goreportcard.com/badge/github.com/josephspurrier/gowebapi)](https://goreportcard.com/report/github.com/josephspurrier/gowebapi)
[![GoDoc](https://godoc.org/github.com/josephspurrier/gowebapi?status.svg)](https://godoc.org/github.com/josephspurrier/gowebapi) 

Basic Web API in Go

This project demonstrates how to structure and build an API using the Go language without a framework. The API is still a work-in-progress, but it's designed to be easy to troubleshoot and easy to modify. Everyone structures their API differently, but ultimately consistency is key. The more consistent your API is, the easier it will be for other people to interact with it.

To download, run the following command:

~~~
go get github.com/josephspurrier/gowebapp
~~~

If you are on Go 1.5, you need to set GOVENDOREXPERIMENT to 1. If you are on Go 1.4 or earlier, the code will not work because it uses the vendor folder.

## Quick Start with MySQL

Start MySQL and import config/mysql.sql to create the database and tables.

Open config/config.json and edit the Database section so the connection information matches your MySQL instance.

Build and run from the root directory. Open your REST client to: http://localhost. You should see the welcome message and status 200.

To create a user, send a POST request to http://localhost/user with the following fields: first_name, last_name, email, and password.

## Available Endpoints

The following endpoints are available:

```
* POST   http://localhost/users		 - Create a new user
* GET	 http://localhost/users/{id} - Retrieve a user by ID
* GET	 http://localhost/users 	 - Retrieve a list of all users
* PUT	 http://localhost/users/{id} - Update a user by ID
* DELETE http://localhost/users/{id} - Delete a user by ID
* DELETE http://localhost/users		 - Delete all users
```

## Structure

The majority of the code is in the **vendor/app** folder. I made this decision originally on
my [GoWebApp](https://github.com/josephspurrier/gowebapp) project because there were
a lot of users trying to use the code on their own, but had to change all the imports path for it to work properly.
The only downside is godoc does not work with the vendor folder method. Luckily, all the code can be moved out of the vendor
folder and then a quick find and replace will get it working again if you want.

The files that are probably of most interest to you are these:

* vendor/app/controller/user.go - Controller with the routes for /users
* vendor/app/model/user/user.go - Model with all the MySQL logic
* vendor/app/shared/form/form.go - Functions that automate the interaction between the model struct and forms
* vendor/app/shared/response/response.go - Handles the output of JSON in 3 different formats

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

## Goals for this project

Integrate security similar to Parse: http://blog.parse.com/learn/secure-your-app-one-class-at-a-time/

Code generation for the following:
* Controllers with routes
* Models
* Endpoint tests
* Swagger spec

## My Other Projects

[GoWebApp](https://github.com/josephspurrier/gowebapp) demonstrates how to build a website using the Go language without a framework. Much
of the structure of this project comes from GoWebApp.

I'll use the [apigen](https://github.com/josephspurrier/apigen) project for the code generation.