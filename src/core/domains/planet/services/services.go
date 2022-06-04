package services

import (
	"context"

	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/domains/planet/repositories"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/domains/planet/services/planets/create"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/domains/planet/services/planets/delete"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/domains/planet/services/planets/index"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/domains/planet/services/planets/show"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/domains/planet/services/planets/update"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/providers/hash"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/providers/logger"
)

type Dependencies struct {
	Context    context.Context
	Repository repositories.IPlanetRepository
	Logger     logger.ILoggerProvider
	Hash       hash.IHashProvider
}

type Services struct {
	Index  index.Service
	Show   show.Service
	Create create.Service
	Update update.Service
	Delete delete.Service
}

func NewPlanet(dep Dependencies) *Services {
	return &Services{
		Index: index.Service{
			Repository: dep.Repository,
			Logger:     dep.Logger,
		},
		Show: show.Service{
			Repository: dep.Repository,
			Logger:     dep.Logger,
		},
		Create: create.Service{
			Repository: dep.Repository,
			Logger:     dep.Logger,
		},
		Update: update.Service{
			Repository: dep.Repository,
			Logger:     dep.Logger,
		},
		Delete: delete.Service{
			Repository: dep.Repository,
			Logger:     dep.Logger,
		},
	}
}
