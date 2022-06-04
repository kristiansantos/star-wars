package update

import (
	"time"

	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/domains/planet/entities"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/domains/planet/repositories"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/providers/logger"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/tools/communication"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/tools/namespace"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/tools/parse"
)

type Dto struct {
	Name    string `json:"name"`
	Climate string `json:"climate"`
	Ground  string `json:"ground"`
}

var Namespace = namespace.New("core.domains.push.services.update.update_service")

type Service struct {
	Repository repositories.IPlanetRepository
	Logger     logger.ILoggerProvider
}

func (service *Service) Execute(id string, dto Dto) communication.Response {
	service.Logger.Info(Namespace.Concat("Execute"), "")

	var planet entities.Planet

	comm := communication.New()
	planet.ID = id
	planet.UpdatedAt = time.Now()

	parse.Unmarshal(dto, &planet)

	// Update document in database
	document, err := service.Repository.Update(id, planet)
	if err != nil {
		service.Logger.Error(Namespace.Concat("Execute", "Update"), err.Error())
		return comm.ResponseError(400, "error_update", err)
	}

	return comm.Response(200, "success", document)
}
