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

type financeRepo struct {
	mgo *mongo.Client
}

func NewFinanceRepo(mgo *mongo.Client) domain.FinanceRepository {
	return &financeRepo{
		mgo: mgo,
	}
}

func (f *financeRepo) Store(ctx context.Context, finance *models.Finance) (*primitive.ObjectID, error) {
	collection := f.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionFinances)
	finance.ID = primitive.NewObjectID()
	finance.CreatedAt = time.Now().Format(time.RFC1123)
	finance.UpdatedAt = finance.CreatedAt
	_, err := collection.InsertOne(ctx, finance)
	if err != nil {
		return nil, err
	}

	return &finance.ID, nil
}

func (f *financeRepo) FindFinanceByID(ctx context.Context, id *primitive.ObjectID) (*models.Finance, error) {
	var finance models.Finance

	collection := f.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionFinances)
	err := collection.FindOne(ctx, bson.M{
		"_id": id,
	}).Decode(&finance)
	if err != nil {
		return nil, err
	}

	return &finance, nil
}

func (f *financeRepo) FindFinanceByType(ctx context.Context, financeType *string) (*models.Finance, error) {
	var finance models.Finance

	collection := f.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionFinances)
	err := collection.FindOne(ctx, bson.M{
		"type": financeType,
	}).Decode(&finance)
	if err != nil {
		return nil, err
	}

	return &finance, nil
}

func (f *financeRepo) FindFinanceByOwnerID(ctx context.Context, ownerId *primitive.ObjectID) ([]models.Finance, error) {
	var finances []models.Finance

	collection := f.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionFinances)
	cursor, err := collection.Find(ctx, bson.M{
		"owner_id": ownerId,
	})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &finances)
	if err != nil {
		return nil, err
	}

	return finances, nil
}

func (f *financeRepo) UpdateFinanceByID(ctx context.Context, id *primitive.ObjectID, finance *models.Finance) (*primitive.ObjectID, error) {
	collection := f.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionFinances)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"type":            finance.Type,
		"fund_allocation": finance.FundAllocation,
		"used":            finance.Used,
		"remaining":       finance.Remaining,
		"updated_at":      time.Now().Format(time.RFC1123),
	}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (f *financeRepo) DeleteFinanceByID(ctx context.Context, id *primitive.ObjectID) error {
	collection := f.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionFinances)
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
