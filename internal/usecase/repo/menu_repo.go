package repo

import (
	"context"
	"errors"
	"github.com/damndelion/sdu-go-final/internal/controller/http/dto"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/damndelion/sdu-go-final/internal/entity"
)

type MenuRepo struct {
	db *gorm.DB
}

// NewMenuRepo -.
func NewMenuRepo(pg *gorm.DB) *MenuRepo {
	return &MenuRepo{pg}
}

func (r *MenuRepo) GetAllMenu(ctx context.Context) (menu []entity.Menu, err error) {
	res := r.db.Find(&menu)
	if res.Error != nil {
		return nil, res.Error
	}

	return menu, nil
}

func (r *MenuRepo) CreateMenuItem(ctx context.Context, item dto.CreateMenuItem) (id string, err error) {
	menuUuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	menu := entity.Menu{
		ID:          menuUuid.String(),
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
	}
	res := r.db.Create(&menu)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrDuplicatedKey) {
			return "", errors.New("menu with that name already exists")
		}
		return "", res.Error
	}

	return item.Name, nil
}

func (r *MenuRepo) UpdateMenuItem(ctx context.Context, item dto.UpdateMenuItem) (id string, err error) {
	menu := entity.Menu{
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
	}
	res := r.db.Model(&menu).Where("id = ?", item.ID).Updates(&menu)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrDuplicatedKey) {
			return "", errors.New("menu with that name already exists")
		}
		return "", res.Error
	}

	return item.Name, nil
}

func (r *MenuRepo) DeleteMenuItem(ctx context.Context, id string) error {
	res := r.db.Model(&entity.Menu{}).Where("id = ?", id).Delete(&entity.Menu{})
	if res.Error != nil {
		return res.Error
	}

	return nil
}
