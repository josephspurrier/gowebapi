package main

import (
	"log"
	"runtime"

	"app/webapi"
)

func init() {
	// Verbose logging with file name and line number.
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores.
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	webapi.Boot()
}
