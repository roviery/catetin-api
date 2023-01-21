package entity

import (
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTCustomClaims struct {
	Name string             `json:"name"`
	ID   primitive.ObjectID `json:"id"`
	jwt.RegisteredClaims
}
