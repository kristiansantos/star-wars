package planet

import (
	"context"

	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/domains/planet/entities"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/database/mongodb"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/providers/logger"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/tools/namespace"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collection = "planets"
	Namespace  = namespace.New("infra.repositories.mongodb.planet.planet_repository")
)

type Repository struct {
	Context    context.Context
	Collection *mongo.Collection
	Logger     logger.ILoggerProvider
}

func Setup(ctx context.Context) *Repository {
	connection := mongodb.New(ctx)
	log := logger.Instance

	return &Repository{
		Context:    ctx,
		Collection: connection.MongoDB.Collection(collection),
		Logger:     log,
	}
}

func (r Repository) GetPlanets(filter bson.M) (planets entities.Planets, err error) {
	r.Logger.Info(Namespace.Concat("GetPlanets"), "")

	cursor, err := r.Collection.Find(r.Context, filter)
	if err != nil {
		return
	}

	for cursor.Next(r.Context) {
		document := &entities.Planet{}
		cursor.Decode(&document)
		planets = append(planets, document)
	}

	return
}

func (r Repository) GetById(id string) (planet entities.Planet, err error) {
	r.Logger.Info(Namespace.Concat("GetById"), "")

	filter := bson.M{"_id": id}

	FindError := r.Collection.FindOne(r.Context, filter).Decode(&planet)
	if FindError != nil {
		return entities.Planet{}, FindError
	}

	return
}

func (r Repository) Create(document entities.Planet) (entities.Planet, error) {
	r.Logger.Info(Namespace.Concat("Create"), "")

	_, InsertOneError := r.Collection.InsertOne(r.Context, document)
	if InsertOneError != nil {
		return entities.Planet{}, InsertOneError
	}

	return r.GetById(document.ID)
}

func (r Repository) Update(id string, document entities.Planet) (entities.Planet, error) {
	r.Logger.Info(Namespace.Concat("Update"), "")

	filter := bson.M{"_id": id}
	update := bson.M{"$set": document}

	_, UpdateOneError := r.Collection.UpdateOne(r.Context, filter, update)
	if UpdateOneError != nil {
		return entities.Planet{}, UpdateOneError
	}

	return r.GetById(id)
}

func (r Repository) Delete(id string) (err error) {
	r.Logger.Info(Namespace.Concat("Update"), "")

	filter := bson.M{"_id": id}

	_, deleteError := r.Collection.DeleteOne(r.Context, filter)
	if deleteError != nil {
		return deleteError
	}

	return
}
