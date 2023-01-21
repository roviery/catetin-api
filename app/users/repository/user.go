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

type userRepo struct {
	mgo *mongo.Client
}

func NewUserRepo(mgo *mongo.Client) domain.UserRepository {
	return &userRepo{
		mgo: mgo,
	}
}

func (u *userRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	collection := u.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionUsers)

	err := collection.FindOne(ctx, bson.M{
		"email": email,
	}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepo) Store(ctx context.Context, user *models.User) (*models.User, error) {
	collection := u.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionUsers)

	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now().Format(time.RFC1123)
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
