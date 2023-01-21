package domain

import (
	"context"

	"github.com/roviery/catetin-api/entity"
	"github.com/roviery/catetin-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuicknoteRepository interface {
	Store(ctx context.Context, quicknote *models.Quicknote) (*primitive.ObjectID, error)
	FindQuicknoteById(ctx context.Context, id *primitive.ObjectID) (*models.Quicknote, error)
	FindQuicknoteByOwnerId(ctx context.Context, ownerId *primitive.ObjectID) ([]models.Quicknote, error)
	UpdateQuicknoteByID(ctx context.Context, id *primitive.ObjectID, quicknote *models.Quicknote) (*primitive.ObjectID, error)
	DeleteQuicknoteByID(ctx context.Context, id *primitive.ObjectID) error
}

type QuicknoteUsecase interface {
	CreateQuicknote(ctx context.Context, ownerId *primitive.ObjectID, req *entity.QuicknoteRequest) (*entity.QuicknoteIDResponse, error)
	GetQuicknote(ctx context.Context, id *primitive.ObjectID) (*entity.QuicknoteResponse, error)
	GetQuicknotes(ctx context.Context, ownerId *primitive.ObjectID) (*entity.QuicknotesResponse, error)
	UpdateQuicknoteByID(ctx context.Context, id *primitive.ObjectID, req *entity.QuicknoteRequest) (*entity.QuicknoteIDResponse, error)
	DeleteQuicknoteByID(ctx context.Context, id *primitive.ObjectID) error
}
