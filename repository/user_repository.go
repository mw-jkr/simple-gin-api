package repository

import (
	"context"
	"example/test-golang/db"
	"example/test-golang/entity"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository interface {
	GetUserById(ctx context.Context, id int) (*entity.User, error)
}

type userRepositoryImpl struct {
	db *db.DB
}

func (u userRepositoryImpl) GetUserById(ctx context.Context, id int) (*entity.User, error) {
	collection := u.getCollection()
	filter := bson.D{{Key: "id", Value: id}}

	result := new(entity.User)
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u userRepositoryImpl) getCollection() *mongo.Collection {
	return u.db.GetDB().Collection("users")
}

func NewUserRepository(db *db.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

func ProvideUserRepository(db *db.DB) UserRepository {
	return NewUserRepository(db)
}
