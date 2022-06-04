package planet

import (
	"context"
	"net/http"

	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/domains/planet/services"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/infra/repositories/mongodb/planet"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/providers/hash"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/providers/logger"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/tools/communication"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/tools/namespace"
)

var Namespace = namespace.New("core.api.handlers.planet")

type IHandler interface {
	IndexPlanetHandler(r *http.Request) communication.Response
	ShowPlanetHandler(r *http.Request) communication.Response
	CreatePlanetHandler(r *http.Request) communication.Response
	UpdatePlanetHandler(r *http.Request) communication.Response
	DeletePlanetHandler(r *http.Request) communication.Response
}

type handler struct {
	Logger logger.ILoggerProvider
}

func NewHandler(logger logger.ILoggerProvider) handler {
	return handler{
		Logger: logger,
	}
}

func (a handler) Service(ctx context.Context) *services.Services {
	planetRepository := planet.Setup(ctx)

	dependencies := services.Dependencies{
		Context:    ctx,
		Repository: planetRepository,
		Logger:     logger.New(),
		Hash:       hash.New(),
	}
	return services.NewPlanet(dependencies)
}
