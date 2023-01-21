package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Finance struct {
	ID             primitive.ObjectID `bson:"_id"`
	OwnerID        primitive.ObjectID `bson:"owner_id"`
	Type           string             `bson:"type"`
	FundAllocation int                `bson:"fund_allocation"`
	Used           int                `bson:"used"`
	Remaining      int                `bson:"remaining"`
	CreatedAt      string             `bson:"created_at"`
	UpdatedAt      string             `bson:"updated_at"`
}
