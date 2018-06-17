package main

import (
	"log"
	"net/http"
	"os"
	"runtime"

	"app/webapi"
	"app/webapi/middleware"
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
		log.Fatal(err)
	}

	// Boot the application.
	mux := webapi.Boot(config, appLogger)

	// Get the router instance.
	r := mux.Instance()

	// Start the web listener(s).
	httpServer := new(http.Server)
	httpServer.Addr = config.Server.HTTPAddress()
	httpServer.Handler = middleware.LoadHTTP(r)
	httpsServer := new(http.Server)
	httpsServer.Addr = config.Server.HTTPSAddress()
	httpsServer.Handler = middleware.LoadHTTPS(r)
	config.Server.Run(httpServer, httpsServer, appLogger)
}
