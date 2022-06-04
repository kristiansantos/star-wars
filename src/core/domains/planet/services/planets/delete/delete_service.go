package delete

import (
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/domains/planet/repositories"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/providers/logger"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/tools/communication"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/tools/namespace"
)

var Namespace = namespace.New("core.domains.push.services.delete.delete_service")

type Service struct {
	Repository repositories.IPlanetRepository
	Logger     logger.ILoggerProvider
}

func (service *Service) Execute(id string) communication.Response {
	service.Logger.Info(Namespace.Concat("Execute"), "")

	comm := communication.New()
	// Get documents in database
	err := service.Repository.Delete(id)

	if err != nil {
		service.Logger.Error(Namespace.Concat("Execute", "Delete"), err.Error())
		return comm.ResponseError(400, "error_delete", err)
	}

	return comm.Response(202, "success", "")
}
