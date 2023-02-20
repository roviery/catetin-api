package usecase

import (
	"context"
	"fmt"

	"github.com/roviery/catetin-api/domain"
	"github.com/roviery/catetin-api/entity"
	"github.com/roviery/catetin-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type financeUsecase struct {
	financeRepo domain.FinanceRepository
}

func NewFinanceUsecase(financeRepo domain.FinanceRepository) domain.FinanceUsecase {
	return &financeUsecase{
		financeRepo: financeRepo,
	}
}

func (f *financeUsecase) CreateFinance(ctx context.Context, ownerId *primitive.ObjectID, req *entity.FinanceRequest) (*entity.FinanceIDResponse, error) {
	check, _ := f.financeRepo.FindFinanceByType(ctx, &req.Type)
	if check != nil {
		return nil, fmt.Errorf("finance type already exist")
	}

	res, err := f.financeRepo.Store(ctx, &models.Finance{
		OwnerID:        *ownerId,
		Type:           req.Type,
		FundAllocation: req.FundAllocation,
		Remaining:      req.FundAllocation,
	})
	if err != nil {
		return nil, err
	}

	return &entity.FinanceIDResponse{ID: *res}, nil
}

func (f *financeUsecase) GetFinance(ctx context.Context, id *primitive.ObjectID) (*entity.FinanceResponse, error) {
	res, err := f.financeRepo.FindFinanceByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.FinanceResponse{
		ID:             res.ID,
		Type:           res.Type,
		FundAllocation: res.FundAllocation,
		Used:           res.Used,
		Remaining:      res.Remaining,
	}, nil
}

func (f *financeUsecase) GetFinances(ctx context.Context, ownerId *primitive.ObjectID) (*entity.FinancesResponse, error) {
	res, err := f.financeRepo.FindFinanceByOwnerID(ctx, ownerId)
	if err != nil {
		return nil, err
	}

	var financeEntity []entity.FinanceResponse
	for _, finance := range res {
		financeEntity = append(financeEntity, entity.FinanceResponse{
			ID:             finance.ID,
			Type:           finance.Type,
			FundAllocation: finance.FundAllocation,
			Used:           finance.Used,
			Remaining:      finance.Remaining,
		})
	}

	return &entity.FinancesResponse{
		Finances: financeEntity,
	}, nil
}

func (f *financeUsecase) UpdateFinanceByID(ctx context.Context, id *primitive.ObjectID, req *entity.FinanceRequest) (*entity.FinanceIDResponse, error) {
	oldFinance, err := f.financeRepo.FindFinanceByID(ctx, id)
	if err != nil {
		return nil, err
	}

	res, err := f.financeRepo.UpdateFinanceByID(ctx, id, &models.Finance{
		Type:           req.Type,
		FundAllocation: req.FundAllocation,
		Used:           oldFinance.Used,
		Remaining:      req.FundAllocation - oldFinance.Used,
	})
	if err != nil {
		return nil, err
	}

	return &entity.FinanceIDResponse{ID: *res}, nil
}

func (f *financeUsecase) DeleteFinanceByID(ctx context.Context, id *primitive.ObjectID) error {
	_, err := f.financeRepo.FindFinanceByID(ctx, id)
	if err != nil {
		return err
	}

	err = f.financeRepo.DeleteFinanceByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
