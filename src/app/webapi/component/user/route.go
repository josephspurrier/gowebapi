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
	router.Post("/v1/user", p.Create)
	router.Get("/v1/user/:user_id", p.Show)
	router.Get("/v1/user", p.Index)
	router.Put("/v1/user/:user_id", p.Update)
	router.Delete("/v1/user/:user_id", p.Destroy)
	router.Delete("/v1/user", p.DestroyAll)
}
