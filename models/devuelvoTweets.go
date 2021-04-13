package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DevuelvoTweets struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID  primitive.ObjectID `bson:"userid" json:"userId,omitempty"`
	Mensaje primitive.ObjectID `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   primitive.ObjectID `bson:"fecha" json:"fecha,omitempty"`
}
