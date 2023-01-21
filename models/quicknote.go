package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Quicknote struct {
	ID        primitive.ObjectID `bson:"_id"`
	OwnerID   primitive.ObjectID `bson:"owner_id"`
	Note      string             `bson:"note"`
	CreatedAt string             `bson:"created_at"`
	UpdatedAt string             `bson:"updated_at"`
}
