package auth

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
	router.Get("/v1/auth", p.Index)
}
