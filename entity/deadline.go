package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type DeadlineRequest struct {
	DeadlineDate string `json:"deadline_date"`
	Task         string `json:"task"`
	Priority     string `json:"priority"`
}

type DeadlineIDResponse struct {
	ID primitive.ObjectID `json:"id"`
}

type DeadlineResponse struct {
	ID           primitive.ObjectID `json:"id"`
	DeadlineDate string             `json:"deadline_date"`
	Task         string             `json:"task"`
	Priority     string             `json:"priority"`
}

type DeadlinesResponse struct {
	Deadlines []DeadlineResponse `json:"deadlines"`
}
