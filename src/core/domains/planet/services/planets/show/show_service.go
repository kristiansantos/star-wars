package show

import (
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/domains/planet/repositories"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/providers/logger"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/tools/communication"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/tools/namespace"
	"go.mongodb.org/mongo-driver/mongo"
)

var Namespace = namespace.New("core.domains.push.services.show.show_service")

type Service struct {
	Repository repositories.IPlanetRepository
	Logger     logger.ILoggerProvider
}

func (service *Service) Execute(id string) communication.Response {
	service.Logger.Info(Namespace.Concat("Execute"), "")

	comm := communication.New()

	// Get documents in database
	switch document, err := service.Repository.GetById(id); err {
	case nil:
		return comm.Response(200, "success", document)
	case mongo.ErrNoDocuments:
		service.Logger.Error(Namespace.Concat("Execute", "GetById"), err.Error())
		return comm.ResponseError(404, "not_found", err)
	default:
		service.Logger.Error(Namespace.Concat("Execute", "GetById"), err.Error())
		return comm.ResponseError(400, "error_list", err)
	}
}
