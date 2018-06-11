package controller

import (
	"app/webapi/route/middleware/pprofhandler"
	"app/webapi/shared/router"
)

func init() {
	// Enable Pprof.
	router.Get("/debug/pprof/*pprof", pprofhandler.Handler)
}
