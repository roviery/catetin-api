package repository

import (
	"context"
	"time"

	"github.com/roviery/catetin-api/constant"
	"github.com/roviery/catetin-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type transactionRepo struct {
	mgo *mongo.Client
}

func NewTransactionRepo(mgo *mongo.Client) *transactionRepo {
	return &transactionRepo{
		mgo: mgo,
	}
}

func (t *transactionRepo) Store(ctx context.Context, transaction *models.Transaction) (*primitive.ObjectID, error) {
	transcationColl := t.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionTransactions)
	transaction.ID = primitive.NewObjectID()
	transaction.CreatedAt = time.Now().Format(time.RFC1123)
	transaction.UpdatedAt = transaction.CreatedAt

	_, err := transcationColl.InsertOne(ctx, transaction)
	if err != nil {
		return nil, err
	}

	return &transaction.ID, nil
}

func (t *transactionRepo) FindTransactionByID(ctx context.Context, id *primitive.ObjectID) (*models.Transaction, error) {
	var transaction models.Transaction

	collection := t.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionTransactions)
	err := collection.FindOne(ctx, bson.M{
		"_id": id,
	}).Decode(&transaction)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (t *transactionRepo) FindTransactionByOwnerID(ctx context.Context, ownerId *primitive.ObjectID, filter *string) ([]models.Transaction, error) {
	// TODO: Filter feature
	var transactions []models.Transaction

	collection := t.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionTransactions)
	cursor, err := collection.Find(ctx, bson.M{
		"owner_id": ownerId,
	})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (t *transactionRepo) UpdateTransactionByID(ctx context.Context, id *primitive.ObjectID, transaction *models.Transaction) (*primitive.ObjectID, error) {
	collection := t.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionTransactions)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"finance_id": transaction.FinanceID,
		"title":      transaction.Title,
		"expense":    transaction.Expense,
		"updated_at": time.Now().Format(time.RFC1123),
	}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (t *transactionRepo) DeleteTransactionByID(ctx context.Context, id *primitive.ObjectID) error {
	collection := t.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionTransactions)
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
