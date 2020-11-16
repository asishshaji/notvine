package repository

import (
	"context"
	"errors"

	"github.com/asishshaji/notvine/app/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Mongorepo create Repo
type Mongorepo struct {
	db *mongo.Collection
}

// NewMongoRepo creates instance of Mongorepo
func NewMongoRepo(db *mongo.Database, collection string) *Mongorepo {
	return &Mongorepo{
		db: db.Collection(collection),
	}
}

func (repo *Mongorepo) CreatePost(ctx context.Context, post *entity.Post) error {
	return nil
}

// CreateUser creates a new user
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

// CheckUserExists checks if the user exists
func (repo *Mongorepo) CheckUserExists(ctx context.Context, user *entity.User) (bool, error) {

	res := repo.db.FindOne(ctx, bson.M{"username": user.Username})

	if res.Err() == nil {
		return true, res.Err()
	}

	return false, res.Err()
}

// CheckUsernamePassword returns user entity with passed username and password
func (repo *Mongorepo) CheckUsernamePassword(ctx context.Context, username, password string) (*entity.User, error) {

	user := entity.User{}

	err := repo.db.FindOne(ctx, bson.M{"username": username, "password": password}).Decode(&user)

	if err != nil {
		return nil, errors.New("username and password doesn't match")
	}

	return &user, nil

}
