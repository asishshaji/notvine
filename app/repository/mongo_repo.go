package repository

import (
	"context"
	"errors"
	"log"

	"github.com/asishshaji/notvine/app/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Mongorepo create Repo
type Mongorepo struct {
	userCollection *mongo.Collection
	postCollection *mongo.Collection
}

// NewMongoRepo creates instance of Mongorepo
func NewMongoRepo(db *mongo.Database, userCollection, postCollection string) RepoInterface {
	return Mongorepo{
		userCollection: db.Collection(userCollection),
		postCollection: db.Collection(postCollection),
	}
}

// CreatePost creates a new post
func (repo Mongorepo) CreatePost(ctx context.Context, post *entity.Post) error {
	res, err := repo.postCollection.InsertOne(ctx, post)
	if err != nil {
		return err
	}

	log.Println("New post created : ", res)

	return nil

}

// CreateUser creates a new user
func (repo Mongorepo) CreateUser(ctx context.Context, user *entity.User) error {

	exists, _ := repo.CheckUserExists(ctx, user)

	if exists == true {
		return errors.New("User already exists")
	}

	result, err := repo.userCollection.InsertOne(ctx, user)

	if err != nil {
		return err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)
	return nil

}

// CheckUserExists checks if the user exists
// Can also be used to check if username is available
func (repo Mongorepo) CheckUserExists(ctx context.Context, user *entity.User) (bool, error) {

	res := repo.userCollection.FindOne(ctx, bson.M{"username": user.Username})

	if res.Err() == nil {
		return true, res.Err()
	}

	return false, res.Err()
}

func (repo Mongorepo) GetUser(ctx context.Context, username string) (*entity.User, error) {

	user := entity.User{}
	res := repo.userCollection.FindOne(ctx, bson.M{"username": username})

	err := res.Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CheckUsernamePassword returns user entity with passed username and password
func (repo Mongorepo) CheckUsernamePassword(ctx context.Context, username, password string) (*entity.User, error) {

	user := entity.User{}

	err := repo.userCollection.FindOne(ctx, bson.M{"username": username, "password": password}).Decode(&user)

	if err != nil {
		return nil, errors.New("username and password doesn't match")
	}

	return &user, nil

}
