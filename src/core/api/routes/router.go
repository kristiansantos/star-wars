package routes

import (
	"github.com/go-chi/chi"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/api/handlers"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/api/routes/planet"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/middlewares"
)

type router struct {
	Client   *chi.Mux
	Handlers handlers.IHandler
}

func NewRoutes(handlers handlers.IHandler) *router {
	return &router{
		Client:   chi.NewRouter(),
		Handlers: handlers,
	}
}

func (r *router) Setup() {
	middlewares.Default(r.Client)

	planet.NewRoutes(r.Client, r.Handlers.NewPlanetHandler())
}
