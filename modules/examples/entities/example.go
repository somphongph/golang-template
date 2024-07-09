package entities

import (
	"golang-template/modules/common/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Example struct {
	Id     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title  string             `json:"title"`
	Detail string             `json:"detail"`

	entities.Entity `bson:"inline"`
}
