package main

import (
	"encoding/json"
	"log"
	"runtime"

	"app/webapi/controller"
	"app/webapi/pkg/database"
	"app/webapi/pkg/email"
	"app/webapi/pkg/jsonconfig"
	"app/webapi/pkg/server"
	"app/webapi/route"
)

// *****************************************************************************
// Application Logic
// *****************************************************************************

func init() {
	// Verbose logging with file name and line number.
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores.
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// Load the configuration file.
	jsonconfig.Load("config.json", config)

	// Configure the email settings.
	email.Configure(config.Email)

	// Connect to database.
	database.Connect(config.Database)

	// Load the controller routes.
	controller.Load()

	// Start the listener.
	server.Run(route.LoadHTTP(), route.LoadHTTPS(), config.Server)
}

// *****************************************************************************
// Application Settings
// *****************************************************************************

// config the settings variable.
var config = &configuration{}

// configuration contains the application settings.
type configuration struct {
	Database database.Info  `json:"Database"`
	Email    email.SMTPInfo `json:"Email"`
	Server   server.Server  `json:"Server"`
}

// ParseJSON unmarshals bytes to structs.
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}
