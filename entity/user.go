package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	UserID primitive.ObjectID `json:"user_id"`
}
