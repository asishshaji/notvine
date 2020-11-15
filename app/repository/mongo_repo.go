package repository

import (
	"context"
	"errors"

	"github.com/asishshaji/notvine/app/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongorepo struct {
	db *mongo.Collection
}

func NewMongoRepo(db *mongo.Database, collection string) *Mongorepo {

	return &Mongorepo{
		db: db.Collection(collection),
	}

}
func (repo *Mongorepo) CreateUser(ctx context.Context, user *entity.User) error {

	exists, _ := repo.CheckUserExists(ctx, user)

	if exists == true {
		return errors.New("User already exists")
	}

	result, err := repo.db.InsertOne(ctx, user)

	if err != nil {
		return err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)
	return nil

}

func (repo *Mongorepo) CheckUserExists(ctx context.Context, user *entity.User) (bool, error) {

	err := repo.db.FindOne(ctx, bson.M{"username": user.Username})

	if err.Err() == nil {
		return true, err.Err()
	}

	return false, err.Err()
}
