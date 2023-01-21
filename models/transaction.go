package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {
	ID        primitive.ObjectID `bson:"_id"`
	FinanceID primitive.ObjectID `bson:"finance_id"`
	OwnerID   primitive.ObjectID `bson:"owner_id"`
	Title     string             `bson:"title"`
	Expense   int                `bson:"expense"`
	CreatedAt string             `bson:"created_at"`
	UpdatedAt string             `bson:"updated_at"`
}
