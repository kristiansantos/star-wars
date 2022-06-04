package planet

import (
	"net/http"

	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/tools/communication"
)

func (handler handler) IndexPlanetHandler(r *http.Request) communication.Response {
	Namespace.AddComponent("index_planet_handler")

	handler.Logger.Info(Namespace.Concat("IndexPlanetHandler"), "")

	ctx := r.Context()
	service := handler.Service(ctx).Index

	return service.Execute()
}
