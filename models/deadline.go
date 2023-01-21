package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Deadline struct {
	ID           primitive.ObjectID `bson:"_id"`
	OwnerID      primitive.ObjectID `bson:"owner_id"`
	DeadlineDate string             `bson:"deadline_date"`
	Task         string             `bson:"task"`
	Priority     string             `bson:"priority"`
	CreatedAt    string             `bson:"created_at"`
	UpdatedAt    string             `bson:"updated_at"`
}
