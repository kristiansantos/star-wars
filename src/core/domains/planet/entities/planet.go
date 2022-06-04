package entities

import (
	"time"

	"github.com/google/uuid"
)

type Planet struct {
	ID             string    `json:"id" bson:"_id"`
	Name           string    `json:"name" bson:"name"`
	Climate        string    `json:"climate" bson:"climate"`
	Ground         string    `json:"ground" bson:"ground"`
	FilmApparences string    `json:"film_apparences" bson:"film_apparences"`
	CreatedAt      time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt" bson:"updatedAt"`
}

type Planets []*Planet

func NewPlanet() *Planet {
	return &Planet{}
}

func (planet *Planet) Populate() {
	planet.ID = uuid.New().String()
	planet.CreatedAt = time.Now()
	planet.UpdatedAt = time.Now()
}
