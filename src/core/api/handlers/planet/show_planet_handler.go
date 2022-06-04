package planet

import (
	"net/http"

	httphelper "gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/providers/http_helper"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/tools/communication"
)

func (handler handler) ShowPlanetHandler(r *http.Request) communication.Response {
	Namespace.AddComponent("show_planet_handler")

	handler.Logger.Info(Namespace.Concat("ShowPlanetHandler"), "")

	id := httphelper.GetParam(r, "planetID")
	ctx := r.Context()
	service := handler.Service(ctx).Show

	return service.Execute(id)
}
