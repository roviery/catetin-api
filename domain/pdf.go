package domain

import (
	"context"

	"github.com/roviery/catetin-api/entity"
	"github.com/roviery/catetin-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PDFRepository interface {
	Store(ctx context.Context, pdf *models.PDF) (*primitive.ObjectID, error)
	FindPDFByID(ctx context.Context, id *primitive.ObjectID) (*models.PDF, error)
}

type PDFUsecase interface {
	UploadPDF(ctx context.Context, req *entity.PDFRequest) (*entity.PDFIDResponse, error)
	DownloadPDF(ctx context.Context, id *primitive.ObjectID) (*entity.PDFResponse, error)
}
