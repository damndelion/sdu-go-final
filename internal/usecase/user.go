package usecase

import (
	"context"
	"fmt"

	"github.com/evrone/go-clean-template/internal/entity"
)

type UserUseCase struct {
	repo UserRepo
}

func New(r UserRepo) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}

func (uc *UserUseCase) GetUsers(ctx context.Context) ([]entity.User, error) {
	users, err := uc.repo.GetUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("UserUseCase - GetUsers: %w", err)
	}

	return users, nil
}

func (uc *UserUseCase) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	user, err := uc.repo.GetUser(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("UserUseCase - GetUserByID: %w", err)
	}

	return user, nil
}
