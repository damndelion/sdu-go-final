package usecase

import (
	"context"
	"fmt"
	"github.com/damndelion/sdu-go-final/internal/controller/http/dto"

	"github.com/damndelion/sdu-go-final/internal/entity"
)

type OrderUseCase struct {
	repo OderRepo
}

func NewOrderUseCase(r OderRepo) *OrderUseCase {
	return &OrderUseCase{
		repo: r,
	}
}

func (mc *MenuUseCase) GetOrder(ctx context.Context) ([]entity.Order, error) {
	menu, err := mc.repo.GetAllMenu(ctx)
	if err != nil {
		return nil, fmt.Errorf("MenuUseCase - GetMenu: %w", err)
	}

	return menu, nil
}

func (mc *MenuUseCase) CreateMenuItem(ctx context.Context, item dto.CreateMenuItem) (id string, err error) {
	menu, err := mc.repo.CreateMenuItem(ctx, item)
	if err != nil {
		return "", fmt.Errorf("MenuUseCase - CreateMenuItem: %w", err)
	}

	return menu, nil
}

func (mc *MenuUseCase) UpdateMenuItem(ctx context.Context, item dto.UpdateMenuItem) (id string, err error) {
	menu, err := mc.repo.UpdateMenuItem(ctx, item)
	if err != nil {
		return "", fmt.Errorf("MenuUseCase - UpdateMenuItem: %w", err)
	}

	return menu, nil
}

func (mc *MenuUseCase) DeleteMenuItem(ctx context.Context, id string) error {
	err := mc.repo.DeleteMenuItem(ctx, id)
	if err != nil {
		return fmt.Errorf("MenuUseCase - DeleteMenuItme: %w", err)
	}

	return nil
}
