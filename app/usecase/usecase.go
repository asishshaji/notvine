package usecase

import (
	"context"

	"github.com/asishshaji/notvine/app/entity"
	"github.com/asishshaji/notvine/app/repository"
)

type AppUsecase struct {
	repo repository.Mongorepo
}

func NewAppUsecase(
	userRepo repository.Mongorepo,
) *AppUsecase {
	return &AppUsecase{
		repo: userRepo,
	}

}

func (a *AppUsecase) Signup(ctx context.Context, username, password string) (*entity.User, error) {

	user := entity.User{
		Username: username,
		Password: password,
	}

	err := a.repo.CreateUser(ctx, &user)

	if err != nil {
		return nil, err
	}
	return &user, nil

}

func (a *AppUsecase) Login(ctx context.Context, username, password string) (*entity.User, error) {

	user, err := a.repo.CheckUsernamePassword(ctx, username, password)

	if err != nil {
		return nil, err
	}

	return user, nil
}
