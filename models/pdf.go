package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PDF struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Data      []byte             `bson:"data"`
	CreatedAt string             `bson:"created_at"`
	UpdatedAt string             `bson:"updated_at"`
}
