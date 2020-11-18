package repository

import (
	"context"

	"github.com/asishshaji/notvine/app/entity"
)

// RepoInterface creates interface for repository
type RepoInterface interface {
	// Users
	CreateUser(ctx context.Context, user *entity.User) error
	CheckUserExists(ctx context.Context, user *entity.User) (bool, error)
	GetUser(ctx context.Context, username string) (*entity.User, error)
	CreatePost(ctx context.Context, post *entity.Post) error
	// Login(ctx context.Context, username, password string) (*entity.User, error)
	CheckUsernamePassword(ctx context.Context, username, password string) (*entity.User, error)
	// GetUserFeed(uid string) ([]*models.Post, error)

	// // Posts
	// LikePost(pid string, uid string) error
	// CommentPost(pid string, uid string) error
}
