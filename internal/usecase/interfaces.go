// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"
	"github.com/damndelion/sdu-go-final/internal/controller/http/dto"
	"github.com/damndelion/sdu-go-final/internal/entity"
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

	Menu interface {
		GetMenu(ctx context.Context) ([]entity.Menu, error)
		CreateMenuItem(ctx context.Context, item dto.CreateMenuItem) (id string, err error)
		UpdateMenuItem(ctx context.Context, item dto.UpdateMenuItem) (id string, err error)
		DeleteMenuItem(ctx context.Context, id string) error
	}
	MenuRepo interface {
		GetAllMenu(ctx context.Context) (users []entity.Menu, err error)
		CreateMenuItem(ctx context.Context, item dto.CreateMenuItem) (id string, err error)
		UpdateMenuItem(ctx context.Context, item dto.UpdateMenuItem) (id string, err error)
		DeleteMenuItem(ctx context.Context, id string) error
	}

	Order interface {
		GetAllOrder(ctx context.Context) (users []entity.Order, err error)
		CreateOrderItem(ctx context.Context, item dto.CreateOderItem) (id string, err error)
		UpdateOderItem(ctx context.Context, item dto.UpdateOrderItem) (id string, err error)
		DeleteOrderItem(ctx context.Context, id string) error
	}

	OderRepo interface {
		GetAllOrder(ctx context.Context) (users []entity.Order, err error)
		CreateOrderItem(ctx context.Context, item dto.CreateOderItem) (id string, err error)
		UpdateOrderItem(ctx context.Context, item dto.UpdateOrderItem) (id string, err error)
		DeleteOrderItem(ctx context.Context, id string) error
	}
)
