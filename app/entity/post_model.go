package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID    primitive.ObjectID
	URL   string
	Owner *User
}
