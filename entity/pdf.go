package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type PDFRequest struct {
	Name string `json:"name"`
	Data []byte `json:"data"`
}

type PDFIDResponse struct {
	ID primitive.ObjectID `json:"id"`
}

type PDFResponse struct {
	ID   primitive.ObjectID `json:"id"`
	Name string             `json:"name"`
	Data []byte             `json:"data"`
}
