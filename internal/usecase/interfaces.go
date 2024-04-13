// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/evrone/go-clean-template/internal/entity"
)

type (
	// Translation -.
	User interface {
		GetUsers(ctx context.Context) ([]entity.User, error)
		GetUserByID(ctx context.Context, id string) (*entity.User, error)
	}

	UserRepo interface {
		GetUsers(ctx context.Context) (users []entity.User, err error)
		GetUser(ctx context.Context, id string) (user *entity.User, err error)
	}
)
