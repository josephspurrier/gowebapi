package controller

import (
	"app/route/middleware/pprofhandler"
	"app/shared/router"
)

func init() {
	// Enable Pprof
	router.Instance().GET("/debug/pprof/*pprof", router.HandlerFunc(pprofhandler.Handler))
}
