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

type deadlineRepo struct {
	mgo *mongo.Client
}

func NewDeadlineRepo(mgo *mongo.Client) domain.DeadlineRepository {
	return &deadlineRepo{
		mgo: mgo,
	}
}

func (d *deadlineRepo) Store(ctx context.Context, deadline *models.Deadline) (*primitive.ObjectID, error) {
	collection := d.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionDeadlines)

	deadline.ID = primitive.NewObjectID()
	deadline.CreatedAt = time.Now().Format(time.RFC1123)
	deadline.UpdatedAt = deadline.CreatedAt
	_, err := collection.InsertOne(ctx, deadline)
	if err != nil {
		return nil, err
	}

	return &deadline.ID, nil
}

func (d *deadlineRepo) FindDeadlineByID(ctx context.Context, id *primitive.ObjectID) (*models.Deadline, error) {
	var deadline models.Deadline

	collection := d.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionDeadlines)
	err := collection.FindOne(ctx, bson.M{
		"_id": id,
	}).Decode(&deadline)
	if err != nil {
		return nil, err
	}

	return &deadline, nil
}

func (d *deadlineRepo) FindDeadlinesByOwnerID(ctx context.Context, ownerId *primitive.ObjectID) ([]models.Deadline, error) {
	var deadlines []models.Deadline

	collection := d.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionDeadlines)
	cursor, err := collection.Find(ctx, bson.M{
		"owner_id": ownerId,
	})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &deadlines)
	if err != nil {
		return nil, err
	}

	return deadlines, nil
}

func (d *deadlineRepo) UpdateDeadlineByID(ctx context.Context, id *primitive.ObjectID, deadline *models.Deadline) (*primitive.ObjectID, error) {
	collection := d.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionDeadlines)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"deadline_date": deadline.DeadlineDate,
		"task":          deadline.Task,
		"priority":      deadline.Priority,
		"updated_at":    time.Now().Format(time.RFC1123),
	}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (d *deadlineRepo) DeleteDeadlineByID(ctx context.Context, id *primitive.ObjectID) error {
	collection := d.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionDeadlines)
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
