package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/damndelion/sdu-go-final/config"
	"github.com/damndelion/sdu-go-final/internal/controller/http/dto"
	"github.com/damndelion/sdu-go-final/internal/entity"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	cfg  *config.Config
	repo UserRepo
}

func NewAuthUseCase(cfg *config.Config, repo UserRepo) *AuthUseCase {
	return &AuthUseCase{cfg, repo}
}

func (u *AuthUseCase) Register(ctx context.Context, name, email, password string) error {
	user, err := u.repo.GetUserByEmail(ctx, email)
	if user.ID != "" {
		return fmt.Errorf("user with this email alraedy exists")
	}
	if err != nil {
		return err
	}

	_, err = u.repo.CreateUser(ctx, &entity.User{
		Name:     name,
		Email:    email,
		Password: password,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *AuthUseCase) Login(ctx context.Context, email, password string) (*dto.LoginResponse, error) {

	user, err := u.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("passwords do not match %v", err))
	}
	accessToken, err := u.generateTokens(ctx, user)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Name:        user.Name,
		Email:       user.Email,
		Role:        user.Role,
		AccessToken: accessToken,
	}, nil
}

func (u *AuthUseCase) generateTokens(ctx context.Context, user *entity.User) (string, error) {
	accessTokenClaims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"name":    user.Name,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Duration(u.cfg.AccessTokenTTL) * time.Second).Unix(),
	}

	access := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := access.SignedString([]byte(u.cfg.SecretKey))
	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}
