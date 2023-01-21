package domain

import (
	"context"

	"github.com/roviery/catetin-api/entity"
	"github.com/roviery/catetin-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionRepository interface {
	Store(ctx context.Context, transaction *models.Transaction) (*primitive.ObjectID, error)
	FindTransactionByID(ctx context.Context, id *primitive.ObjectID) (*models.Transaction, error)
	FindTransactionByOwnerID(ctx context.Context, ownerId *primitive.ObjectID, filter *string) ([]models.Transaction, error)
	UpdateTransactionByID(ctx context.Context, id *primitive.ObjectID, transaction *models.Transaction) (*primitive.ObjectID, error)
	DeleteTransactionByID(ctx context.Context, id *primitive.ObjectID) error
}

type TransactionUsecase interface {
	CreateTransaction(ctx context.Context, ownerId *primitive.ObjectID, req *entity.TransactionRequest) (*entity.TransactionIDResponse, error)
	GetTransaction(ctx context.Context, id *primitive.ObjectID) (*entity.TransactionResponse, error)
	GetTransactions(ctx context.Context, ownerId *primitive.ObjectID, filter *string) (*entity.TransactionsResponse, error)
	UpdateTransactionByID(ctx context.Context, id *primitive.ObjectID, req *entity.TransactionRequest) (*entity.TransactionIDResponse, error)
	DeleteTransactionByID(ctx context.Context, id *primitive.ObjectID) error
}
