package entity

import (
	"github.com/golang-jwt/jwt/v4"
)

type JWTCustomClaims struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	jwt.RegisteredClaims
}
