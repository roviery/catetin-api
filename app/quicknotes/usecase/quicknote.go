package usecase

import (
	"context"

	"github.com/roviery/catetin-api/domain"
	"github.com/roviery/catetin-api/entity"
	"github.com/roviery/catetin-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type quicknoteUsecase struct {
	quicknoteRepo domain.QuicknoteRepository
}

func NewQuicknoteUsecase(quicknoteRepo domain.QuicknoteRepository) domain.QuicknoteUsecase {
	return &quicknoteUsecase{
		quicknoteRepo: quicknoteRepo,
	}
}

func (q *quicknoteUsecase) CreateQuicknote(ctx context.Context, ownerId *primitive.ObjectID, req *entity.QuicknoteRequest) (*entity.QuicknoteIDResponse, error) {
	res, err := q.quicknoteRepo.Store(ctx, &models.Quicknote{
		OwnerID: *ownerId,
		Note:    req.Note,
	})
	if err != nil {
		return nil, err
	}

	return &entity.QuicknoteIDResponse{ID: *res}, nil
}

func (q *quicknoteUsecase) GetQuicknote(ctx context.Context, id *primitive.ObjectID) (*entity.QuicknoteResponse, error) {
	res, err := q.quicknoteRepo.FindQuicknoteById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.QuicknoteResponse{
		ID:   res.ID,
		Note: res.Note,
	}, nil
}

func (q *quicknoteUsecase) GetQuicknotes(ctx context.Context, ownerId *primitive.ObjectID) (*entity.QuicknotesResponse, error) {
	res, err := q.quicknoteRepo.FindQuicknoteByOwnerId(ctx, ownerId)
	if err != nil {
		return nil, err
	}

	var quicknotesEntity []entity.QuicknoteResponse
	for _, quicknote := range res {
		quicknotesEntity = append(quicknotesEntity, entity.QuicknoteResponse{
			ID:   quicknote.ID,
			Note: quicknote.Note,
		})
	}

	return &entity.QuicknotesResponse{
		Quicknotes: quicknotesEntity,
	}, nil
}

func (q *quicknoteUsecase) UpdateQuicknoteByID(ctx context.Context, id *primitive.ObjectID, req *entity.QuicknoteRequest) (*entity.QuicknoteIDResponse, error) {
	_, err := q.quicknoteRepo.FindQuicknoteById(ctx, id)
	if err != nil {
		return nil, err
	}

	res, err := q.quicknoteRepo.UpdateQuicknoteByID(ctx, id, &models.Quicknote{
		Note: req.Note,
	})
	if err != nil {
		return nil, err
	}

	return &entity.QuicknoteIDResponse{
		ID: *res,
	}, nil
}

func (q *quicknoteUsecase) DeleteQuicknoteByID(ctx context.Context, id *primitive.ObjectID) error {
	_, err := q.quicknoteRepo.FindQuicknoteById(ctx, id)
	if err != nil {
		return err
	}

	err = q.quicknoteRepo.DeleteQuicknoteByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
