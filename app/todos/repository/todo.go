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

type todoRepository struct {
	mgo *mongo.Client
}

func NewTodoRepository(mgo *mongo.Client) domain.TodoRepository {
	return &todoRepository{
		mgo: mgo,
	}
}

func (to *todoRepository) Store(ctx context.Context, todo *models.Todo) (*primitive.ObjectID, error) {
	collection := to.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionTodos)

	todo.ID = primitive.NewObjectID()
	todo.CreatedAt = time.Now().Format(time.RFC1123)
	todo.UpdatedAt = todo.CreatedAt
	_, err := collection.InsertOne(ctx, todo)
	if err != nil {
		return nil, err
	}

	return &todo.ID, nil
}

func (to *todoRepository) FindTodoByID(ctx context.Context, id *primitive.ObjectID) (*models.Todo, error) {
	var todo models.Todo

	collection := to.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionTodos)
	err := collection.FindOne(ctx, bson.M{
		"_id": id,
	}).Decode(&todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (to *todoRepository) FindTodoByOwnerID(ctx context.Context, ownerId *primitive.ObjectID) ([]models.Todo, error) {
	var todos []models.Todo

	collection := to.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionTodos)
	cursor, err := collection.Find(ctx, bson.M{
		"owner_id": ownerId,
	})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &todos)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (to *todoRepository) UpdateTodoByID(ctx context.Context, id *primitive.ObjectID, todo *models.Todo) (*primitive.ObjectID, error) {
	collection := to.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionTodos)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"todo":       todo.Todo,
		"status":     todo.Status,
		"updated_at": time.Now().Format(time.RFC1123),
	}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (to *todoRepository) DeleteTodoByID(ctx context.Context, id *primitive.ObjectID) error {
	collection := to.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionTodos)
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
