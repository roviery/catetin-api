package usecase

import (
	"context"

	"github.com/roviery/catetin-api/domain"
	"github.com/roviery/catetin-api/entity"
	"github.com/roviery/catetin-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type deadlineUsecase struct {
	deadlineRepo domain.DeadlineRepository
}

func NewDeadlineUsecase(deadlineRepo domain.DeadlineRepository) domain.DeadlineUsecase {
	return &deadlineUsecase{
		deadlineRepo: deadlineRepo,
	}
}

func (d *deadlineUsecase) CreateDeadline(ctx context.Context, ownerId *primitive.ObjectID, req *entity.DeadlineRequest) (*entity.DeadlineIDResponse, error) {
	res, err := d.deadlineRepo.Store(ctx, &models.Deadline{
		OwnerID:      *ownerId,
		DeadlineDate: req.DeadlineDate,
		Task:         req.Task,
		Priority:     req.Priority,
	})
	if err != nil {
		return nil, err
	}

	return &entity.DeadlineIDResponse{
		ID: *res,
	}, nil
}

func (d *deadlineUsecase) GetDeadline(ctx context.Context, id *primitive.ObjectID) (*entity.DeadlineResponse, error) {
	res, err := d.deadlineRepo.FindDeadlineByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.DeadlineResponse{
		ID:           res.ID,
		DeadlineDate: res.DeadlineDate,
		Task:         res.Task,
		Priority:     res.Priority,
	}, nil
}

func (d *deadlineUsecase) GetDeadlines(ctx context.Context, ownerId *primitive.ObjectID) (*entity.DeadlinesResponse, error) {
	res, err := d.deadlineRepo.FindDeadlinesByOwnerID(ctx, ownerId)
	if err != nil {
		return nil, err
	}

	var deadlinesEntity []entity.DeadlineResponse
	for _, deadline := range res {
		deadlinesEntity = append(deadlinesEntity, entity.DeadlineResponse{
			ID:           deadline.ID,
			DeadlineDate: deadline.DeadlineDate,
			Task:         deadline.Task,
			Priority:     deadline.Priority,
		})
	}

	return &entity.DeadlinesResponse{
		Deadlines: deadlinesEntity,
	}, nil
}

func (d *deadlineUsecase) UpdateDeadlineByID(ctx context.Context, id *primitive.ObjectID, req *entity.DeadlineRequest) (*entity.DeadlineIDResponse, error) {
	res, err := d.deadlineRepo.UpdateDeadlineByID(ctx, id, &models.Deadline{
		DeadlineDate: req.DeadlineDate,
		Task:         req.Task,
		Priority:     req.Priority,
	})
	if err != nil {
		return nil, err
	}

	return &entity.DeadlineIDResponse{
		ID: *res,
	}, nil
}

func (d *deadlineUsecase) DeleteDeadlineByID(ctx context.Context, id *primitive.ObjectID) error {
	_, err := d.deadlineRepo.FindDeadlineByID(ctx, id)
	if err != nil {
		return err
	}

	err = d.deadlineRepo.DeleteDeadlineByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
