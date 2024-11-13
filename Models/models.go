package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	Movie   string             `json:"movie,omitempty"`
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"` // bson is binary json used in mongo db
	Watched bool               `json:"watched,omitempty"`
}
