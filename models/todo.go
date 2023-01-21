package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID        primitive.ObjectID `bson:"_id"`
	OwnerID   primitive.ObjectID `bson:"owner_id"`
	Todo      string             `bson:"todo"`
	Status    string             `bson:"status"`
	CreatedAt string             `bson:"created_at"`
	UpdatedAt string             `bson:"updated_at"`
}
