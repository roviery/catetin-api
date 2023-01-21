package domain

import (
	"context"

	"github.com/roviery/catetin-api/entity"
	"github.com/roviery/catetin-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeadlineRepository interface {
	Store(ctx context.Context, deadline *models.Deadline) (*primitive.ObjectID, error)
	FindDeadlineByID(ctx context.Context, id *primitive.ObjectID) (*models.Deadline, error)
	FindDeadlinesByOwnerID(ctx context.Context, ownerId *primitive.ObjectID) ([]models.Deadline, error)
	UpdateDeadlineByID(ctx context.Context, id *primitive.ObjectID, deadline *models.Deadline) (*primitive.ObjectID, error)
	DeleteDeadlineByID(ctx context.Context, id *primitive.ObjectID) error
}

type DeadlineUsecase interface {
	CreateDeadline(ctx context.Context, ownerId *primitive.ObjectID, req *entity.DeadlineRequest) (*entity.DeadlineIDResponse, error)
	GetDeadline(ctx context.Context, id *primitive.ObjectID) (*entity.DeadlineResponse, error)
	GetDeadlines(ctx context.Context, ownerId *primitive.ObjectID) (*entity.DeadlinesResponse, error)
	UpdateDeadlineByID(ctx context.Context, id *primitive.ObjectID, req *entity.DeadlineRequest) (*entity.DeadlineIDResponse, error)
	DeleteDeadlineByID(ctx context.Context, id *primitive.ObjectID) error
}
