package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID           primitive.ObjectID `json:"postid"`
	URL          string             `json:"videourl"`
	Owner        *User              `json:"posted_by"`
	Caption      string             `json:"caption"`
	LikesCount   int                `json:"likes"`
	LikedBy      []string           `json:"liked_by"`
	ThumbnailURL string             `json:"thumbnail_url"`
	CreatedAt    time.Time          `json:"created_at"`
}
