package planet

import (
	"net/http"

	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/domains/planet/entities"
	httphelper "gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/providers/http_helper"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/tools/communication"
)

func (handler handler) CreatePlanetHandler(r *http.Request) communication.Response {
	Namespace.AddComponent("create_planet_handler")

	handler.Logger.Info(Namespace.Concat("CreatePlanetHandler"), "")

	var dto entities.PlanetCreateDto

	ctx := r.Context()
	service := handler.Service(ctx).Create
	comm := communication.New()

	if err := httphelper.GetBody(r.Body, &dto); err != nil {
		handler.Logger.Error(Namespace.Concat("CreatePlanetHandler", "GetBody"), err.Error())

		return comm.ResponseError(400, "error_create", err)
	}

	return service.Execute(dto)

}
