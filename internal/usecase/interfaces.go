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
		GetOrder(ctx context.Context) (users []entity.Order, err error)
		GetCurrentOrder(ctx context.Context) ([]entity.Order, error)
		CreateOrderItem(ctx context.Context, item dto.CreateOrderItemRequest, userID string) (id string, err error)
		UpdateOrderItem(ctx context.Context, item dto.UpdateOrderItem) (id string, err error)
		DeleteOrderItem(ctx context.Context, id string) error
		GetUserAllOrder(ctx context.Context, userId string) ([]entity.Order, error)
		GetUserCurrentOrder(ctx context.Context, userId string) ([]entity.Order, error)
	}

	OderRepo interface {
		GetAllOrder(ctx context.Context) (order []entity.Order, err error)
		CreateOrderItem(ctx context.Context, item dto.CreateOrderItemRequest, userId string) (id string, err error)
		UpdateOrderItem(ctx context.Context, item dto.UpdateOrderItem) (id string, err error)
		DeleteOrderItem(ctx context.Context, id string) error
		GetUserCurrentOrders(ctx context.Context, userId string) (order []entity.Order, err error)
		GetUserAllOrders(ctx context.Context, userId string) (order []entity.Order, err error)
		GetAllCurrentOrder(ctx context.Context) (order []entity.Order, err error)
	}

	OrderMenu interface {
		GetOrderMenu(ctx context.Context) (users []entity.OrderMenu, err error)
		UpdateOrderMenuItem(ctx context.Context, item dto.UpdateOrderMenuItem) (id string, err error)
		DeleteOrderMenuItem(ctx context.Context, id string) error
	}
	OrderMenuRepo interface {
		GetAllOrderMenu(ctx context.Context) (users []entity.OrderMenu, err error)
		UpdateOrderMenuItem(ctx context.Context, item dto.UpdateOrderMenuItem) (id string, err error)
		DeleteOrderMenuItem(ctx context.Context, id string) error
	}

	Stock interface {
		GetStock(ctx context.Context) (users []entity.Stock, err error)
		CreateStockItem(ctx context.Context, item dto.CreateStockItem) (id string, err error)
		UpdateStockItem(ctx context.Context, item dto.UpdateStockItem) (id string, err error)
		DeleteStockItem(ctx context.Context, id string) error
	}

	StockRepo interface {
		GetAllStock(ctx context.Context) (users []entity.Stock, err error)
		CreateStockItem(ctx context.Context, item dto.CreateStockItem) (id string, err error)
		UpdateStockItem(ctx context.Context, item dto.UpdateStockItem) (id string, err error)
		DeleteStockItem(ctx context.Context, id string) error
	}
)
