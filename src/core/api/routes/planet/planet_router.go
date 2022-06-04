package planet

import (
	"github.com/go-chi/chi"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/api/handlers/planet"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/middlewares"
)

func NewRoutes(router *chi.Mux, planetHandler planet.IHandler) {
	router.Route("/api/v1/planets", func(r chi.Router) {
		r.Get("/", middlewares.Handler(planetHandler.IndexPlanetHandler))
		r.Get("/{planetID}", middlewares.Handler(planetHandler.ShowPlanetHandler))
		r.Post("/", middlewares.Handler(planetHandler.CreatePlanetHandler))
		r.Put("/{planetID}", middlewares.Handler(planetHandler.UpdatePlanetHandler))
		r.Delete("/{planetID}", middlewares.Handler(planetHandler.DeletePlanetHandler))
	})
}
