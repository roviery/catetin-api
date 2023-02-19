package usecase

import (
	"context"

	"github.com/roviery/catetin-api/domain"
	"github.com/roviery/catetin-api/entity"
	"github.com/roviery/catetin-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type pdfUsecase struct {
	pdfRepo domain.PDFRepository
}

func NewPDFUsecase(pdfRepo domain.PDFRepository) domain.PDFUsecase {
	return &pdfUsecase{
		pdfRepo: pdfRepo,
	}
}

func (p *pdfUsecase) UploadPDF(ctx context.Context, req *entity.PDFRequest) (*entity.PDFIDResponse, error) {
	res, err := p.pdfRepo.Store(ctx, &models.PDF{
		Name: req.Name,
		Data: req.Data,
	})
	if err != nil {
		return nil, err
	}

	return &entity.PDFIDResponse{
		ID: *res,
	}, nil
}

func (p *pdfUsecase) DownloadPDF(ctx context.Context, id *primitive.ObjectID) (*entity.PDFResponse, error) {
	pdf, err := p.pdfRepo.FindPDFByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.PDFResponse{
		ID:   pdf.ID,
		Name: pdf.Name,
		Data: pdf.Data,
	}, nil
}
