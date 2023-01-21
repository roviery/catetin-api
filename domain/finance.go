package domain

import (
	"context"

	"github.com/roviery/catetin-api/entity"
	"github.com/roviery/catetin-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FinanceRepository interface {
	Store(ctx context.Context, finance *models.Finance) (*primitive.ObjectID, error)
	FindFinanceByID(ctx context.Context, id *primitive.ObjectID) (*models.Finance, error)
	FindFinanceByType(ctx context.Context, financeType *string) (*models.Finance, error)
	FindFinanceByOwnerID(ctx context.Context, ownerId *primitive.ObjectID) ([]models.Finance, error)
	UpdateFinanceByID(ctx context.Context, id *primitive.ObjectID, finance *models.Finance) (*primitive.ObjectID, error)
	DeleteFinanceByID(ctx context.Context, id *primitive.ObjectID) error
}

type FinanceUsecase interface {
	CreateFinance(ctx context.Context, ownerId *primitive.ObjectID, req *entity.FinanceRequest) (*entity.FinanceIDResponse, error)
	GetFinance(ctx context.Context, id *primitive.ObjectID) (*entity.FinanceResponse, error)
	GetFinances(ctx context.Context, ownerId *primitive.ObjectID) (*entity.FinancesResponse, error)
	UpdateFinanceByID(ctx context.Context, id *primitive.ObjectID, req *entity.FinanceRequest) (*entity.FinanceIDResponse, error)
	DeleteFinanceByID(ctx context.Context, id *primitive.ObjectID) error
}
