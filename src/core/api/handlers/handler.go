package handlers

import (
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/api/handlers/planet"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/providers/logger"
)

type Dependencies struct {
	Logger logger.ILoggerProvider
}

type IHandler interface {
	NewPlanetHandler() planet.IHandler
}

type handler struct {
	PlanetHandler planet.IHandler
}

func NewHandler(dep Dependencies) handler {
	return handler{
		PlanetHandler: planet.NewHandler(dep.Logger),
	}
}

func (h handler) NewPlanetHandler() planet.IHandler {
	return h.PlanetHandler
}
