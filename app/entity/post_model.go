package entity

import (
	"time"
)

type Post struct {
	URL          string    `json:"videourl"`
	Owner        *User     `json:"posted_by"`
	Caption      string    `json:"caption"`
	LikesCount   int       `json:"likes"`
	LikedBy      []string  `json:"liked_by"`
	ThumbnailURL string    `json:"thumbnail_url"`
	CreatedAt    time.Time `json:"created_at"`
}
