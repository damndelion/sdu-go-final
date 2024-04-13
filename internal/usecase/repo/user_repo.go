package repo

import (
	"context"
	"gorm.io/gorm"

	"github.com/evrone/go-clean-template/internal/entity"
)

type UserRepo struct {
	db *gorm.DB
}

// New -.
func New(pg *gorm.DB) *UserRepo {
	return &UserRepo{pg}
}

func (r *UserRepo) GetUsers(ctx context.Context) (users []entity.User, err error) {
	res := r.db.Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}

	return users, nil

}

func (r *UserRepo) GetUser(ctx context.Context, id string) (user *entity.User, err error) {
	err = r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
