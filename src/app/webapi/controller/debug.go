package controller

import (
	"app/webapi/middleware/pprofhandler"
	"app/webapi/pkg/router"
)

func init() {
	// Enable Pprof.
	router.Get("/debug/pprof/*pprof", pprofhandler.Handler)
}
