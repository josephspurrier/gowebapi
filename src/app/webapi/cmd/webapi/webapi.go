package main

import (
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

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

// Clock is a clock.
type Clock struct{}

// Now returns the current time.
func (c *Clock) Now() time.Time {
	return time.Now()
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

	// Set up the routes.
	mux := webapi.Routes(config, appLogger)

	// Set up the HTTP listener.
	httpServer := new(http.Server)
	httpServer.Addr = config.Server.HTTPAddress()

	// Determine if HTTP should redirect to HTTPS.
	if config.Server.ForceHTTPSRedirect {
		httpServer.Handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			http.Redirect(w, req, "https://"+req.Host, http.StatusMovedPermanently)
		})
	} else {
		httpServer.Handler = middleware.Wrap(mux, appLogger, config.JWT.Secret)
	}

	// Set up the HTTPS listener.
	httpsServer := new(http.Server)
	httpsServer.Addr = config.Server.HTTPSAddress()
	httpsServer.Handler = middleware.Wrap(mux, appLogger, config.JWT.Secret)

	// Start the listeners based on the config.
	config.Server.Run(httpServer, httpsServer, appLogger)
}
