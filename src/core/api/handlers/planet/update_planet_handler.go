package planet

import (
	"net/http"

	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/domains/planet/services/planets/update"
	httphelper "gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/providers/http_helper"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/tools/communication"
)

func (handler handler) UpdatePlanetHandler(r *http.Request) communication.Response {
	Namespace.AddComponent("create_planet_handler")

	handler.Logger.Info(Namespace.Concat("CreatePlanetHandler"), "")

	var dto update.Dto

	ctx := r.Context()
	service := handler.Service(ctx).Update
	comm := communication.New()
	id := httphelper.GetParam(r, "planetID")

	if err := httphelper.GetBody(r.Body, &dto); err != nil {
		handler.Logger.Error(Namespace.Concat("CreatePlanetHandler", "GetBody"), err.Error())

		return comm.ResponseError(400, "error_create", err)
	}

	return service.Execute(id, dto)
}
