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

type quicknoteRepo struct {
	mgo *mongo.Client
}

func NewQuicknoteRepo(mgo *mongo.Client) domain.QuicknoteRepository {
	return &quicknoteRepo{
		mgo: mgo,
	}
}

func (q *quicknoteRepo) Store(ctx context.Context, quicknote *models.Quicknote) (*primitive.ObjectID, error) {
	collection := q.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionQuicknotes)
	quicknote.ID = primitive.NewObjectID()
	quicknote.CreatedAt = time.Now().Format(time.RFC1123)
	quicknote.UpdatedAt = quicknote.CreatedAt
	_, err := collection.InsertOne(ctx, quicknote)
	if err != nil {
		return nil, err
	}

	return &quicknote.ID, nil
}

func (q *quicknoteRepo) FindQuicknoteById(ctx context.Context, id *primitive.ObjectID) (*models.Quicknote, error) {
	var quicknote models.Quicknote

	collection := q.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionQuicknotes)
	err := collection.FindOne(ctx, bson.M{
		"_id": id,
	}).Decode(&quicknote)
	if err != nil {
		return nil, err
	}

	return &quicknote, nil
}

func (q *quicknoteRepo) FindQuicknoteByOwnerId(ctx context.Context, ownerId *primitive.ObjectID) ([]models.Quicknote, error) {
	var quicknotes []models.Quicknote

	collection := q.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionQuicknotes)
	cursor, err := collection.Find(ctx, bson.M{
		"owner_id": ownerId,
	})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &quicknotes)
	if err != nil {
		return nil, err
	}

	return quicknotes, nil
}

func (q *quicknoteRepo) UpdateQuicknoteByID(ctx context.Context, id *primitive.ObjectID, quicknote *models.Quicknote) (*primitive.ObjectID, error) {
	collection := q.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionQuicknotes)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"note":       quicknote.Note,
		"updated_at": time.Now().Format(time.RFC1123),
	}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (q *quicknoteRepo) DeleteQuicknoteByID(ctx context.Context, id *primitive.ObjectID) error {
	collection := q.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionQuicknotes)
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
