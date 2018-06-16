package static

import (
	"app/webapi/component"
)

// New returns a new instance of the endpoint.
func New(bc component.Core) *Endpoint {
	return &Endpoint{
		Core: bc,
	}
}

// Endpoint contains the dependencies.
type Endpoint struct {
	component.Core
}

// Routes will set up the endpoints.
func (p *Endpoint) Routes(router component.IRouter) {
	// Serve just the favorite icon and don't allow directory browsing.
	router.Get("/favicon.ico", component.H(p.Static))
}
