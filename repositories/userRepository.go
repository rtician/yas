package repositories

import (
	"context"
	"yas/database/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type userRepository struct {
	Database   *mongo.Database
	Collection string
}

func NewUserRepository(d *mongo.Database) UserRepository {
	return &userRepository{
		Database:   d,
		Collection: "user",
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	user.Id = primitive.NewObjectID()
	collection := r.Database.Collection(r.Collection)
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	collection := r.Database.Collection(r.Collection)
	user := &models.User{}
	if err := collection.FindOne(ctx, bson.M{"email": email}).Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}
