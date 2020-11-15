package repository

import (
	"context"
	"log"

	"github.com/asishshaji/notvine/app/entity"
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

	log.Println(user)
	result, err := repo.db.InsertOne(ctx, user)

	if err != nil {
		return err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)
	return nil

}
