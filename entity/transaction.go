package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type TransactionRequest struct {
	FinanceID primitive.ObjectID `json:"finance_id"`
	Title     string             `json:"title"`
	Expense   int                `json:"expense"`
}

type TransactionIDResponse struct {
	ID primitive.ObjectID `json:"id"`
}

type TransactionResponse struct {
	ID        primitive.ObjectID `json:"id"`
	FinanceID primitive.ObjectID `json:"finance_id"`
	Title     string             `json:"title"`
	Expense   int                `json:"expense"`
	CreatedAt string             `json:"created_at"`
}

type TransactionsResponse struct {
	Transactions []TransactionResponse `json:"transactions"`
}
