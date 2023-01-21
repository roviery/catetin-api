package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type TodoRequest struct {
	Todo   string `json:"todo"`
	Status string `json:"status"`
}

type TodoIDResponse struct {
	ID primitive.ObjectID `json:"id"`
}

type TodoResponse struct {
	ID     primitive.ObjectID `json:"id"`
	Todo   string             `json:"todo"`
	Status string             `json:"Status"`
}

type TodosResponse struct {
	Todos []TodoResponse
}
