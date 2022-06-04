package index

import (
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/domains/planet/repositories"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/providers/logger"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/tools/communication"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/tools/namespace"
	"go.mongodb.org/mongo-driver/bson"
)

var Namespace = namespace.New("core.domains.push.services.list.list_service")

type Service struct {
	Repository repositories.IPlanetRepository
	Logger     logger.ILoggerProvider
}

func (service *Service) Execute() communication.Response {
	service.Logger.Info(Namespace.Concat("Execute"), "")

	comm := communication.New()

	// Get documents in database
	filter := bson.M{}
	documents, err := service.Repository.GetPlanets(filter)
	if err != nil {
		service.Logger.Error(Namespace.Concat("Execute", "FindAll"), err.Error())
		return comm.ResponseError(400, "error_list", err)
	}

	return comm.Response(200, "success", documents)
}
