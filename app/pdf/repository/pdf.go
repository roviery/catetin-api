package repository

import (
	"context"
	"time"

	"github.com/roviery/catetin-api/constant"
	"github.com/roviery/catetin-api/domain"
	"github.com/roviery/catetin-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type pdfRepo struct {
	mgo *mongo.Client
}

func NewPDFRepo(mgo *mongo.Client) domain.PDFRepository {
	return &pdfRepo{
		mgo: mgo,
	}
}

func (p *pdfRepo) Store(ctx context.Context, pdf *models.PDF) (*primitive.ObjectID, error) {
	collection := p.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionPDF)

	pdf.ID = primitive.NewObjectID()
	pdf.CreatedAt = time.Now().Format(time.RFC1123)
	pdf.UpdatedAt = pdf.CreatedAt
	_, err := collection.InsertOne(ctx, pdf)
	if err != nil {
		return nil, err
	}

	return &pdf.ID, nil
}

func (p *pdfRepo) FindPDFByID(ctx context.Context, id *primitive.ObjectID) (*models.PDF, error) {
	var pdf models.PDF
	collection := p.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionPDF)
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&pdf)
	if err != nil {
		return nil, err
	}

	return &pdf, nil
}
