package usecase

import (
	"context"

	"github.com/asishshaji/notvine/app/entity"
)

type UsecaseInterface interface {
	Signup(ctx context.Context, username, password string) (*entity.User, error)
	// Signup(ctx context.Context, username, password string)
}
