package repo

import (
	"context"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (user *entity.User, err error) {
	res := r.db.Where("email = ?", email).WithContext(ctx).Find(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (r *UserRepo) CreateUser(ctx context.Context, user *entity.User) (string, error) {
	generatedHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user_uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	newUser := entity.User{
		ID:       user_uuid.String(),
		Name:     user.Name,
		Email:    user.Email,
		Password: string(generatedHash),
		Role:     "user",
	}

	res := r.db.WithContext(ctx).Create(&newUser)
	if res.Error != nil {
		return "", res.Error
	}

	return user_uuid.String(), nil
}
