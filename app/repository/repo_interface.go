package repository

import (
	"context"

	"github.com/asishshaji/notvine/app/entity"
)

type RepositoryInterface interface {
	// User
	CreateUser(ctx context.Context, user *entity.User) error
	CheckUserExists(ctx context.Context, user *entity.User) (bool, error)
	// GetUser(username, password string) (*models.User, error)
	// GetUserFeed(uid string) ([]*models.Post, error)

	// // Posts
	// LikePost(pid string, uid string) error
	// CommentPost(pid string, uid string) error
}
