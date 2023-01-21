package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Fullname  string             `bson:"fullname"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	CreatedAt string             `bson:"created_at"`
}
