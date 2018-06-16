package user

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
	router.Post("/v1/user", component.H(p.Create))
	router.Get("/v1/user/:user_id", component.H(p.Show))
	router.Get("/v1/user", component.H(p.Index))
	router.Put("/v1/user/:user_id", component.H(p.Update))
	router.Delete("/v1/user/:user_id", component.H(p.Destroy))
	router.Delete("/v1/user", component.H(p.DestroyAll))
}
