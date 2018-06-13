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
	router.Post("/v1/user", component.F(p.Create))
	router.Get("/v1/user/:user_id", component.F(p.Show))
	router.Get("/v1/user", component.F(p.Index))
	router.Put("/v1/user/:user_id", component.F(p.Update))
	router.Delete("/v1/user/:user_id", component.F(p.Destroy))
	router.Delete("/v1/user", component.F(p.DestroyAll))
}
