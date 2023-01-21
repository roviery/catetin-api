package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type QuicknoteRequest struct {
	Note string `json:"note"`
}

type QuicknoteIDResponse struct {
	ID primitive.ObjectID `json:"id"`
}

type QuicknoteResponse struct {
	ID   primitive.ObjectID `json:"id"`
	Note string             `json:"note"`
}

type QuicknotesResponse struct {
	Quicknotes []QuicknoteResponse `json:"quicknotes"`
}
