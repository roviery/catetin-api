package usecase

import (
	"context"

	"github.com/roviery/catetin-api/domain"
	"github.com/roviery/catetin-api/entity"
	"github.com/roviery/catetin-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type transactionUsecase struct {
	transactionRepo domain.TransactionRepository
	financeRepo     domain.FinanceRepository
}

func NewTransactionUsecase(transactionRepo domain.TransactionRepository, financeRepo domain.FinanceRepository) *transactionUsecase {
	return &transactionUsecase{
		transactionRepo: transactionRepo,
		financeRepo:     financeRepo,
	}
}

func (t *transactionUsecase) CreateTransaction(ctx context.Context, ownerId *primitive.ObjectID, req *entity.TransactionRequest) (*entity.TransactionIDResponse, error) {
	// Better if use transaction
	// Get Finance
	finance, err := t.financeRepo.FindFinanceByID(ctx, &req.FinanceID)
	if err != nil {
		return nil, err
	}

	// Update Finance
	finance.Used += req.Expense
	finance.Remaining = finance.FundAllocation - finance.Used
	_, err = t.financeRepo.UpdateFinanceByID(ctx, &finance.ID, finance)
	if err != nil {
		return nil, err
	}

	// Insert Transaction
	res, err := t.transactionRepo.Store(ctx, &models.Transaction{
		OwnerID:   *ownerId,
		FinanceID: req.FinanceID,
		Title:     req.Title,
		Expense:   req.Expense,
	})
	if err != nil {
		return nil, err
	}

	return &entity.TransactionIDResponse{
		ID: *res,
	}, nil
}

func (t *transactionUsecase) GetTransaction(ctx context.Context, id *primitive.ObjectID) (*entity.TransactionResponse, error) {
	res, err := t.transactionRepo.FindTransactionByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.TransactionResponse{
		ID:        res.ID,
		FinanceID: res.FinanceID,
		Title:     res.Title,
		Expense:   res.Expense,
		CreatedAt: res.CreatedAt,
	}, nil
}

func (t *transactionUsecase) GetTransactions(ctx context.Context, ownerId *primitive.ObjectID, filter *string) (*entity.TransactionsResponse, error) {
	res, err := t.transactionRepo.FindTransactionByOwnerID(ctx, ownerId, filter)
	if err != nil {
		return nil, err
	}

	var transactionEntity []entity.TransactionResponse
	for _, transaction := range res {
		transactionEntity = append(transactionEntity, entity.TransactionResponse{
			ID:        transaction.ID,
			FinanceID: transaction.FinanceID,
			Title:     transaction.Title,
			Expense:   transaction.Expense,
			CreatedAt: transaction.CreatedAt,
		})
	}

	return &entity.TransactionsResponse{
		Transactions: transactionEntity,
	}, nil
}

func (t *transactionUsecase) UpdateTransactionByID(ctx context.Context, id *primitive.ObjectID, req *entity.TransactionRequest) (*entity.TransactionIDResponse, error) {
	// Get Old Transaction
	oldTransaction, err := t.transactionRepo.FindTransactionByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Get Finance
	finance, err := t.financeRepo.FindFinanceByID(ctx, &req.FinanceID)
	if err != nil {
		return nil, err
	}

	// Update Finance
	finance.Used += req.Expense - oldTransaction.Expense
	finance.Remaining = finance.FundAllocation - finance.Used
	_, err = t.financeRepo.UpdateFinanceByID(ctx, &finance.ID, finance)
	if err != nil {
		return nil, err
	}

	// Update Transaction
	res, err := t.transactionRepo.UpdateTransactionByID(ctx, id, &models.Transaction{
		FinanceID: req.FinanceID,
		Title:     req.Title,
		Expense:   req.Expense,
	})
	if err != nil {
		return nil, err
	}

	return &entity.TransactionIDResponse{
		ID: *res,
	}, nil
}

func (t *transactionUsecase) DeleteTransactionByID(ctx context.Context, id *primitive.ObjectID) error {
	// Get Transaction
	transaction, err := t.transactionRepo.FindTransactionByID(ctx, id)
	if err != nil {
		return err
	}

	// Get Finance
	finance, err := t.financeRepo.FindFinanceByID(ctx, &transaction.FinanceID)
	if err != nil {
		return err
	}

	// Update Finance
	finance.Used -= transaction.Expense
	finance.Remaining = finance.FundAllocation - finance.Used
	_, err = t.financeRepo.UpdateFinanceByID(ctx, &finance.ID, finance)
	if err != nil {
		return err
	}

	// Delete Transaction
	err = t.transactionRepo.DeleteTransactionByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
