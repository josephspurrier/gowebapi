package main

import (
	"log"
	"os"
	"runtime"

	"app/webapi"
	"app/webapi/pkg/jsonconfig"
	"app/webapi/pkg/logger"
)

func init() {
	// Verbose logging with file name and line number.
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores.
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// Create the logger.
	l := logger.New(log.New(os.Stderr, "", log.LstdFlags))

	// Get the config file name from the first argument.
	configFile := "config.json"
	if len(os.Args) > 1 {
		configFile = os.Args[1]
	}

	// Load the configuration file.
	config := new(webapi.AppConfig)
	err := jsonconfig.Load(configFile, config)
	if err != nil {
		l.Fatalf("%v", err)
	}

	// Set up the service, routes, and the handlers.
	core := webapi.Services(config, l)
	mux := webapi.Routes(core)
	httpServer, httpsServer := webapi.Handlers(config, l, mux)

	// Start the listeners based on the config.
	config.Server.Run(httpServer, httpsServer, l)
}
