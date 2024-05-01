package repo

import "C"
import (
	"context"
	"errors"
	"github.com/damndelion/sdu-go-final/internal/controller/http/dto"
	"github.com/damndelion/sdu-go-final/internal/entity"
	"gorm.io/gorm"
)

type OrderMenuRepo struct {
	db *gorm.DB
}

func NewOderMenuRepo(pg *gorm.DB) *OrderMenuRepo {
	return &OrderMenuRepo{pg}
}

func (r *OrderMenuRepo) GetAllOrderMenu(ctx context.Context) (orderMenu []entity.OrderMenu, err error) {
	res := r.db.Find(&orderMenu)
	if res.Error != nil {
		return nil, res.Error
	}

	return orderMenu, nil
}

func (r *OrderMenuRepo) UpdateOrderMenuItem(ctx context.Context, item dto.UpdateOrderMenuItem) (id string, err error) {
	orderMenu := entity.OrderMenu{
		OrderID: item.OrderID,
		MenuID:  item.MenuID,
	}
	res := r.db.Model(&orderMenu).Where("id = ?", item.ID).Updates(&orderMenu)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrDuplicatedKey) {
			return "", errors.New("order with that name already exists")
		}
		return "", res.Error
	}

	return item.ID, nil
}

func (r *OrderMenuRepo) DeleteOrderMenuItem(ctx context.Context, id string) error {
	res := r.db.Model(&entity.OrderMenu{}).Where("id = ?", id).Delete(&entity.OrderMenu{})
	if res.Error != nil {
		return res.Error
	}

	return nil
}
