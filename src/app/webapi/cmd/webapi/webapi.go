package main

import (
	"log"
	"os"
	"runtime"

	"app/webapi"
	"app/webapi/pkg/jsonconfig"
)

func init() {
	// Verbose logging with file name and line number.
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores.
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// Create the logger.
	appLogger := log.New(os.Stderr, "", log.LstdFlags)

	// Load the configuration file.
	config := new(webapi.AppConfig)
	err := jsonconfig.Load("config.json", config)
	if err != nil {
		appLogger.Fatalf("%v", err)
	}

	// Set up the routes.
	_, httpServer, httpsServer := webapi.Routes(config, appLogger)

	// Start the listeners based on the config.
	config.Server.Run(httpServer, httpsServer, appLogger)
}
