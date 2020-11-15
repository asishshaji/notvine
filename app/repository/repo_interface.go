package repository

import (
	"context"

	"github.com/asishshaji/notvine/app/entity"
)

// RepositoryInterface creates interface for repository
type RepositoryInterface interface {
	// Users
	CreateUser(ctx context.Context, user *entity.User) error
	CheckUserExists(ctx context.Context, user *entity.User) (bool, error)
	GetUser(username, password string) (*entity.User, error)
	Login(ctx context.Context, username, password string) (*entity.User, error)
	// GetUserFeed(uid string) ([]*models.Post, error)

	// // Posts
	// LikePost(pid string, uid string) error
	// CommentPost(pid string, uid string) error
}
