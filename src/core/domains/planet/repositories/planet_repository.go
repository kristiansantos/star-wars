package repositories

import (
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/domains/planet/entities"
	"go.mongodb.org/mongo-driver/bson"
)

type IPlanetRepository interface {
	GetPlanets(filter bson.M) (entities.Planets, error)
	GetById(id string) (entities.Planet, error)
	Create(planet entities.Planet) (entities.Planet, error)
	Update(id string, planet entities.Planet) (entities.Planet, error)
	Delete(id string) error
}
