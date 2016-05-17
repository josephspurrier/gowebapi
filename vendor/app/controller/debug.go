package controller

import (
	"app/route/middleware/pprofhandler"
	"app/shared/router"
)

func init() {
	// Enable Pprof
	router.Get("/debug/pprof/*pprof", pprofhandler.Handler)
}
