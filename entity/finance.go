package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type FinanceRequest struct {
	Type           string `json:"type"`
	FundAllocation int    `json:"fund_allocation"`
}

type FinanceIDResponse struct {
	ID primitive.ObjectID `json:"id"`
}

type FinanceResponse struct {
	ID             primitive.ObjectID `json:"id"`
	Type           string             `json:"type"`
	FundAllocation int                `json:"fund_allocation"`
	Used           int                `json:"used"`
	Remaining      int                `json:"remaining"`
}

type FinancesResponse struct {
	Finances []FinanceResponse `json:"finances"`
}
