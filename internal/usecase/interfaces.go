// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"
	"github.com/evrone/go-clean-template/internal/controller/http/dto"
	"github.com/evrone/go-clean-template/internal/entity"
)

type (
	User interface {
		GetUsers(ctx context.Context) ([]entity.User, error)
		GetUserByID(ctx context.Context, id string) (*entity.User, error)
	}

	UserRepo interface {
		GetUsers(ctx context.Context) (users []entity.User, err error)
		GetUser(ctx context.Context, id string) (user *entity.User, err error)
		GetUserByEmail(ctx context.Context, email string) (user *entity.User, err error)
		CreateUser(ctx context.Context, user *entity.User) (string, error)
	}
	Auth interface {
		Register(ctx context.Context, name, email, password string) error
		Login(ctx context.Context, email, password string) (*dto.LoginResponse, error)
	}
)
